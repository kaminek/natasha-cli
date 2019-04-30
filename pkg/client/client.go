package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"gopkg.in/urfave/cli.v2"
)

// Connect to the server
func Connect(c *cli.Context) error {

	server := c.String("address") + ":" + c.String("port")
	fmt.Printf("Connecting to %s...\n", server)

	conn, err := net.Dial("tcp4", server)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}

	go readConnection(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}
	return nil
}

func readConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()

			fmt.Printf("\b\b** %s\n> ", text)

			if !ok {
				fmt.Println("Reached EOF on server connection.")
				break
			}
		}
	}
}
