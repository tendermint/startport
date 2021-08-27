// std and 3rd party imports.
import os from "os";
import fs from "fs";
import { join } from "path";
import yaml from "js-yaml";

// cosmosjs related imports.
import { Bip39, Random } from "@cosmjs/crypto";
import { DirectSecp256k1Wallet } from "@cosmjs/proto-signing";
import { 
	QueryClient,
  	setupBankExtension,

	GasPrice,
} from "@cosmjs/stargate";

import { Tendermint34Client } from "@cosmjs/tendermint-rpc";
import { stringToPath } from "@cosmjs/crypto";
import { Coin } from "@cosmjs/stargate";
import { Link, IbcClient } from "@confio/relayer/build";
import { orderFromJSON } from "@confio/relayer/build/codec/ibc/core/channel/v1/channel";

// local imports.
import Errors from "./errors";
import ConsoleLogger from "./logger";

const calcGasLimits = (limit: number) => ({
	initClient: 150000,
	updateClient: 600000,
	initConnection: 150000,
	connectionHandshake: limit,
	initChannel: 150000,
	channelHandshake: limit,
	receivePacket: limit,
	ackPacket: limit,
	timeoutPacket: limit,
	transfer: 180000,
});

// ***
// define types for relayer's config.yml.
// ***
//
type RelayerConfig = {
	chains?: Array<ChainConfig>;
	paths?: Array<PathConfig>;
};

type PathConfig = {
	path: Path;
	options?: ConnectOptions;
	connections?: Connections;
	relayerData?: PacketHeights;
};

type ChainConfig = {
	chainId: string;
	account: string,
	rpcAddr: string;
	addressPrefix: string;
	gasPrice: string;
	gasLimit: number;
};

// ***
// define internal types.
// ***
//
type Account = {
	address: string;
};

type ConnectOptions = {
	sourcePort: string;
	sourceVersion: string;
	targetPort: string;
	targetVersion: string;
	ordering: "ORDER_UNORDERED" | "ORDER_ORDERED";
};

type Connections = {
	srcConnection: string;
	destConnection: string;
};

type Endpoint = {
	channelID?: string;
	chainID: string;
	portID: string;
};

type Path = {
	id: string;
	isLinked: boolean;
	src: Endpoint;
	dst: Endpoint;
};

type PacketHeights = {
	packetHeightA: number;
	packetHeightB: number;
	ackHeightA: number;
	ackHeightB: number;
};


interface ChainSetupOptions {
	addressPrefix: string;
	gasPrice: string;
	gasLimit: number;
}

type EnsureChainSetupResponse = {
	// id is chain id.
	id: string;
};

type LinkError = {
	pathName: string;
	error: string;
};

type LinkResponse = {
	linkedPaths: string[];
	alreadyLinkedPaths: string[];
	failedToLinkPaths: LinkError[];
};

type LinkStatus = {
	status: boolean;
	pathName: string;
	error?: string;
};

type StartResponse = {};

type InfoResponse = {
	configPath: string;
};

export default class Relayer {
	private configDir: string = ".starport/relayer";
	private configFile: string = "config.yml";
	private ibcSetupGas: number = 2256000;
	private defaultMaxAge: number = 86400;
	private pollTime: number;
	private config: RelayerConfig;
	private homedir: string;

	constructor(pollTime = 5000) {
		this.homedir = os.homedir();
		this.pollTime = pollTime;
		this.ensureConfigDirCreated();
		this.initConfigProxy();
	}

	public async link([paths]: [string[]]): Promise<LinkResponse> {
		if (!this.config.paths)
			throw Errors.PathsNotDefined;

		let response: LinkResponse = {
			linkedPaths: [],
			alreadyLinkedPaths: [],
			failedToLinkPaths: [],
		};

		const results = [];

		for (let pathName of paths) {
			const path = this.pathById(pathName);

			if (path?.path.isLinked) {
				response.alreadyLinkedPaths.push(pathName);
				continue;
			}

			results.push(await this.createLink(path));
		}

		for (let result of results) {
			if (result.status) {
				response.linkedPaths.push(result.pathName);
			} else {
				response.failedToLinkPaths.push({
					pathName: result.pathName,
					error: result.error,
				});
			}
		}

		return response;
	}

	public async start([paths]: [string[]]): Promise<StartResponse> {
		if (!this.config.paths)
			throw Errors.PathsNotDefined;

		for (let pathName of paths) {
			const path = this.pathById(pathName);

			if (path?.path.isLinked) {
				const link = await this.getLink(path);
				setInterval(async () => {
					let heights = this.pathById(pathName).relayerData;
					let newHeights = await this.relayPackets(link, heights);
					this.pathById(pathName).relayerData = newHeights;
				}, this.pollTime)

				continue;
			}

			throw Errors.PathNotLinked;
		}

		return {};
	}

	public async info(): Promise<InfoResponse> {
		return { configPath: this.getConfigPath() };
	}


	private initConfigProxy() {
		const nestedProxy = {
			set: (target, prop, value) => {
				target[prop] = value;
				this.writeConfig(this.config);
				return true;
			},

			get: (target, prop) => {
				if (typeof target[prop] === "object" && target[prop] !== null)
					return new Proxy(target[prop], nestedProxy);
				return target[prop];
			},
		};

		this.config = new Proxy(this.readOrCreateConfig(), nestedProxy);
	}

	private getConfigDirPath() {
		return join(this.homedir, this.configDir);
	}

	private getConfigPath() {
		return join(this.getConfigDirPath(), this.configFile);
	}

	private ensureConfigDirCreated() {
		try {
			if (!fs.existsSync(this.getConfigDirPath()))
				fs.mkdirSync(this.getConfigDirPath(), { recursive: true });
		} catch (e) {
			throw Errors.ConfigFolderFailed(e);
		}
	}

	private readOrCreateConfig(): RelayerConfig {
		// return the config if already exists.
		try {
			if (fs.existsSync(this.getConfigPath())) {
				let configFile = fs.readFileSync(this.getConfigPath(), "utf8");
				return yaml.load(configFile);
			}
		} catch (e) {
			throw Errors.ConfigReadFailed(e);
		}

		// there is no config, create one and return it.
		let config = {
		};

		this.writeConfig(config);

		return config;
	}

	private writeConfig(config) {
		try {
			let configFile = yaml.dump(config);
			fs.writeFileSync(this.getConfigPath(), configFile, "utf8");
		} catch (e) {
			throw Errors.ConfigWriteFailed(e);
		}
	}

	private chainById(chainID: string): ChainConfig {
		return this.config.chains
			? this.config.chains.find((x) => x.chainId == chainID)
			: null;
	}

	private pathById(pathID: string): PathConfig {
		return this.config.paths
			? this.config.paths.find((x) => x.path.id == pathID)
			: null;
	}
	private async balanceCheck(chain: ChainConfig): Promise<boolean> {
		let chainBalances = await this.getAccountBalance([chain.chainId]);
		let chainGP = GasPrice.fromString(chain.gasPrice);
		if (!chainBalances.find((x) => x.denom == chainGP.denom)) return false;

		return !chainBalances.find(
			(x) =>
				x.denom == chainGP.denom &&
				parseInt(x.amount) < chainGP.amount.toFloatApproximation() * this.ibcSetupGas
		);
	}

	private notEnoughBalanceError(chain, gasPrice) {
		const { chainId } = chain;
		const { amount, denom } = gasPrice;
		const calcAmount = amount.toFloatApproximation() * this.ibcSetupGas;

		return Errors.NotEnoughBalance(`${calcAmount} ${denom} (${chainId})`);
	}

	private async createLink({
		path,
		options,
	}: PathConfig): Promise<LinkStatus> {
		let chainA = this.chainById(path.src.chainID);
		let chainB = this.chainById(path.dst.chainID);
		let chainAGP = GasPrice.fromString(chainA.gasPrice);
		let chainBGP = GasPrice.fromString(chainA.gasPrice);

		if (!(await this.balanceCheck(chainA)))
			return {
				status: false,
				pathName: path.id,
				error: this.notEnoughBalanceError(chainA, chainAGP).message,
			};

		if (!(await this.balanceCheck(chainB)))
			return {
				status: false,
				pathName: path.id,
				error: this.notEnoughBalanceError(chainB, chainBGP).message,
			};

		// create IBC clients.
		const clientA = await this.getIBCClient(chainA);
		const clientB = await this.getIBCClient(chainB);
		try {
			const link = await Link.createWithNewConnections(clientA, clientB);

			const channels = await link.createChannel(
				"A",
				options.sourcePort,
				options.targetPort,
				orderFromJSON(options.ordering),
				options.targetVersion
			);

			let configPath = this.pathById(path.id);
			configPath.path.src.channelID = channels.src.channelId;
			configPath.path.dst.channelID = channels.dest.channelId;
			configPath.path.isLinked = true;
			configPath.connections = {
				srcConnection: link.endA.connectionID,
				destConnection: link.endB.connectionID,
			};
			configPath.relayerData = null;

			return {
				status: true,
				pathName: path.id,
			};
		} catch (e) {
			return {
				status: false,
				pathName: path.id,
				error: e.toString(),
			};
		}
	}

	private async getLink({ path, connections }: PathConfig): Promise<Link> {
		let chainA = this.chainById(path.src.chainID);
		let chainB = this.chainById(path.dst.chainID);

		// create IBC clients.
		const clientA = await this.getIBCClient(chainA);
		const clientB = await this.getIBCClient(chainB);

		const link = Link.createWithExistingConnections(
			clientA,
			clientB,
			connections.srcConnection,
			connections.destConnection,
			new ConsoleLogger()
		);

		return link;
	}

	private async queryClient(rpcAddr: string): Promise<QueryClient> {
		return QueryClient.withExtensions(
			await Tendermint34Client.connect(rpcAddr),
      		setupBankExtension,
    	);
	}

	private async getIBCClient(chain: ChainConfig): Promise<IbcClient> {
		let chainGP = GasPrice.fromString(chain.gasPrice);
		let signer = await DirectSecp256k1Wallet.fromKey(
			this.config.mnemonic,
			{
				hdPaths: [stringToPath("m/44'/118'/0'/0/0")],
				prefix: chain.addressPrefix,
			}
		);

		const [account] = await signer.getAccounts();

		const client = await IbcClient.connectWithSigner(
			chain.rpcAddr,
			signer,
			account.address,
			{
				prefix: chain.addressPrefix,
				gasPrice: chainGP,
				gasLimits: calcGasLimits(chain.gasLimit),
			}
		);

		return client;
	}

	private async relayPackets(
		link,
		relayHeights,
		options = { maxAgeDest: this.defaultMaxAge, maxAgeSrc: this.defaultMaxAge }
	) {
		try {
			const heights = await link.checkAndRelayPacketsAndAcks(
				relayHeights ?? {},
				2,
				6
			);

			await link.updateClientIfStale("A", options.maxAgeDest);
			await link.updateClientIfStale("B", options.maxAgeSrc);

			return heights;
		} catch (e) {
			throw Errors.RelayPacketError;
		}
	}
}
