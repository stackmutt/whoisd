package service

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/takama/whoisd/client"
	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/storage"
)

const (
	Version = "0.06"
	Date    = "2014-06-29T11:02:14Z"
)

type ServiceRecord struct {
	Name   string
	Config *config.ConfigRecord
}

func New(name string) *ServiceRecord {

	return &ServiceRecord{name, config.New()}
}

func (srv *ServiceRecord) Check() (doRun bool, err error) {
	doRun = true
	err = nil

	return doRun, err
}

func (srv *ServiceRecord) Run() error {
	mapp, err := srv.Config.Load()
	if err != nil {
		return err
	}
	serviceHostPort := fmt.Sprintf("%s:%d", srv.Config.Host, srv.Config.Port)
	fmt.Printf("%s started on %s\n", srv.Name, serviceHostPort)
	fmt.Printf("Used storage %s on %s:%d\n",
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

	return nil
}

func executablePath() (string, error) {
	return filepath.Abs(os.Args[0])
}
