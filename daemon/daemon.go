package daemon

import (
	"os/user"
)

const (
	rootPrivileges = "You must have root user privileges. Possibly using 'sudo' command should help"
	success        = "\t\t\t\t\t[  \033[32mOK\033[0m  ]"
	failed         = "\t\t\t\t\t[\033[31mFAILED\033[0m]"
)

type Daemon interface {
	Install() (string, error)
	Remove() (string, error)
	Start() (string, error)
	Stop() (string, error)
	Status() (string, error)
}

func New(name, description string) (Daemon, error) {
	return newDaemon(name, description)
}

func executablePath() (string, error) {
	return execPath()
}

func checkPrivileges() bool {

	if user, err := user.Current(); err == nil && user.Gid == "0" {
		return true
	}
	return false
}
