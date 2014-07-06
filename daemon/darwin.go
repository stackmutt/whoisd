package daemon

import (
	"errors"
	"log"
	"os"
	"text/template"
)

type DarwinRecord struct {
	name        string
	description string
}

func newDaemon(name, description string) (*DarwinRecord, error) {

	return &DarwinRecord{name, description}, nil
}

func (darwin *DarwinRecord) servicePath() string {
	return "/Library/LaunchDaemons/" + darwin.name + ".plist"
}

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
	err = templ.Execute(
		file,
		&struct {
			Name, Path string
		}{darwin.name, execPatch},
	)
	if err != nil {
		return err
	}

	log.Println(darwin.description, "has been installed")

	return nil
}

func (darwin *DarwinRecord) Remove() error {
	if err := os.Remove(darwin.servicePath()); err != nil {
		return err
	}
	log.Println(darwin.description, "has been removed")

	return nil
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
