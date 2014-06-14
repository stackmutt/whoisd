package client

import (
	"bytes"
	"log"
	"net"

	"github.com/takama/whoisd/storage"
)

const (
	queryBufferSize = 256
)

type ClientRecord struct {
	Conn  net.Conn
	Query []byte
}

// Sends the client data into the channel
func (client *ClientRecord) HandleClient(channel chan ClientRecord) {
	buffer := make([]byte, queryBufferSize)
	numBytes, err := client.Conn.Read(buffer)
	if numBytes == 0 || err != nil {
		return
	}
	client.Query = bytes.ToLower(bytes.Trim(buffer, "\u0000\u000a\u000d"))
	channel <- *client
}

// Asynchronous the client handling
func ProcessClient(channel chan ClientRecord, repository *storage.StorageRecord) {
	for {
		message := <-channel
		query := string(message.Query)
		data, ok := repository.Search(query)
		message.Conn.Write([]byte(data))
		log.Println(message.Conn.RemoteAddr().String(), query, ok)
		message.Conn.Close()
	}
}
