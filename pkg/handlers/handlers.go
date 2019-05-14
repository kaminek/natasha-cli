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

// SendCmd common func handling basic command
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
	_, err = conn.Write(sendBuf.Bytes())
	if err != nil {
		log.Fatal("Connection write error: ", err)
		return err
	}

	recvBuf := make([]byte, unsafe.Sizeof(*reply))
	_, err = conn.Read(recvBuf)
	if err != nil {
		log.Fatal("Connection read error: ", err)
		return err
	}

	// Write byte stream to struct
	r := bytes.NewReader(recvBuf)
	err = binary.Read(r, binary.BigEndian, reply)
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

// NatashaReload Triggers rules reload
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

// NatashaExit Force server stop
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

// NatashaResetStats reset all stats
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

// NatashaVersion Returns server version
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

// NatashaDPDKStats Show DPDK stats
func NatashaDPDKStats(c *cli.Context) error {

	reply := headers.NatashaCmdReply{}
	DpdkPortStats := headers.NatashaEthStats{}

	conn, err := client.Connect(c)
	if err != nil {
		log.Fatal("Socket connection error: ", err)
		return err
	}
	defer conn.Close()

	err = SendCmd(conn, headers.NatashaCmdDpdkStats, &reply)
	if err != nil {
		log.Fatal("version: ", err)
		return err
	}

	ports := int(reply.DataSize) / int(unsafe.Sizeof(DpdkPortStats))
	recvBuf := make([]byte, unsafe.Sizeof(DpdkPortStats))

	for p := 0; p < ports; p++ {
		_, err = conn.Read(recvBuf)
		if err != nil {
			log.Fatal("Failed to read data", err)
			return err
		}

		r := bytes.NewReader(recvBuf)
		err = binary.Read(r, binary.BigEndian, &DpdkPortStats)
		if err != nil {
			log.Fatal("Write to data structure error: ", err)
			return err
		}

		fmt.Println("Port ", p)
		fmt.Printf("%+v\n", DpdkPortStats)
	}

	return nil
}

// NatashaAppStats handle app-stats command
func NatashaAppStats(c *cli.Context) error {

	reply := headers.NatashaCmdReply{}
	AppCoreStats := headers.NatashaAppStats{}
	var coreID uint8

	conn, err := client.Connect(c)
	if err != nil {
		log.Fatal("Socket connection error: ", err)
		return err
	}
	defer conn.Close()

	err = SendCmd(conn, headers.NatashaCmdAppStats, &reply)
	if err != nil {
		log.Fatal("version: ", err)
		return err
	}

	cores := int(reply.DataSize) /
		int(unsafe.Sizeof(AppCoreStats)+unsafe.Sizeof(coreID))

	for c := 0; c < cores; c++ {
		recvBuf := make([]byte, unsafe.Sizeof(coreID))
		_, err := conn.Read(recvBuf)
		if err != nil {
			log.Fatal("Failed to read data", err)
			return err
		}
		// it's a uint8 same as one byte
		coreID = recvBuf[0]

		recvBuf = make([]byte, unsafe.Sizeof(AppCoreStats))
		_, err = conn.Read(recvBuf)
		if err != nil {
			log.Fatal("Failed to read data", err)
			return err
		}

		r := bytes.NewReader(recvBuf)
		err = binary.Read(r, binary.BigEndian, &AppCoreStats)
		if err != nil {
			log.Fatal("Write to data structure error: ", err)
			return err
		}

		fmt.Println("Core ", coreID)
		fmt.Printf("%+v\n", AppCoreStats)
	}

	return nil
}
