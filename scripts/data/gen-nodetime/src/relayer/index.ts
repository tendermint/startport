import run from "./jsonrpc";

import Relayer from "./lib/relayer";

const relayer = new Relayer();

run([
	["link", relayer.link.bind(relayer)],
	["start", relayer.start.bind(relayer)],
	["ensureChainSetup", relayer.ensureChainSetup.bind(relayer)],
	["createPath", relayer.createPath.bind(relayer)],
	["getPath", relayer.getPath.bind(relayer)],
	["listPaths", relayer.listPaths.bind(relayer)],
	["getAccountBalance", relayer.getAccountBalance.bind(relayer)],
	["info", relayer.info.bind(relayer)],
]);
