package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/service"
)

// Init "Usage" helper
func init() {
	flag.Usage = func() {
		fmt.Println(config.Usage())
	}
}

func main() {
	serviceName := "Whois Daemon"
	serviceInstance, err := service.New(serviceName)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	flag.Parse()
	if serviceInstance.Config.ShowVersion {
		buildTime, err := time.Parse(time.RFC3339, service.Date)
		if err != nil {
			buildTime = time.Now()
		}
		fmt.Println(serviceName, service.Version, buildTime.Format(time.RFC3339))
		os.Exit(0)
	}
	doRun, err := serviceInstance.Manage()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	if doRun {
		if err := serviceInstance.Run(); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
