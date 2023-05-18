package env

func httpRegex() string {
	// Define a simple regex pattern for an http or https URL.
	// Starts with http:// or https://
	// Followed by: 
	//   - One or more alphanumeric characters, hyphens or periods (the domain name and possibly subdomains)
	//   - Optional :port number (one or more digits)
	//   - The remainder of the URL, which can include alphanumeric characters, hyphens, periods, slashes, and a few other valid URL characters
	//return `^(http://|https://)?([a-zA-Z0-9-]+\.)*[a-zA-Z0-9][a-zA-Z0-9-]+\.[a-zA-Z]{2,6}?(:[0-9]{1,5})?(\/.*)?$`
    return `^(http|https):\/\/[^\s/:]+(:[0-9]+)?$`
}

func websocketRegex() string {

    // scheme matches the optional ws:// or wss:// at the beginning.
    scheme := `^(ws://|wss://)?`
    // subdomains matches any number of optional subdomains, each ending with a dot.
    subdomains := `([a-zA-Z0-9-]+\.)*`
    // domain matches the domain name, which must contain at least one character before the dot, followed by a top-level domain of 2 to 6 characters.
    domain := `[a-zA-Z0-9][a-zA-Z0-9-]+\.[a-zA-Z]{2,6}?`
    // port matches an optional :port part, where the port is 1 to 5 digits.
    port := `(:[0-9]{1,5})?`
    // path matches the remainder of the URL, which starts with a slash and can contain any characters.
    path := `(\/.*)?$`

    return scheme + subdomains + domain + port + path
}

func grpcRegex() string {
	// Define a simple regex pattern for a gRPC address.
	// It can be a domain name or an IP address, optionally followed by a port number.
	// For domain names:
	//   - One or more alphanumeric characters, hyphens or periods (the domain name and possibly subdomains)
	//   - Optional :port number (one or more digits)
	// For IP addresses:
	//   - Four groups of one to three digits, separated by periods (IPv4)
	//   - Optional :port number
	return `^(([a-zA-Z0-9-]+\.)*[a-zA-Z0-9][a-zA-Z0-9-]+\.[a-zA-Z]{2,6}?|(\d{1,3}\.){3}\d{1,3})(:[0-9]{1,5})?$`
}

func mongoRegex() string {

    // userPassword matches an optional username:password@ part.
    userPassword := `([a-zA-Z0-9]+:[a-zA-Z0-9]+@)?`
    // host matches the mandatory hostname:port part.
    host := `([a-zA-Z0-9.-]+:[0-9]+)`
    // additionalHosts matches any additional hostname:port pairs, separated by commas.
    additionalHosts := `(,[a-zA-Z0-9.-]+:[0-9]+)*`
    // database matches an optional /database part.
    database := `(/([a-zA-Z0-9]+))?`
    // options matches an optional ?options part.
    options := `(\?[a-zA-Z0-9=&]+)?`

    return "^mongodb:\\/\\/" + userPassword + host + additionalHosts + database + options + "$"
}

