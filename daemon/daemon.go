package daemon

type DaemonRecord struct {
	Name string
}

func New(name string) (Daemon, error) {
	return newDaemon(name)
}

type Daemon interface {
	Install() error
	Remove() error
}
