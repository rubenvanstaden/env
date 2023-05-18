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
