// Wrap proto structs to allow easier creation, protobuf lang is small enough
// to easily allow this.package protoutil
package protoutil

import (
	"fmt"
	"strconv"

	"github.com/emicklei/proto"
)

// TODO: Can also support comments/inline comments? -- Probably, formatting is currently
// flaky with how it prints them, though.

// Create a new Literal:
//
// // true
// l := NewLiteral("true")
//
// // 1
// l := NewLiteral("1")
//
// // "foo"
// l := NewLiteral("foo")
//
// Currently doesn't support creating compound literals (arrays/maps).
func NewLiteral(lit string) *proto.Literal {
	return &proto.Literal{
		Source:   lit,
		IsString: isString(lit),
	}
}

// ImportSpec holds information relevant to the import statement.
type ImportSpec struct {
	path string
	kind string
}

// Type alias for a callable accepting an ImportSpec.
type ImportSpecOpts func(i *ImportSpec)

// Weak allows you to set the kind of the import statement to 'weak'.
func Weak() ImportSpecOpts {
	return func(i *ImportSpec) {
		i.kind = "weak"
	}
}

// Public allows you to set the kind of the import statement to 'public'.
func Public() ImportSpecOpts {
	return func(i *ImportSpec) {
		i.kind = "public"
	}
}

// NewImport creates a new import statement node:
//
//		// import "myproto.proto";
//	 imp := NewImport("myproto.proto")
//
// By default, no kind is assigned to it, by using Weak or Public, this can be specified:
//
//	// import weak "myproto.proto";
//	imp := NewImport("myproto.proto", Weak())
func NewImport(path string, opts ...ImportSpecOpts) *proto.Import {
	i := ImportSpec{path: path}
	for _, opt := range opts {
		opt(&i)
	}

	return &proto.Import{
		Filename: i.path,
		Kind:     i.kind,
	}
}

// NewPackage creates a new package statement node:
//
//	// package foo.bar;
//	pkg := NewPackage("foo.bar")
func NewPackage(path string) *proto.Package {
	return &proto.Package{
		Name: path,
	}
}

// OptionSpec holds information relevant to the option statement.
type OptionSpec struct {
	name     string
	setter   string
	constant string
	custom   bool
}

// OptionSpecOpts is a function that accepts an OptionSpec.
type OptionSpecOpts func(o *OptionSpec)

// Custom denotes the option as being a custom option.
func Custom() OptionSpecOpts {
	return func(f *OptionSpec) {
		f.custom = true
	}
}

// Setter allows setting specific fields for a given option
// that denotes a type with fields.
//
//	// option (my_opt).field = "Value";
//	opt := NewOption("my_opt", "Value", Custom(), Setter("field"))
func SetField(name string) OptionSpecOpts {
	return func(f *OptionSpec) {
		f.setter = name
	}
}

// NewOption creates a new option statement node:
//
//	// option foo = 1;
//	opt := NewOption("foo", "1")
//
// Custom options can be marked as such by using Custom, this wraps the option name
// in parenthesis:
//
//	// option (foo) = 1;
//	opt := NewOption("foo", "1", Custom())
//
// Since option constants can accept a number of types, strings that require quotation
// should be passed as raw strings:
//
//	// option foo = "bar";
//	opt := NewOption("foo", `bar`)
func NewOption(name, constant string, opts ...OptionSpecOpts) *proto.Option {
	o := OptionSpec{name: name, constant: constant}
	for _, opt := range opts {
		opt(&o)
	}
	if o.custom {
		o.name = fmt.Sprintf("(%s)", o.name)
	}
	// add the field we are setting outside the parentheses.
	if o.setter != "" {
		o.name = fmt.Sprintf("%s.%s", o.name, o.setter)
	}
	return &proto.Option{
		Name:     o.name,
		Constant: *NewLiteral(o.constant),
	}
}

/// Service + PRC

// RPCSpec holds information relevant to the rpc statement.
type RPCSpec struct {
	name, inputType, outputType string
	streamsReq, streamsResp     bool
	options                     []*proto.Option
}

// Type alias for a callable accepting an RPCSpec.
type RPCSpecOpts func(i *RPCSpec)

// Mark request as streaming.
func StreamRequest() RPCSpecOpts {
	return func(r *RPCSpec) {
		r.streamsReq = true
	}
}

// Mark response as streaming.
func StreamResponse() RPCSpecOpts {
	return func(r *RPCSpec) {
		r.streamsResp = true
	}
}

// WithRPCOptions adds options to the RPC.
func WithRPCOptions(option ...*proto.Option) RPCSpecOpts {
	return func(o *RPCSpec) {
		o.options = append(o.options, option...)
	}
}

// NewRPC creates a new RPC statement node:
//
//	// rpc Foo(Bar) returns(Bar) {}
//	rpc := NewRPC("Foo", "Bar", "Bar")
//
// No options are attached by default, use WithRPCOptions to add options as required:
//
//	// rpc Foo(Bar) returns(Bar) {
//	//  option (foo) = 1;
//	// }
//	rpc := NewRPC("Foo", "Bar", "Bar", WithRPCOptions(NewOption("foo", "1")))
func NewRPC(name string, inputType string, outputType string, opts ...RPCSpecOpts) *proto.RPC {
	r := RPCSpec{name: name, inputType: inputType, outputType: outputType}
	for _, opt := range opts {
		opt(&r)
	}

	rpc := &proto.RPC{
		Name:           r.name,
		RequestType:    r.inputType,
		ReturnsType:    r.outputType,
		StreamsRequest: r.streamsReq,
		StreamsReturns: r.streamsResp,
	}
	if len(r.options) > 0 {
		for _, opt := range r.options {
			rpc.Elements = append(rpc.Elements, opt)
		}
	}
	return rpc
}

// ServiceSpec holds information relevant to the service statement.
type ServiceSpec struct {
	name string
	rpcs []*proto.RPC
	opts []*proto.Option
}

// ServiceSpecOpts is a type alias for a callable accepting a ServiceSpec.
type ServiceSpecOpts func(i *ServiceSpec)

// WithRPCs adds rpcs to the service.
func WithRPCs(rpcs ...*proto.RPC) ServiceSpecOpts {
	return func(s *ServiceSpec) {
		s.rpcs = append(s.rpcs, rpcs...)
	}
}

// WithServiceOptions adds options to the service.
func WithServiceOptions(options ...*proto.Option) ServiceSpecOpts {
	return func(s *ServiceSpec) {
		s.opts = append(s.opts, options...)
	}
}

// NewService creates a new service statement node:
//
//	// service Foo {}
//	service := NewService("Foo")
//
// No rpcs/options are attached by default, use WithRPCs and
// WithServiceOptions to add them as required:
//
//	 // service Foo {
//	 //  option (foo) = 1;
//	 //  rpc Bar(Bar) returns (Bar) {}
//	 // }
//		opt := NewOption("foo", "1")
//	 rpc := NewRPC("Bar", "Bar", "Bar")
//	 service := NewService("Foo", WithServiceOptions(opt), WithRPCs(rpc))
//
// By default, options are added first and then the rpcs.
func NewService(name string, opts ...ServiceSpecOpts) *proto.Service {
	s := ServiceSpec{name: name}
	for _, opt := range opts {
		opt(&s)
	}
	service := &proto.Service{
		Name: s.name,
	}
	for _, opt := range s.opts {
		service.Elements = append(service.Elements, opt)
	}
	for _, rpc := range s.rpcs {
		service.Elements = append(service.Elements, rpc)
	}
	return service
}

/// Message + NormalField

// FieldSpec holds information relevant to the field statement.
type FieldSpec struct {
	name, typ                    string
	sequence                     int
	repeated, optional, required bool
	options                      []*proto.Option
}

// FieldSpecOpts is a type alias for a callable accepting a FieldSpec.
type FieldSpecOpts func(f *FieldSpec)

// Repeated marks the field as repeated.
func Repeated() FieldSpecOpts {
	return func(f *FieldSpec) {
		f.repeated = true
	}
}

// Optional marks the field as optional.
func Optional() FieldSpecOpts {
	return func(f *FieldSpec) {
		f.optional = true
	}
}

// Required marks the field as required.
func Required() FieldSpecOpts {
	return func(f *FieldSpec) {
		f.required = true
	}
}

// WithFieldOptions adds options to the field.
func WithFieldOptions(options ...*proto.Option) FieldSpecOpts {
	return func(f *FieldSpec) {
		f.options = append(f.options, options...)
	}
}

// NewField creates a new field statement node:
//
//	// int32 Foo = 1;
//	field := NewField("Foo", "int32", 1)
//
// Fields aren't marked as repeated, required or optional. Use Repeated, Optional
// and Required to mark the field as such.
//
//	// repeated int32 Foo = 1;
//	field := NewField("Foo", "int32", 1, Repeated())
func NewField(typ, name string, sequence int, opts ...FieldSpecOpts) *proto.NormalField {
	f := FieldSpec{name: name, typ: typ, sequence: sequence}
	for _, opt := range opts {
		opt(&f)
	}

	// Check qualifiers? Though protoc will shout if we do stupid things.
	field := &proto.NormalField{
		Field: &proto.Field{
			Name:     f.name,
			Sequence: f.sequence,
			Type:     f.typ,
			Options:  []*proto.Option{},
		},
		Repeated: f.repeated,
		Required: f.required,
		Optional: f.optional,
	}
	if len(f.options) > 0 {
		field.Options = append(field.Options, f.options...)
	}
	return field
}

// MessageSpec holds information relevant to the message statement.
type MessageSpec struct {
	name     string
	fields   []*proto.NormalField // needs a good amount of work.
	enums    []*proto.Enum
	options  []*proto.Option
	isExtend bool
}

// MessageSpecOpts is a type alias for a callable accepting a MessageSpec.
type MessageSpecOpts func(i *MessageSpec)

// WithMessageOptions adds options to the message.
func WithMessageOptions(options ...*proto.Option) MessageSpecOpts {
	return func(m *MessageSpec) {
		m.options = append(m.options, options...)
	}
}

// WithMessageFields adds fields to the message.
func WithFields(fields ...*proto.NormalField) MessageSpecOpts {
	return func(m *MessageSpec) {
		m.fields = append(m.fields, fields...)
	}
}

// WithEnums adds enums to the message.
func WithEnums(enum ...*proto.Enum) MessageSpecOpts {
	return func(m *MessageSpec) {
		m.enums = append(m.enums, enum...)
	}
}

func Extend() MessageSpecOpts {
	return func(m *MessageSpec) {
		m.isExtend = true
	}
}

// NewMessage creates a new message statement node:
//
//	// message Foo {}
//	message := NewMessage("Foo")
//
// No fields/enums/options are attached by default, use WithMessageFields, WithEnums,
// and WithMessageOptions to add them as required:
//
//	 // message Foo {
//	 //  option (foo) = 1;
//	 //  int32 Bar = 1;
//	 // }
//		opt := NewOption("foo", "1")
//	 field := NewField("int32", "Bar", 1)
//	 message := NewMessage("Foo", WithMessageOptions(opt), WithFields(field))
//
// By default, options are added first, then fields and then enums.
func NewMessage(name string, opts ...MessageSpecOpts) *proto.Message {
	m := MessageSpec{name: name}
	for _, opt := range opts {
		opt(&m)
	}
	message := &proto.Message{
		Name:     m.name,
		IsExtend: m.isExtend,
	}
	for _, opt := range m.options {
		message.Elements = append(message.Elements, opt)
	}

	// Verify that fields have unique sequence? Though, again, protoc will shout if
	// it isn't the case.
	for _, field := range m.fields {
		message.Elements = append(message.Elements, field)
	}
	for _, enum := range m.enums {
		message.Elements = append(message.Elements, enum)
	}
	return message
}

// EnumFieldSpec holds information relevant to the enum field statement.
type EnumFieldSpec struct {
	name    string
	value   int
	options []*proto.Option
}

// EnumFieldSpecOpts is a type alias for a callable accepting an EnumFieldSpec.
type EnumFieldSpecOpts func(f *EnumFieldSpec)

// WithEnumFieldOptions adds options to the enum field.
func WithEnumFieldOptions(options ...*proto.Option) EnumFieldSpecOpts {
	return func(f *EnumFieldSpec) {
		f.options = append(f.options, options...)
	}
}

// NewEnumField creates a new enum field statement node:
//
//	// BAR = 1;
//	field := NewEnumField("BAR", 1)
//
// No options are attached by default, use WithEnumFieldOptions to add them as
// required:
//
//	// BAR = 1 [option (foo) = 1];
//	field := NewEnumField("BAR", 1, WithEnumFieldOptions(NewOption("foo", "1")))
func NewEnumField(name string, value int, opts ...EnumFieldSpecOpts) *proto.EnumField {
	f := EnumFieldSpec{name: name, value: value}
	for _, opt := range opts {
		opt(&f)
	}

	field := &proto.EnumField{
		Name:    f.name,
		Integer: f.value,
	}
	for _, opt := range f.options {
		field.Elements = append(field.Elements, opt)
	}
	return field
}

// EnumSpec holds information relevant to the enum statement.
type EnumSpec struct {
	name    string
	fields  []*proto.EnumField
	options []*proto.Option
}

// EnumSpecOpts is a type alias for a callable accepting an EnumSpec.
type EnumSpecOpts func(i *EnumSpec)

// WithEnumOptions adds options to the enum.
func WithEnumOptions(options ...*proto.Option) EnumSpecOpts {
	return func(e *EnumSpec) {
		e.options = append(e.options, options...)
	}
}

// WithEnumFields adds fields to the enum.
func WithEnumFields(fields ...*proto.EnumField) EnumSpecOpts {
	return func(e *EnumSpec) {
		e.fields = append(e.fields, fields...)
	}
}

// NewEnum creates a new enum statement node:
//
//	// enum Foo {
//	//  BAR = 1;
//	// }
//	enum := NewEnum("Foo", WithEnumFields(NewEnumField("BAR", 1)))
//
// No options are attached by default, use WithEnumOptions to add them as
// required:
//
//	// enum Foo {
//	//  BAR = 1 [option (foo) = 1];
//	// }
//	enum := NewEnum("Foo", WithEnumOptions(NewOption("foo", "1")), WithEnumFields(NewEnumField("BAR", 1)))
//
// By default, options are added first, then fields.
func NewEnum(name string, opts ...EnumSpecOpts) *proto.Enum {
	e := EnumSpec{name: name}
	for _, opt := range opts {
		opt(&e)
	}
	enum := &proto.Enum{
		Name: e.name,
	}
	for _, opt := range e.options {
		enum.Elements = append(enum.Elements, opt)
	}
	for _, field := range e.fields {
		enum.Elements = append(enum.Elements, field)
	}
	return enum
}

// OneOfField holds information relevant to the oneof field statement.
type OneOfFieldSpec struct {
	name, typ string
	sequence  int
	options   []*proto.Option
}

// OneOfFieldOpts is a type alias for a callable accepting a OneOfField.
type OneOfFieldOpts func(f *OneOfFieldSpec)

// WithOneOfFieldOptions adds options to the oneof field.
func WithOneOfFieldOptions(options ...*proto.Option) OneOfFieldOpts {
	return func(f *OneOfFieldSpec) {
		f.options = append(f.options, options...)
	}
}

// NewOneOfField creates a new oneof field statement node:
//
//		// Needs to placed in a oneof block.
//	 // int32 Foo = 1;
//	 field := NewOneOfField("Foo", "int32", 1)
//
// Additional options can be created and attached to the field to the field via
// WithOneOfFieldOptions:
//
//	// int32 Foo = 1 [option (foo) = 1];
//	field := NewOneOfField("Foo", "int32", 1, WithOneOfFieldOptions(NewOption("foo", "1")))
func NewOneOfField(typ, name string, sequence int, opts ...OneOfFieldOpts) *proto.OneOfField {
	f := OneOfFieldSpec{name: name, typ: typ, sequence: sequence}
	for _, opt := range opts {
		opt(&f)
	}
	field := &proto.OneOfField{
		Field: &proto.Field{
			Name:     f.name,
			Sequence: f.sequence,
			Type:     f.typ,
			Options:  []*proto.Option{},
		},
	}
	field.Options = append(field.Options, f.options...)
	return field
}

// OneOfSpec holds information relevant to the enum statement.
type OneofSpec struct {
	name    string
	options []*proto.Option
	fields  []*proto.OneOfField
}

// OneOfSpecOpts is a type alias for a callable accepting a OneOfSpec.
type OneOfSpecOpts func(o *OneofSpec)

// WithOneOfOptions adds options to the oneof.
func WithOneOfOptions(options ...*proto.Option) OneOfSpecOpts {
	return func(o *OneofSpec) {
		o.options = append(o.options, options...)
	}
}

// WithOneOfFields adds fields to the oneof.
func WithOneOfFields(fields ...*proto.OneOfField) OneOfSpecOpts {
	return func(o *OneofSpec) {
		o.fields = append(o.fields, fields...)
	}
}

// NewOneOf creates a new oneof statement node:
//
//	// oneof Foo {
//	//  int32 Foo = 1;
//	// }
//	oneof := NewOneOf("Foo", WithOneOfFields(NewOneOfField("Foo", "int32", 1)))
//
// No options are attached by default, use WithOneOfOptions to add them as required.
func NewOneOf(name string, opts ...OneOfSpecOpts) *proto.Oneof {
	o := OneofSpec{name: name}
	for _, opt := range opts {
		opt(&o)
	}
	oneof := &proto.Oneof{
		Name: o.name,
	}
	for _, opt := range o.options {
		oneof.Elements = append(oneof.Elements, opt)
	}
	for _, field := range o.fields {
		oneof.Elements = append(oneof.Elements, field)
	}
	return oneof
}

// Handle this better. Currently s with at least one digit is
// considered a number and if not, a string.
func isString(s string) bool {
	if s == "true" || s == "false" {
		return false
	}
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return false
	}
	return true
}
