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
	// Define a simple regex pattern for a ws or wss URL.
	// Starts with ws:// or wss://
	// Followed by: 
	//   - One or more alphanumeric characters, hyphens or periods (the domain name and possibly subdomains)
	//   - Optional :port number (one or more digits)
	//   - The remainder of the URL, which can include alphanumeric characters, hyphens, periods, slashes, and a few other valid URL characters
	return `^(ws://|wss://)?([a-zA-Z0-9-]+\.)*[a-zA-Z0-9][a-zA-Z0-9-]+\.[a-zA-Z]{2,6}?(:[0-9]{1,5})?(\/.*)?$`
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
    // Starts with mongodb://
    // Optionally includes username:password@
    // Includes a hostname which can be a domain name (alphanumeric and hyphen characters plus optional periods) or an IP address
    // Optionally includes :port
    // Optionally includes additional hostname:port pairs, separated by commas
    // Optionally includes /database
    // Optionally includes ?options, with options as key=value pairs separated by &
    return `^mongodb:\/\/([a-zA-Z0-9]+:[a-zA-Z0-9]+@)?([a-zA-Z0-9.-]+(:[0-9]+)?)(,[a-zA-Z0-9.-]+(:[0-9]+)?)*(/([a-zA-Z0-9]+)(\?[a-zA-Z0-9=&]+)?)?$`
}
