package client

import (
	"fmt"
	"net"

	"gopkg.in/urfave/cli.v2"
)

// Connect to the server
func Connect(c *cli.Context) (net.Conn, error) {

	server := c.String("address") + ":" + c.String("port")
	fmt.Printf("Connecting to %s...\n", server)
	conn, err := net.Dial("tcp4", server)
	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem while connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
	}

	return conn, err
}
