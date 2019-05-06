package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"unsafe"

	"github.com/kaminek/natasha-cli/pkg/client"
	"github.com/kaminek/natasha-cli/pkg/headers"
	"gopkg.in/urfave/cli.v2"
)

// NatashaQuery struct
type NatashaQuery struct {
	Type uint8
}

// SendCmd ere
func SendCmd(conn net.Conn, queryType uint8, reply *headers.NatashaCmdReply) error {

	query := NatashaQuery{
		Type: queryType,
	}
	sendBuf := &bytes.Buffer{}

	// Write struct to byte stream
	err := binary.Write(sendBuf, binary.BigEndian, query)
	if err != nil {
		log.Fatal("Write to byte stream error: ", err)
		return err
	}

	// Write to socket
	nb, err := conn.Write(sendBuf.Bytes())
	if err != nil {
		log.Fatal("Connection write error: ", err)
		return err
	}

	fmt.Printf("%d bytes sent\n", nb)

	recvBuf := make([]byte, unsafe.Sizeof(*reply))
	nb, err = conn.Read(recvBuf)
	if err != nil {
		log.Fatal("Connection read error: ", err)
		return err
	}

	fmt.Printf("%d bytes received\n", nb)

	// Write byte stream to struct
	buf := bytes.NewReader(recvBuf)
	err = binary.Read(buf, binary.BigEndian, reply)
	if err != nil {
		log.Fatal("Write to data structure error: ", err)
		return err
	}

	if reply.Status != 0 {
		return fmt.Errorf("Command Failed (%g)", reply.Status)
	}

	return nil
}

// NatashaStatus Status checker
func NatashaStatus(c *cli.Context) error {

	reply := headers.NatashaCmdReply{}

	conn, err := client.Connect(c)
	if err != nil {
		log.Fatal("Socket connection error: ", err)
		return err
	}
	defer conn.Close()

	err = SendCmd(conn, headers.NatashaCmdStatus, &reply)
	if err != nil {
		log.Fatal("status: ", err)
		fmt.Println("Natasha status KO")
		return err
	}

	fmt.Println("Natasha status OK")

	return nil
}

// NatashaReload Status checker
func NatashaReload(c *cli.Context) error {

	reply := headers.NatashaCmdReply{}

	conn, err := client.Connect(c)
	if err != nil {
		log.Fatal("Socket connection error: ", err)
		return err
	}
	defer conn.Close()

	err = SendCmd(conn, headers.NatashaCmdReload, &reply)
	if err != nil {
		log.Fatal("reload: ", err)
		return err
	}

	fmt.Println("Natasha rules reload succeeded")

	return nil
}

// NatashaExit Status checker
func NatashaExit(c *cli.Context) error {

	reply := headers.NatashaCmdReply{}

	conn, err := client.Connect(c)
	if err != nil {
		log.Fatal("Socket connection error: ", err)
		return err
	}
	defer conn.Close()

	err = SendCmd(conn, headers.NatashaCmdExit, &reply)
	if err != nil {
		log.Fatal("exit: ", err)
		return err
	}

	fmt.Println("Natasha exited successfuly")

	return nil
}

// NatashaResetStats Status checker
func NatashaResetStats(c *cli.Context) error {

	reply := headers.NatashaCmdReply{}

	conn, err := client.Connect(c)
	if err != nil {
		log.Fatal("Socket connection error: ", err)
		return err
	}
	defer conn.Close()

	err = SendCmd(conn, headers.NatashaCmdResetStats, &reply)
	if err != nil {
		log.Fatal("reset-stats: ", err)
		return err
	}

	fmt.Println("Natasha stats reset succeeded")

	return nil
}

// NatashaVersion Status checker
func NatashaVersion(c *cli.Context) error {

	reply := headers.NatashaCmdReply{}

	conn, err := client.Connect(c)
	if err != nil {
		log.Fatal("Socket connection error: ", err)
		return err
	}
	defer conn.Close()

	err = SendCmd(conn, headers.NatashaCmdVersion, &reply)
	if err != nil {
		log.Fatal("version: ", err)
		return err
	}

	recvBuf := make([]byte, reply.DataSize)
	_, err = conn.Read(recvBuf)
	if err != nil {
		log.Fatal("Connection error", err)
		return err
	}

	fmt.Println("Natasha version ", string(recvBuf))

	return nil
}
