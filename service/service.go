package service

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

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
func New(name string) (*ServiceRecord, error) {
	daemonInstance, err := daemon.New(name)
	if err != nil {
		return nil, err
	}

	return &ServiceRecord{name, config.New(), daemonInstance}, nil
}

// Manage a service (Install or Remove)
func (srv *ServiceRecord) Manage() (doRun bool, err error) {
	doRun = true
	err = nil
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			doRun = false
			if err = srv.Install(); err != nil {
				return doRun, err
			}
		case "remove":
			doRun = false
			if err = srv.Remove(); err != nil {
				return doRun, err
			}
		default:
			doRun = false
			return doRun, errors.New("Unrecognized command: " + command)
		}
	}

	return doRun, err
}

// Run the service
func (srv *ServiceRecord) Run() error {
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

func executablePath() (string, error) {
	return filepath.Abs(os.Args[0])
}
