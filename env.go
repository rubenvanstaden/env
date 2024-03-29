package env

import (
	"log"
	"os"
	"regexp"
)

func String(key string) string {

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("address env variable \"%s\" not set, usual", key)
	}

	return value
}

func HttpAddr(key string) string {

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Regex env variable \"%s\" not set, usual", key)
	}

	re := regexp.MustCompile(httpRegex())
	match := re.MatchString(value)

	if !match {
		log.Fatalf("Invalid HTTP address env variable \"%s\"", key)
	}

	return value
}

func WebsocketAddr(key string) string {

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Regex env variable \"%s\" not set, usual", key)
	}

	re := regexp.MustCompile(websocketRegex())
	match := re.MatchString(value)

	if !match {
		log.Fatalf("Invalid WS address env variable \"%s\"", key)
	}

	return value
}

func RpcAddr(key string) string {

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Regex env variable \"%s\" not set, usual", key)
	}

	re := regexp.MustCompile(grpcRegex())
	match := re.MatchString(value)

	if !match {
		log.Fatalf("Invalid RPC address env variable \"%s\"", key)
	}

	return value
}

func MongoAddr(key string) string {

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Regex env variable \"%s\" not set, usual", key)
	}

	re := regexp.MustCompile(mongoRegex())
	match := re.MatchString(value)

	if !match {
		log.Fatalf("Invalid WS address env variable \"%s\"", key)
	}

	return value
}
