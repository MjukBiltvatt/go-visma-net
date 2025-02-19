package vismanet

import "fmt"

// debugDumpResponse prints the request to stdout if debugging is enabled
func debugDumpResponse(c *Client, resp interface{}) {
	if c.Debug {
		dumpInterface("Response", resp)
	}
}

// dumpInterface prints the description and the object to stdout
func dumpInterface(description string, obj interface{}) {
	fmt.Printf("%s: %+v\n\n", description, obj)
}
