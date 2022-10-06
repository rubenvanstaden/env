package main

import (
	"log"
	"os"
)

func HasValue(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

func RpcAddr(key string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Fatalf("address env variable \"%s\" not set, usual", key)

	return ""
}

func HttpAddr(key string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Fatalf("address env variable \"%s\" not set, usual", key)

	return ""
}

func Test(key string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return ""
}
