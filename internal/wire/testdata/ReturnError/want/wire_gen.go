// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/zeriontech/google-wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFoo() (Foo, error) {
	foo, err := provideFoo()
	if err != nil {
		return 0, err
	}
	return foo, nil
}
