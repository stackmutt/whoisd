package storage

import (
	"flag"
	"testing"

	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/mapper"
)

func TestStorage(t *testing.T) {

	conf := config.New()
	flag.Parse()
	bundle := make(mapper.Bundle, 1)
	storage := New(conf, bundle)
	answer, ok := storage.Search("")
	if ok != false {
		t.Error("Expected ok is false, got", ok)
	}
	if answer != "not found\n" {
		t.Error("Expected answer is not found, got", answer)
	}
	answer, ok = storage.Search("aaa")
	if ok != false {
		t.Error("Expected ok is false, got", ok)
	}
	if answer != "not found\n" {
		t.Error("Expected answer is not found, got", answer)
	}
	entry := new(mapper.Entry)
	entry.TLDs = []string{"com"}
	entry.Fields = make(map[string]mapper.Field)
	entry.Fields["01"] = mapper.Field{
		Key:     "Domain Name: ",
		Name:    []string{"name"},
		Format:  "{idn}",
		Related: "name",
	}
	entry.Fields["03"] = mapper.Field{
		Key:   "Registrar WHOIS Server: ",
		Value: []string{"whois.markmonitor.com"},
	}
	entry.Fields["05"] = mapper.Field{
		Key:     "Updated Date: ",
		Name:    []string{"updatedDate"},
		Format:  "{date}",
		Related: "name",
	}
	entry.Fields["12"] = mapper.Field{
		Key:      "Domain Status: ",
		Name:     []string{"domainStatus"},
		Multiple: true,
		Related:  "name",
	}
	entry.Fields["13"] = mapper.Field{
		Key:       "Registry Registrant ID: ",
		Name:      []string{"handle"},
		Hide:      true,
		Related:   "ownerHandle",
		RelatedBy: "handle",
		RelatedTo: "customer",
	}
	entry.Fields["21"] = mapper.Field{
		Key: "Registrant Phone: ",
		Name: []string{
			"phone.countryCode",
			"phone.areaCode",
			"phone.subscriberNumber",
		},
		Format:    "{string}.{string}{string}",
		Related:   "ownerHandle",
		RelatedBy: "handle",
		RelatedTo: "customer",
	}
	entry.Fields["52"] = mapper.Field{
		Key:       "Name Server: ",
		Name:      []string{"name"},
		Multiple:  true,
		Related:   "nsgroupId",
		RelatedBy: "nsgroupId",
		RelatedTo: "nameserver",
	}
	entry.Fields["55"] = mapper.Field{
		Key:       "",
		Value:     []string{"1"},
		Name:      []string{"updatedDate"},
		Format:    ">>> Last update of WHOIS database: {date} <<<",
		Related:   "whois",
		RelatedBy: "id",
		RelatedTo: "whois",
	}
	bundle = append(bundle, *entry)
	storage = New(conf, bundle)
	answer, ok = storage.Search("google.com")
	if ok != true {
		t.Error("Expected ok is true, got", ok)
	}
	expected := `Domain Name: google.com
Registrar WHOIS Server: whois.markmonitor.com
Updated Date: 2014-05-19T04:00:17Z
Domain Status: clientUpdateProhibited
Domain Status: clientTransferProhibited
Domain Status: clientDeleteProhibited
Registry Registrant ID: 
Registrant Phone: +1.6502530000
Name Server: ns1.google.com
Name Server: ns2.google.com
Name Server: ns3.google.com
Name Server: ns4.google.com
>>> Last update of WHOIS database: 2014-06-01T11:00:07Z <<<
`
	if answer != expected {
		t.Error("Expected answer:\n", expected, "\n, got:\n", answer)
	}
	conf.Storage.StorageType = "mysql"
	storage = New(conf, bundle)
	answer, ok = storage.Search("mmm")
	if ok != false {
		t.Error("Expected ok is false, got", ok)
	}
	if answer != "not found\n" {
		t.Error("Expected answer is not found, got", answer)
	}
	conf.Storage.StorageType = "elasticsearch"
	storage = New(conf, bundle)
	answer, ok = storage.Search("eee")
	if ok != false {
		t.Error("Expected ok is false, got", ok)
	}
	if answer != "not found\n" {
		t.Error("Expected answer is not found, got", answer)
	}
}
