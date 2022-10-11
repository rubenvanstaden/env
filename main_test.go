package main_test

import (
	"testing"

	env "github.com/rubenvanstaden/env"
	"github.com/rubenvanstaden/test"
)

func TestUnit(t *testing.T) {

	t.Run("Grpc", func(t *testing.T) {

		v := env.RpcAddr("RPC_LOCAL_URL")
		test.Equals(t, "0.0.0.0:8080", v)

	})

	t.Run("Http", func(t *testing.T) {

		v := env.HttpAddr("HTTP_LOCAL_URL")
		test.Equals(t, "http://0.0.0.0:8081", v)

	})
}
