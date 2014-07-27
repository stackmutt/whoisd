package daemon

import (
	"os"
	"os/exec"
	"os/user"
)

const (
	rootPrivileges = "You must have root user privileges. Possibly using 'sudo' command should help"
	success        = "\t\t\t\t\t[  \033[32mOK\033[0m  ]"
	failed         = "\t\t\t\t\t[\033[31mFAILED\033[0m]"
)

// Daemon interface has standard set of methods/commands
// install, remove, start, stop, status
type Daemon interface {
	Install() (string, error)
	Remove() (string, error)
	Start() (string, error)
	Stop() (string, error)
	Status() (string, error)
}

// New - Create new daemon
func New(name, description string) (Daemon, error) {
	return newDaemon(name, description)
}

func executablePath() (string, error) {
	if path, err := exec.LookPath("whoisd"); err == nil {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			return execPath()
		}
		return path, nil
	}
	return execPath()
}

func checkPrivileges() bool {

	if user, err := user.Current(); err == nil && user.Gid == "0" {
		return true
	}
	return false
}
