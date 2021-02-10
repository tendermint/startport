package ibc

import (
	"fmt"

	"github.com/tendermint/starport/starport/templates/typed"

	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
)

var (
	ibcTemplate = packr.New("ibc/templates/packet", "./packet")
)

// Options ...
type PacketOptions struct {
	AppName    string
	ModuleName string
	ModulePath string
	OwnerName  string
	PacketName string
	Fields     []typed.Field
}

// New ...
func NewIBC(opts *PacketOptions) (*genny.Generator, error) {
	g := genny.New()

	g.RunFn(moduleModify(opts))
	g.RunFn(protoModify(opts))
	g.RunFn(typeModify(opts))
	g.RunFn(eventModify(opts))

	// TODO: CODEC!!!

	if err := g.Box(ibcTemplate); err != nil {
		return g, err
	}
	ctx := plush.NewContext()
	ctx.Set("moduleName", opts.ModuleName)
	ctx.Set("modulePath", opts.ModulePath)
	ctx.Set("appName", opts.AppName)
	ctx.Set("packetName", opts.PacketName)
	ctx.Set("ownerName", opts.OwnerName)
	ctx.Set("fields", opts.Fields)
	ctx.Set("title", strings.Title)

	ctx.Set("nodash", func(s string) string {
		return strings.ReplaceAll(s, "-", "")
	})

	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Replace("{{moduleName}}", opts.ModuleName))
	g.Transformer(genny.Replace("{{packetName}}", opts.PacketName))
	return g, nil
}

func moduleModify(opts *PacketOptions) genny.RunFn {
	return func(r *genny.Runner) error {
		path := fmt.Sprintf("x/%s/module-ibc.go", opts.ModuleName)
		f, err := r.Disk.Find(path)
		if err != nil {
			return err
		}

		// Recv packet dispatch
		templateRecv := `%[1]v
case *types.%[2]vPacketData_%[3]vPacket:
	err := am.keeper.OnRecv%[3]vPacket(ctx, modulePacket, packet.%[3]vPacket)
	if err != nil {
		acknowledgement = channeltypes.NewErrorAcknowledgement(err.Error())
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventType%[3]vPacket,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
		),
	)`
		replacementRecv := fmt.Sprintf(
			templateRecv,
			PlaceholderIBCPacketModuleRecv,
			strings.Title(opts.ModuleName),
			strings.Title(opts.PacketName),
		)
		content := strings.Replace(f.String(), PlaceholderIBCPacketModuleRecv, replacementRecv, 1)

		// Ack packet dispatch
		templateAck := `%[1]v
case *types.%[2]vPacketData_%[3]vPacket:
	err := am.keeper.OnAcknowledgement%[3]vPacket(ctx, modulePacket, packet.%[3]vPacket, ack)
	if err != nil {
		return nil, err
	}
	eventType = types.EventType%[3]vPacket`
		replacementAck := fmt.Sprintf(
			templateAck,
			PlaceholderIBCPacketModuleAck,
			strings.Title(opts.ModuleName),
			strings.Title(opts.PacketName),
		)
		content = strings.Replace(content, PlaceholderIBCPacketModuleAck, replacementAck, 1)

		// Timeout packet dispatch
		templateTimeout := `%[1]v
case *types.%[2]vPacketData_%[3]vPacket:
	err := am.keeper.OnTimeoutPacket%[3]vPacket(ctx, modulePacket, packet.%[3]vPacket)
	if err != nil {
		return nil, err
	}`
		replacementTimeout := fmt.Sprintf(
			templateTimeout,
			PlaceholderIBCPacketModuleTimeout,
			strings.Title(opts.ModuleName),
			strings.Title(opts.PacketName),
		)
		content = strings.Replace(content, PlaceholderIBCPacketModuleTimeout, replacementTimeout, 1)

		newFile := genny.NewFileS(path, content)
		return r.File(newFile)
	}
}

func protoModify(opts *PacketOptions) genny.RunFn {
	return func(r *genny.Runner) error {
		path := fmt.Sprintf("proto/%s/packet.proto", opts.ModuleName)
		f, err := r.Disk.Find(path)
		if err != nil {
			return err
		}

		content := f.String()

		// Add the field in the module packet
		fieldCount := strings.Count(content, PlaceholderIBCPacketProtoFieldNumber)
		templateField := `%[1]v
		%[2]vPacketData %[3]vpacket = %[4]v; %[5]v`
		replacementField := fmt.Sprintf(
			templateField,
			PlaceholderIBCPacketProtoField,
			strings.Title(opts.PacketName),
			opts.PacketName,
			fieldCount + 2,
			PlaceholderIBCPacketProtoFieldNumber,
		)
		content = strings.Replace(content, PlaceholderIBCPacketProtoField, replacementField, 1)

		// Add the message definition
		var messageFields string
		for i, field := range opts.Fields {
			messageFields += fmt.Sprintf("  %s %s = %d;\n", field.Datatype, field.Name, i+1)
		}
		templateMessage := `%[1]v
// %[2]vPacketData defines a struct for the packet payload
message %[2]vPacketData {
	%[3]v
}
`
		replacementMessage := fmt.Sprintf(
			templateMessage,
			PlaceholderIBCPacketProtoMessage,
			strings.Title(opts.PacketName),
			messageFields,
			)
		content = strings.Replace(content, PlaceholderIBCPacketProtoMessage, replacementMessage, 1)

		newFile := genny.NewFileS(path, content)
		return r.File(newFile)
	}
}

func typeModify(opts *PacketOptions) genny.RunFn {
	return func(r *genny.Runner) error {
		path := fmt.Sprintf("x/%s/types/packet.go", opts.ModuleName)
		f, err := r.Disk.Find(path)
		if err != nil {
			return err
		}

		template := `%[1]v
// ValidateBasic is used for validating the packet
func (p %[3]vPacketData) ValidateBasic() error {
	
	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p %[3]vPacketData) GetBytes() []byte {
	var modulePacket %[2]vPacketData

	modulePacket.Packet = &%[2]vPacketData_%[3]vPacket{p}

	return ModuleCdc.MustMarshalBinaryBare(&p)
}`
		replacement := fmt.Sprintf(
			template,
			PlaceholderIBCPacketType,
			strings.Title(opts.ModuleName),
			strings.Title(opts.PacketName),
			)
		content := strings.Replace(f.String(), PlaceholderIBCPacketType, replacement, 1)

		newFile := genny.NewFileS(path, content)
		return r.File(newFile)
	}
}

func eventModify(opts *PacketOptions) genny.RunFn {
	return func(r *genny.Runner) error {
		path := fmt.Sprintf("x/%s/types/events_ibc.go", opts.ModuleName)
		f, err := r.Disk.Find(path)
		if err != nil {
			return err
		}

		template := `%[1]v
EventType%[2]vPacket       = "%[3]v_packet"
`
		replacement := fmt.Sprintf(
			template,
			PlaceholderIBCPacketEvent,
			strings.Title(opts.PacketName),
			opts.PacketName,
			)
		content := strings.Replace(f.String(), PlaceholderIBCPacketEvent, replacement, 1)

		newFile := genny.NewFileS(path, content)
		return r.File(newFile)
	}
}
