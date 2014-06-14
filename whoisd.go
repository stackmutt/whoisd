package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/takama/whoisd/config"
	"github.com/takama/whoisd/server"
)

// Init "Usage" helper
func init() {
	flag.Usage = func() {
		fmt.Println(config.Usage())
	}
}

func main() {
	conf := config.New()
	flag.Parse()
	if conf.ShowVersion {
		buildTime, err := time.Parse(time.RFC3339, server.Date)
		if err != nil {
			buildTime = time.Now()
		}
		fmt.Println("Whois Daemon", server.Version, buildTime.Format(time.RFC3339))
		os.Exit(0)
	}
	mapp, err := conf.Load()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}
	var daemon = server.New(conf, mapp)
	fmt.Printf("Whois Daemon started on %s:%d\n", conf.Host, conf.Port)
	fmt.Printf("Used storage %s on %s:%d\n", conf.Storage.StorageType, conf.Storage.Host, conf.Storage.Port)
	daemon.Run()
}
