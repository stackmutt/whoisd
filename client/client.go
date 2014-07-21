package client

import (
	"bytes"
	"log"
	"net"

	"code.google.com/p/go.net/idna"
	"github.com/takama/whoisd/storage"
)

const (
	queryBufferSize = 256
)

type ClientRecord struct {
	Conn  net.Conn
	Query []byte
}

// Sends a client data into the channel
func (client *ClientRecord) HandleClient(channel chan<- ClientRecord) {
	defer func() {
		if recovery := recover(); recovery != nil {
			log.Println("Recovered in HandleClient:", recovery)
			channel <- *client
		}
	}()
	buffer := make([]byte, queryBufferSize)
	numBytes, err := client.Conn.Read(buffer)
	if numBytes == 0 || err != nil {
		return
	}
	client.Query = bytes.ToLower(bytes.Trim(buffer, "\u0000\u000a\u000d"))
	channel <- *client
}

// Asynchronous a client handling
func ProcessClient(channel <-chan ClientRecord, repository *storage.StorageRecord) {
	message := ClientRecord{}
	defer func() {
		if recovery := recover(); recovery != nil {
			log.Println("Recovered in ProcessClient:", recovery)
			if message.Conn != nil {
				message.Conn.Close()
			}
		}
	}()
	for {
		message = <-channel
		query, err := idna.ToASCII(string(message.Query))
		if err != nil {
			query = string(message.Query)
		}
		data, ok := repository.Search(query)
		message.Conn.Write([]byte(data))
		log.Println(message.Conn.RemoteAddr().String(), query, ok)
		message.Conn.Close()
	}
}
