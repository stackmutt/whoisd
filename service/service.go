package service

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/takama/whoisd/client"
	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/daemon"
	"github.com/takama/whoisd/storage"
)

const (
	Version = "0.07"
	Date    = "2014-07-06T23:25:01Z"
)

type ServiceRecord struct {
	Name   string
	Config *config.ConfigRecord
	daemon.Daemon
}

// Create a new service record
func New(name, description string) (*ServiceRecord, error) {
	daemonInstance, err := daemon.New(name, description)
	if err != nil {
		return nil, err
	}

	return &ServiceRecord{name, config.New(), daemonInstance}, nil
}

// Run or manage the service
func (srv *ServiceRecord) Run() (string, error) {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return "Install whois service ...", srv.Install()
		case "remove":
			return "Remove whois service ...", srv.Remove()
		case "start":
			return "Starting whois service ...", srv.Start()
		case "stop":
			return "Stoping whois service ...", srv.Stop()
		case "status":
			return srv.Status()
		}
	}
	mapp, err := srv.Config.Load()
	if err != nil {
		return "Loading mapping file was unsuccessful", err
	}
	serviceHostPort := fmt.Sprintf("%s:%d", srv.Config.Host, srv.Config.Port)
	log.Printf("%s started on %s\n", srv.Name, serviceHostPort)
	log.Printf("Used storage %s on %s:%d\n",
		srv.Config.Storage.StorageType,
		srv.Config.Storage.Host,
		srv.Config.Storage.Port,
	)
	listener, err := net.Listen("tcp", serviceHostPort)
	if err != nil {
		return "Probably use 'sudo' command to grant permission", err
	}
	channel := make(chan client.ClientRecord, srv.Config.Connections)
	repository := storage.New(srv.Config, mapp)
	for i := 0; i < srv.Config.Workers; i++ {
		go client.ProcessClient(channel, repository)
	}
	if srv.Config.TestMode == true {
		// make pipe connections for testing
		// connIn will ready to write into by function ProcessClient
		connIn, connOut := net.Pipe()
		defer connIn.Close()
		defer connOut.Close()
		newClient := client.ClientRecord{Conn: connIn}

		// prepare query for ProcessClient
		newClient.Query = []byte(srv.Config.TestQuery)

		// send it into channel
		channel <- newClient
		// just read answer from channel pipe
		buffer := make([]byte, 4096)
		numBytes, err := connOut.Read(buffer)
		log.Println("Read bytes:", numBytes)
		return string(buffer), err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		newClient := client.ClientRecord{Conn: conn}
		go newClient.HandleClient(channel)
	}

	// never happen, but need to complete code
	return "If you see that, you lucky bastard", nil
}
