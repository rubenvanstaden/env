package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func main() {

	match, _ := regexp.MatchString(`(\d{1,3})\.(\d)\.(\d)\.(\d):(\d{1,4})`, "123450.0.0.0:8080")
	fmt.Println(match)

}
