package client

import (
	"net"
	"strings"
	"testing"

	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/storage"
)

func TestClientHandling(t *testing.T) {
	conf := config.New()
	conf.ConfigPath = "../test/testconfig.conf"
	conf.MapperPath = "../test/testmapper.json"
	mapp, err := conf.Load()
	if err != nil {
		t.Error("Expected config loading without error, got", err.Error())
	}
	channel := make(chan ClientRecord, conf.Connections)
	repository := storage.New(conf, mapp)
	go ProcessClient(channel, repository)

	connIn, connOut := net.Pipe()
	newClient := ClientRecord{Conn: connIn}
	newClient.Query = []byte("google.com")
	channel <- newClient
	buffer := make([]byte, 256)
	numBytes, err := connOut.Read(buffer)
	if err != nil {
		t.Error("Network communication error", err.Error())
	}
	if numBytes == 0 || len(buffer) == 0 {
		t.Error("Expexted some data resd, got", string(buffer))
	}
	partAnswer := "Updated Date: 2014-05-19T04:00:17Z"
	if !strings.Contains(string(buffer), partAnswer) {
		t.Error("Expexted that contains", partAnswer, ", got", string(buffer))
	}
}
