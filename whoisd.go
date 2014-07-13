package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/service"
)

// Init "Usage" helper
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Usage = func() {
		fmt.Println(config.Usage())
	}
}

func main() {
	serviceName, serviceDescription := "whoisd", "Whois Daemon"
	serviceInstance, err := service.New(serviceName, serviceDescription)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
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
	status, err := serviceInstance.Run()
	if err != nil {
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	// Wait for logger output
	time.Sleep(100 * time.Millisecond)
	fmt.Println(status)
}
