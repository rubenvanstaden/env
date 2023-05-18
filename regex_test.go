package env

import (
	"regexp"
	"testing"

	"github.com/rubenvanstaden/test"
)

func TestUnit_Http(t *testing.T) {

	cases := []struct {
        name string
        addr string
        want bool
	}{
		{
            name: "Valid Http",
            addr: "http://127.0.0.1:8080",
            want: true,
		},
		{
            name: "Invalid Prefix",
            addr: "htt://127.0.0.1:8080",
            want: false,
		},
		{
            name: "Invalid Domain",
            addr: "htt://127.0.01:8080",
            want: false,
		},
		{
            name: "Invalid Port",
            addr: "htt://127.0.01:808",
            want: false,
		},
	}

	for _, c := range cases {
        re := regexp.MustCompile(httpRegex())
        got := re.MatchString(c.addr)
        test.Equals(t, c.want, got)
	}
}

func TestUnit_Websocket(t *testing.T) {

	cases := []struct {
        name string
        addr string
        want bool
	}{
		{
            name: "Valid Websocket",
            addr: "wss://nostr.wine",
            want: true,
		},
		{
            name: "Valid Prefix",
            addr: "ws://nostr.wine",
            want: true,
		},
		{
            name: "Invalid Prefix",
            addr: "w://nostr.wine",
            want: false,
		},
		{
            name: "Invalid URL",
            addr: "wss:/nostr.wine",
            want: false,
		},
	}

	for _, c := range cases {
        re := regexp.MustCompile(websocketRegex())
        got := re.MatchString(c.addr)
        test.Equals(t, c.want, got)
	}
}

func TestUnit_Mongo(t *testing.T) {

	cases := []struct {
        name string
        addr string
        want bool
	}{
		{
            name: "Valid address",
            addr: "mongodb://admin:password@mongodb:27017",
            want: true,
		},
		{
            name: "Valid username and password",
            addr: "mongodb://mongodb:27017",
            want: true,
		},
		{
            name: "Valid port size 4",
            addr: "mongodb://admin:password@mongodb:2701",
            want: true,
		},
		{
            name: "Valid port size 6",
            addr: "mongodb://admin:password@mongodb:270178",
            want: true,
		},
		{
            name: "Invalid username and password",
            addr: "mongodb://:@mongodb:27017",
            want: false,
		},
		{
            name: "Invalid Prefix",
            addr: "mongo://admin:password@mongodb:27017",
            want: false,
		},
	}

	for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            re := regexp.MustCompile(mongoRegex())
            got := re.MatchString(c.addr)
            test.Equals(t, c.want, got)
        })
	}
}
