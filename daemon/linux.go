package daemon

import (
	"log"
)

type LinuxRecord struct {
	name        string
	description string
}

func newDaemon(name, description string) (*LinuxRecord, error) {

	return &LinuxRecord{name, description}, nil
}

func (linux *LinuxRecord) Install() error {
	log.Println(linux.description, "has not been installed due to dummy mode")

	return nil
}

func (linux *LinuxRecord) Remove() error {
	log.Println(linux.description, "has not been removed due to dummy mode")

	return nil
}

func (linux *LinuxRecord) Start() error {
	log.Println(linux.description, "has not been started due to dummy mode")

	return nil
}

func (linux *LinuxRecord) Stop() error {
	log.Println(linux.description, "has not been stoped due to dummy mode")

	return nil
}

func (linux *LinuxRecord) Status() (string, error) {
	log.Println(linux.description, "has not benn used due to dummy mode")

	return "", nil
}
