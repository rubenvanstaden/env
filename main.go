package main

import (
	"fmt"
	"regexp"
)

func ipAddress() string {
    domain := `(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
    port := `:(\d{4})`
    return domain + port
}

func main() {
	re := regexp.MustCompile(ipAddress())
    match := re.MatchString("http://000.12.12.034:0000")
	fmt.Println(match)
}
