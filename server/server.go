package server

import (
	"fmt"
	"log"
	"net"

	"github.com/takama/whoisd/client"
	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/mapper"
	"github.com/takama/whoisd/storage"
)

const (
	Version = "0.05"
	Date    = "2014-06-14T21:26:17Z"
)

type ServerRecord struct {
	Config *config.ConfigRecord
	Mapper *mapper.MapperRecord
}

// Returns new Server instance
func New(conf *config.ConfigRecord, mapp *mapper.MapperRecord) *ServerRecord {
	return &ServerRecord{
		conf,
		mapp,
	}
}

// Run the Server instance
func (server *ServerRecord) Run() {
	address := fmt.Sprintf("%s:%d", server.Config.Host, server.Config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	channel := make(chan client.ClientRecord, server.Config.Connections)
	repository := storage.New(server.Config, server.Mapper)
	for i := 0; i < server.Config.Workers; i++ {
		go client.ProcessClient(channel, repository)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		newClient := client.ClientRecord{Conn: conn}
		go newClient.HandleClient(channel)
	}
}
