package daemon

import (
	"log"
)

type DarwinRecord struct {
	name string
}

func newDaemon(name string) (*DarwinRecord, error) {

	return &DarwinRecord{name}, nil
}

func (darwin *DarwinRecord) Install() error {
	log.Println("Mac OS X service has not been installed due to dummy mode")

	return nil
}

func (darwin *DarwinRecord) Remove() error {
	log.Println("Mac OS X service has not been removed due to dummy mode")

	return nil
}
