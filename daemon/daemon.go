package daemon

import (
	"os"
	"path/filepath"
)

type Daemon interface {
	Install() error
	Remove() error
}

func New(name, description string) (Daemon, error) {
	return newDaemon(name, description)
}

func executablePath() (string, error) {
	return filepath.Abs(os.Args[0])
}
