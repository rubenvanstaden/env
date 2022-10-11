package main

import (
	"log"
	"os"
	"regexp"
)

func RpcAddr(key string) string {

    value, ok := os.LookupEnv(key)
	if !ok {
        log.Fatalf("address env variable \"%s\" not set, usual", key)
	}

	re := regexp.MustCompile(rpcAddress())
    match := re.MatchString(value)

    if !match {
        log.Fatalf("invalid RPC address env variable \"%s\"", key)
    }

	return value
}

func HttpAddr(key string) string {

    value, ok := os.LookupEnv(key)
	if !ok {
        log.Fatalf("address env variable \"%s\" not set, usual", key)
	}

    re := regexp.MustCompile(httpAddress())
    match := re.MatchString(value)

    if !match {
        log.Fatalf("invalid HTTP address env variable \"%s\"", key)
    }

	return value
}

func Test(key string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return ""
}
