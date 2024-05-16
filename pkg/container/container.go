package container

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

const (
	msgErrNotDetectType       = "cannot detect type of the receiver, make sure your are passing reference of the object"
	msgErrNotFound            = "no concrete found for the abstraction "
	msgErrInvalidReceiver     = "the receiver must be either a reference or a callback"
	msgErrInvalidReceiverType = "container: invalid structure"
)

type binding struct {
	resolver interface{} // resolver function
	instance interface{} // instance stored for singleton bindings
}

// arguments will return resolved arguments of the given function.
func arguments(function interface{}) []reflect.Value {
	functionTypeOf := reflect.TypeOf(function)
	argumentsCount := functionTypeOf.NumIn()
	arguments := make([]reflect.Value, argumentsCount)

	for i := 0; i < argumentsCount; i++ {
		abstraction := functionTypeOf.In(i)

		var instance interface{}

		if concrete, ok := container[abstraction]; ok {
			instance = concrete.resolve()
		} else {
			panic(msgErrNotFound + abstraction.String())
		}

		arguments[i] = reflect.ValueOf(instance)
	}

	return arguments
}

// invoke will call the given function and return its returned value.
// It only works for functions that return a single value.
func invoke(function interface{}) interface{} {
	return reflect.ValueOf(function).Call(arguments(function))[0].Interface()
}

// resolve will return the concrete of related abstraction.
func (b binding) resolve() interface{} {
	if b.instance != nil {
		return b.instance
	}

	return invoke(b.resolver)
}

// container is the IoC container that will keep all of the bindings.
var container = map[reflect.Type]binding{}

// bind will map an abstraction to a concrete and set instance if it's a singleton binding.
func bind(resolver interface{}, singleton bool) {
	resolverTypeOf := reflect.TypeOf(resolver)
	if resolverTypeOf.Kind() != reflect.Func {
		panic("the resolver must be a function")
	}

	for i := 0; i < resolverTypeOf.NumOut(); i++ {
		var instance interface{}
		if singleton {
			instance = invoke(resolver)
		}

		container[resolverTypeOf.Out(i)] = binding{
			resolver: resolver,
			instance: instance,
		}
	}
}

// Register will bind an abstraction to a concrete for further singleton resolves.
// It takes a resolver function which returns the concrete and its return type matches the abstraction (interface).
// The resolver function can have arguments of abstraction that have bound already in Container.
func Register[T any](instance T) {
	resolverTypeOf := reflect.TypeOf(instance)
	if resolverTypeOf.Kind() == reflect.Func {
		bind(instance, true)
	}

	tmp := new(T)
	receiverTypeOf := reflect.TypeOf(tmp)

	container[receiverTypeOf.Elem()] = binding{
		resolver: nil,
		instance: instance,
	}
}

// ResolverMust will resolve the dependency and return a appropriate concrete of the given abstraction.
// It can take an abstraction (interface reference) and fill it with the related implementation.
// It also can take a function (receiver) with one or more arguments of the abstractions (interfaces) that need to be
// resolved, Container will invoke the receiver function and pass the related implementations.
func ResolverMust[T any]() (receiver T) {
	receiver, err := Resolver[T]()
	if err != nil {
		panic(err)
	}
	return receiver
}

func Resolver[T any]() (receiver T, err error) {
	tmp := new(T)
	receiverTypeOf := reflect.TypeOf(tmp)
	if receiverTypeOf == nil {
		return receiver, errors.New(msgErrNotDetectType)
	}

	if receiverTypeOf.Kind() == reflect.Ptr {
		abstraction := receiverTypeOf.Elem()
		if concrete, ok := container[abstraction]; ok {
			instance := concrete.resolve()
			aaa := reflect.New(abstraction).Elem()
			aaa.Set(reflect.ValueOf(instance))
			return (aaa.Interface()).(T), nil
		}

		return receiver, errors.New(msgErrNotFound + abstraction.String())
	}

	return receiver, errors.New(msgErrInvalidReceiver)
}

// Reset the container, remove all the bindings
func Reset() {
	container = map[reflect.Type]binding{}
}

// Fill takes a struct and fills the fields with the tag `di:"inject"`
func Fill(structure interface{}) {
	receiverType := reflect.TypeOf(structure)
	if receiverType == nil {
		panic(msgErrInvalidReceiverType)
	}

	if receiverType.Kind() != reflect.Ptr {
		panic(msgErrInvalidReceiverType)
	}
	elem := receiverType.Elem()
	if elem.Kind() != reflect.Struct {
		panic(msgErrInvalidReceiverType)
	}
	s := reflect.ValueOf(structure).Elem()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if t, ok := s.Type().Field(i).Tag.Lookup("di"); ok && t == "inject" {
			if concrete, ok := container[f.Type()]; ok {
				instance := concrete.resolve()

				ptr := reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
				ptr.Set(reflect.ValueOf(instance))
				continue
			}

			panic(fmt.Sprintf("container: cannot resolve %v has type: %v", s.Type().Field(i).Name, s.Field(i).Type()))
		}
	}

	return
}
