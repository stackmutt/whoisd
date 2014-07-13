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

func newDaemon(name, description string) (*DarwinRecord, error) {

	return &DarwinRecord{name, description}, nil
}

// Standard service path for system daemons
func (darwin *DarwinRecord) servicePath() string {
	return "/Library/LaunchDaemons/" + darwin.name + ".plist"
}

// Check service is running
func (darwin *DarwinRecord) checkStatus() (string, bool) {
	output, err := exec.Command("launchctl", "list", darwin.name).Output()
	if err == nil {
		if matched, err := regexp.MatchString(darwin.name, string(output)); err == nil && matched {
			reg := regexp.MustCompile("PID\" = ([0-9]+);")
			data := reg.FindStringSubmatch(string(output))
			if len(data) > 1 {
				return "Service (pid  " + data[1] + ") is running...", true
			} else {
				return "Service is running...", true
			}
		}
	}

	return "Service is stoped", false
}

// Install the service
func (darwin *DarwinRecord) Install() (string, error) {
	installAction := "Install " + darwin.description + ":"

	if checkPrivileges() == false {
		return installAction + failed, errors.New(rootPrivileges)
	}

	srvPath := darwin.servicePath()

	if _, err := os.Stat(srvPath); err == nil {
		return installAction + failed, errors.New(darwin.description + " already installed")
	}

	file, err := os.Create(srvPath)
	if err != nil {
		return installAction + failed, err
	}
	defer file.Close()

	execPatch, err := executablePath()
	if err != nil {
		return installAction + failed, err
	}

	templ, err := template.New("propertyList").Parse(propertyList)
	if err != nil {
		return installAction + failed, err
	}

	if err := templ.Execute(
		file,
		&struct {
			Name, Path string
		}{darwin.name, execPatch},
	); err != nil {
		return installAction + failed, err
	}

	return installAction + success, nil
}

// Remove the service
func (darwin *DarwinRecord) Remove() (string, error) {
	removeAction := "Removing " + darwin.description + ":"

	if checkPrivileges() == false {
		return removeAction + failed, errors.New(rootPrivileges)
	}

	if err := os.Remove(darwin.servicePath()); err != nil {
		return removeAction + failed, err
	}

	return removeAction + success, nil
}

func (darwin *DarwinRecord) Start() (string, error) {
	startAction := "Starting " + darwin.description + ":"

	if checkPrivileges() == false {
		return startAction + failed, errors.New(rootPrivileges)
	}

	if _, status := darwin.checkStatus(); status == true {
		return startAction + failed, errors.New("service already running")
	}

	if err := exec.Command("launchctl", "load", darwin.servicePath()).Run(); err != nil {
		return startAction + failed, err
	}

	return startAction + success, nil
}

func (darwin *DarwinRecord) Stop() (string, error) {
	stopAction := "Stopping " + darwin.description + ":"

	if checkPrivileges() == false {
		return stopAction + failed, errors.New(rootPrivileges)
	}

	if _, status := darwin.checkStatus(); status == false {
		return stopAction + failed, errors.New("service already stopped")
	}

	if err := exec.Command("launchctl", "unload", darwin.servicePath()).Run(); err != nil {
		return stopAction + failed, err
	}

	return stopAction + success, nil
}

func (darwin *DarwinRecord) Status() (string, error) {

	if checkPrivileges() == false {
		return "", errors.New(rootPrivileges)
	}
	statusAction, _ := darwin.checkStatus()

	return statusAction, nil
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
