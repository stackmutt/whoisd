package server

import (
	"testing"

	"github.com/takama/whoisd/config"
)

func TestFerver(t *testing.T) {
	conf := config.New()
	conf.ConfigPath = "../test/testconfig.conf"
	conf.MapperPath = "../test/testmapper.json"
	mapp, err := conf.Load()
	if err != nil {
		t.Error("Expected config loading without error, got", err.Error())
	}
	srv := New(conf, mapp)
	if srv.Config.Host != "localhost" {
		t.Error("Expected server host is localhost, got", srv.Config.Host)
	}
}
