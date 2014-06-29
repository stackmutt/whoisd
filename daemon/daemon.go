package daemon

type DaemonRecord struct {
	Name string
}

type Daemon interface {
	Install() error
	Remove() error
}

func New(name string) (Daemon, error) {
	return newDaemon(name)
}
