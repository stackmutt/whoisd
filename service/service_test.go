package service

import (
	"strings"
	"testing"
)

func TestService(t *testing.T) {
	serviceName, serviceDescription := "whoisd", "Whois Daemon"
	srv, err := New(serviceName, serviceDescription)
	if err != nil {
		t.Error("Expected service create without error, got", err.Error())
	}
	if srv.Name != serviceName {
		t.Error("Expected service name must be ", serviceName, ", got", srv.Name)
	}
	if srv.Config.Host != "localhost" {
		t.Error("Expected server host is localhost, got", srv.Config.Host)
	}
	srv.Config.ConfigPath = "../test/testconfig.conf"
	srv.Config.MappingPath = "../test/testmapping.json"
	answer, err := srv.Run()
	if err != nil {
		t.Error("Expected service run without error, got", err.Error())
	}
	partAnswer := "Updated Date: 2014-05-19T04:00:17Z"
	if !strings.Contains(string(answer), partAnswer) {
		t.Error("Expexted that contains", partAnswer, ", got", string(answer))
	}
}
