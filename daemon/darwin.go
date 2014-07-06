package daemon

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"text/template"
)

type DarwinRecord struct {
	name        string
	description string
}

type Status struct {
	Label string
	PID   int
}

func newDaemon(name, description string) (*DarwinRecord, error) {

	return &DarwinRecord{name, description}, nil
}

// Standard service path for system daemons
func (darwin *DarwinRecord) servicePath() string {
	return "/Library/LaunchDaemons/" + darwin.name + ".plist"
}

// Install the service
func (darwin *DarwinRecord) Install() error {

	srvPath := darwin.servicePath()

	if _, err := os.Stat(srvPath); err == nil {
		return errors.New(darwin.description + " already installed")
	}

	file, err := os.Create(srvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	execPatch, err := executablePath()
	if err != nil {
		return err
	}

	templ, err := template.New("propertyList").Parse(propertyList)
	if err != nil {
		return err
	}

	if err := templ.Execute(
		file,
		&struct {
			Name, Path string
		}{darwin.name, execPatch},
	); err != nil {
		return err
	}

	return nil
}

// Remove the service
func (darwin *DarwinRecord) Remove() error {
	if err := os.Remove(darwin.servicePath()); err != nil {
		return err
	}

	return nil
}

func (darwin *DarwinRecord) Start() error {

	return exec.Command("launchctl", "load", darwin.servicePath()).Run()
}

func (darwin *DarwinRecord) Stop() error {

	return exec.Command("launchctl", "unload", darwin.servicePath()).Run()
}

func (darwin *DarwinRecord) Status() (string, error) {

	output, err := exec.Command("launchctl", "list", darwin.name).Output()
	if err != nil {
		return "service probably is stoped", nil
	}

	if matched, err := regexp.MatchString("whoisd", string(output)); err == nil && matched {
		reg := regexp.MustCompile("PID\" = ([0-9]+);")
		data := reg.FindStringSubmatch(string(output))
		if len(data) > 1 {
			return "service (pid  " + data[1] + ") is running...", nil
		}
	}

	return "service is stoped", nil
}

var propertyList = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>KeepAlive</key>
	<true/>
	<key>Label</key>
	<string>{{.Name}}</string>
	<key>ProgramArguments</key>
	<array>
	    <string>{{.Path}}</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
    <key>WorkingDirectory</key>
    <string>/usr/local/var</string>
    <key>StandardErrorPath</key>
    <string>/usr/local/var/log/{{.Name}}.log</string>
    <key>StandardOutPath</key>
    <string>/usr/local/var/log/{{.Name}}.log</string>
</dict>
</plist>
`
