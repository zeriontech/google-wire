// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/zeriontech/google-wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooer() Fooer {
	bar := provideBar()
	return bar
}
