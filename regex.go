package env

import (
	"fmt"
)

func rpcAddress() string {
	domain := `(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	port := `(\d{4})`
	return fmt.Sprintf("^%s:%s$", domain, port)
}

func httpAddress() string {
	domain := `(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	port := `(\d{4})`
	return fmt.Sprintf("^http://%s:%s$", domain, port)
}
