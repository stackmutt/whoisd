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
	Version = "0.06"
	Date    = "2014-06-29T11:02:14Z"
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
func (srv *ServiceRecord) Run() error {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return srv.Install()
		case "remove":
			return srv.Remove()
		}
	}
	mapp, err := srv.Config.Load()
	if err != nil {
		return err
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
		return err
	}
	channel := make(chan client.ClientRecord, srv.Config.Connections)
	repository := storage.New(srv.Config, mapp)
	for i := 0; i < srv.Config.Workers; i++ {
		go client.ProcessClient(channel, repository)
	}
	if srv.Config.TestMode == true {
		return nil
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
	return nil
}
