package daemon

import (
	"log"
)

type LinuxRecord struct {
	name string
}

func newDaemon(name string) (*LinuxRecord, error) {

	return &LinuxRecord{name}, nil
}

func (linux *LinuxRecord) Install() error {
	log.Println("Linux service has not been installed due to dummy mode")

	return nil
}

func (linux *LinuxRecord) Remove() error {
	log.Println("Linux service has not been removed due to dummy mode")

	return nil
}
