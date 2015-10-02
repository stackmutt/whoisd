package config

import (
	"testing"
)

func TestConfig(t *testing.T) {

	conf := New()
	conf.ConfigPath = ""
	conf.MappingPath = ""
	bundle, err := conf.Load()
	if err == nil {
		t.Error("Expected error of loading mapping file, got nothing")
	}
	if conf.Connections != 1000 {
		t.Error("Expected 100 active connections, got", conf.Connections)
	}
	if conf.Workers != 1000 {
		t.Error("Expected 100 workers, got", conf.Workers)
	}
	if len(bundle) != 0 {
		t.Error("Expected empty bundle slice, got not empty bundle slice")
	}
	conf.ConfigPath = "../test/testconfig.conf"
	conf.MappingPath = "../test/testmapping.json"
	bundle, err = conf.Load()
	if err != nil {
		t.Error("Expected config loading without error, got", err.Error())
	}
	if conf.Connections != 100 {
		t.Error("Expected 100 active connections, got", conf.Connections)
	}
	if conf.Workers != 100 {
		t.Error("Expected 100 workers, got", conf.Workers)
	}
	if len(bundle) == 0 {
		t.Error("Expected loading of bundle, got empty bundle")
	}
	entry := bundle.EntryByTLD("com")
	key := "01"
	expected := "Domain Name: "
	if entry.Fields[key].Key != expected {
		t.Error("Expected", expected, ", got", entry.Fields[key].Key)
	}
	key = "02"
	expected = "name"
	if entry.Fields[key].Related != expected {
		t.Error("Expected", expected, ", got", entry.Fields[key].Related)
	}
	key = "05"
	expected = "{date}"
	if entry.Fields[key].Format != expected {
		t.Error("Expected", expected, ", got", entry.Fields[key].Format)
	}
}
