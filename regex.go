package env

func httpRegex() string {

    // scheme matches either http:// or https:// at the beginning of the string.
    scheme := `^(http|https):\/\/`
    // domainName matches one or more characters that are not whitespace, forward slash, or colon.
    domainName := `[^\s/:]+`
    // port matches an optional :port part, where the port is one or more digits.
    port := `(:[0-9]+)?`

    return scheme + domainName + port + "$"
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
    // domainName matches a domain name, which can include one or more subdomains, each ending with a dot, followed by a domain name with at least one character before the dot and a top-level domain of 2 to 6 characters.
    domainName := `(([a-zA-Z0-9-]+\.)*[a-zA-Z0-9][a-zA-Z0-9-]+\.[a-zA-Z]{2,6}?)`
    // ipAddress matches an IPv4 address, which is four groups of one to three digits separated by periods.
    ipAddress := `((\d{1,3}\.){3}\d{1,3})`
    // port matches an optional :port part, where the port is 1 to 5 digits.
    port := `(:[0-9]{1,5})?$`

    return "^" + domainName + "|" + ipAddress + port
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

