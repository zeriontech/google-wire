// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/zeriontech/google-wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"io"
	"strings"
)

// Injectors from wire.go:

func injectedReader() io.Reader {
	reader := _wireReaderValue
	return reader
}

var (
	_wireReaderValue = strings.NewReader("hello world")
)
