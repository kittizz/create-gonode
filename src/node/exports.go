//go:build js

package main

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
	randstr "github.com/kittizz/create-gonode/src/go"
)

var (
	global   = js.Global
	exports  = js.Module.Get("exports")
	errorObj = global.Get("Error")
)

func throwError(err error) {
	panic(errorObj.New(err.Error()))
}

func randomString(length int) string {
	if length < 1 {
		throwError(errors.New("length must be greater than 0"))
	}
	return randstr.RandomString(length)
}
func main() {
	exports.Set("randomString", randomString)
}
