// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Person struct {
	pulumi.CustomResourceState

	Name pulumi.StringPtrOutput `pulumi:"name"`
	Pets PetTypeArrayOutput     `pulumi:"pets"`
}

// NewPerson registers a new resource with the given unique name, arguments, and options.
func NewPerson(ctx *pulumi.Context,
	name string, args *PersonArgs, opts ...pulumi.ResourceOption) (*Person, error) {
	if args == nil {
		args = &PersonArgs{}
	}

	var resource Person
	err := ctx.RegisterResource("example::Person", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPerson gets an existing Person resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPerson(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersonState, opts ...pulumi.ResourceOption) (*Person, error) {
	var resource Person
	err := ctx.ReadResource("example::Person", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Person resources.
type personState struct {
}

type PersonState struct {
}

func (PersonState) ElementType() reflect.Type {
	return reflect.TypeOf((*personState)(nil)).Elem()
}

type personArgs struct {
	Name *string   `pulumi:"name"`
	Pets []PetType `pulumi:"pets"`
}

// The set of arguments for constructing a Person resource.
type PersonArgs struct {
	Name pulumi.StringPtrInput
	Pets PetTypeArrayInput
}

func (PersonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*personArgs)(nil)).Elem()
}

type PersonInput interface {
	pulumi.Input

	ToPersonOutput() PersonOutput
	ToPersonOutputWithContext(ctx context.Context) PersonOutput
}

func (*Person) ElementType() reflect.Type {
	return reflect.TypeOf((**Person)(nil)).Elem()
}

func (i *Person) ToPersonOutput() PersonOutput {
	return i.ToPersonOutputWithContext(context.Background())
}

func (i *Person) ToPersonOutputWithContext(ctx context.Context) PersonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonOutput)
}

// PersonArrayInput is an input type that accepts PersonArray and PersonArrayOutput values.
// You can construct a concrete instance of `PersonArrayInput` via:
//
//          PersonArray{ PersonArgs{...} }
type PersonArrayInput interface {
	pulumi.Input

	ToPersonArrayOutput() PersonArrayOutput
	ToPersonArrayOutputWithContext(context.Context) PersonArrayOutput
}

type PersonArray []PersonInput

func (PersonArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Person)(nil)).Elem()
}

func (i PersonArray) ToPersonArrayOutput() PersonArrayOutput {
	return i.ToPersonArrayOutputWithContext(context.Background())
}

func (i PersonArray) ToPersonArrayOutputWithContext(ctx context.Context) PersonArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonArrayOutput)
}

// PersonMapInput is an input type that accepts PersonMap and PersonMapOutput values.
// You can construct a concrete instance of `PersonMapInput` via:
//
//          PersonMap{ "key": PersonArgs{...} }
type PersonMapInput interface {
	pulumi.Input

	ToPersonMapOutput() PersonMapOutput
	ToPersonMapOutputWithContext(context.Context) PersonMapOutput
}

type PersonMap map[string]PersonInput

func (PersonMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Person)(nil)).Elem()
}

func (i PersonMap) ToPersonMapOutput() PersonMapOutput {
	return i.ToPersonMapOutputWithContext(context.Background())
}

func (i PersonMap) ToPersonMapOutputWithContext(ctx context.Context) PersonMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonMapOutput)
}

type PersonOutput struct{ *pulumi.OutputState }

func (PersonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Person)(nil)).Elem()
}

func (o PersonOutput) ToPersonOutput() PersonOutput {
	return o
}

func (o PersonOutput) ToPersonOutputWithContext(ctx context.Context) PersonOutput {
	return o
}

type PersonArrayOutput struct{ *pulumi.OutputState }

func (PersonArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Person)(nil)).Elem()
}

func (o PersonArrayOutput) ToPersonArrayOutput() PersonArrayOutput {
	return o
}

func (o PersonArrayOutput) ToPersonArrayOutputWithContext(ctx context.Context) PersonArrayOutput {
	return o
}

func (o PersonArrayOutput) Index(i pulumi.IntInput) PersonOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Person {
		return vs[0].([]*Person)[vs[1].(int)]
	}).(PersonOutput)
}

type PersonMapOutput struct{ *pulumi.OutputState }

func (PersonMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Person)(nil)).Elem()
}

func (o PersonMapOutput) ToPersonMapOutput() PersonMapOutput {
	return o
}

func (o PersonMapOutput) ToPersonMapOutputWithContext(ctx context.Context) PersonMapOutput {
	return o
}

func (o PersonMapOutput) MapIndex(k pulumi.StringInput) PersonOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Person {
		return vs[0].(map[string]*Person)[vs[1].(string)]
	}).(PersonOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersonInput)(nil)).Elem(), &Person{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersonArrayInput)(nil)).Elem(), PersonArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersonMapInput)(nil)).Elem(), PersonMap{})
	pulumi.RegisterOutputType(PersonOutput{})
	pulumi.RegisterOutputType(PersonArrayOutput{})
	pulumi.RegisterOutputType(PersonMapOutput{})
}
