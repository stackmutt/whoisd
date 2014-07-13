package daemon

import (
	"errors"
)

type WindowsRecord struct {
	name        string
	description string
}

func newDaemon(name, description string) (*WindowsRecord, error) {

	return &WindowsRecord{name, description}, nil
}

// Install the service
func (windows *WindowsRecord) Install() (string, error) {
	installAction := "Install " + windows.description + ":"

	return installAction + failed, errors.New("Windows daemon not supported")
}

// Remove the service
func (windows *WindowsRecord) Remove() (string, error) {
	removeAction := "Removing " + windows.description + ":"

	return removeAction + failed, errors.New("Windows daemon not supported")
}

// Start the service
func (windows *WindowsRecord) Start() (string, error) {
	startAction := "Starting " + windows.description + ":"

	return startAction + failed, errors.New("Windows daemon not supported")
}

// Stop the service
func (windows *WindowsRecord) Stop() (string, error) {
	stopAction := "Stopping " + windows.description + ":"

	return stopAction + failed, errors.New("Windows daemon not supported")
}

// Get service status
func (windows *WindowsRecord) Status() (string, error) {

	return "Status could not defined", errors.New("Windows daemon not supported")
}
