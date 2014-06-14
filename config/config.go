package config

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/takama/whoisd/mapper"
)

// The default values: path of the configuration file, host, port
const (
	defaultConfigPath  = "/etc/whoisd/whoisd.conf"
	defaultMapperPath  = "mapper.json"
	defaultHost        = "localhost"
	defaultPort        = 43
	defaultWorkers     = 1000
	defaultConnections = 1000
	defaultStorageType = "Elasticsearch"
	defaultStorageHost = "localhost"
	defaultStoragePort = 9200
	defaultIndexBase   = "whois"
	defaultTypeTable   = "domain"
)

type ConfigRecord struct {
	configPath string
	mapperPath string

	ShowVersion bool

	Host        string
	Port        int
	Workers     int
	Connections int
	Storage     StorageConfig
}

type StorageConfig struct {
	StorageType string
	Host        string
	Port        int
	IndexBase   string
	TypeTable   string
}

// returns the configuration initialized with default values
func New() *ConfigRecord {
	config := new(ConfigRecord)
	flag.BoolVar(&config.ShowVersion, "version", false, "show version")
	flag.BoolVar(&config.ShowVersion, "v", false, "show version")
	flag.StringVar(&config.configPath, "config", defaultConfigPath, "path to configuration file")
	flag.StringVar(&config.mapperPath, "mapper", defaultMapperPath, "path to mapper file")
	flag.StringVar(&config.Host, "host", defaultHost, "host name or IP address")
	flag.IntVar(&config.Port, "port", defaultPort, "port number")
	flag.IntVar(&config.Workers, "work", defaultWorkers, "number of active workers")
	flag.IntVar(&config.Connections, "conn", defaultConnections, "number of active conections")
	flag.StringVar(&config.Storage.StorageType, "storage", defaultStorageType, "type of storage (Elasticsearch, Mysql)")
	flag.StringVar(&config.Storage.Host, "shost", defaultStorageHost, "storage host name or IP address")
	flag.IntVar(&config.Storage.Port, "sport", defaultStoragePort, "storage port number")
	flag.StringVar(&config.Storage.IndexBase, "base", defaultIndexBase, "storage index or database name")
	flag.StringVar(&config.Storage.TypeTable, "table", defaultTypeTable, "storage type or table name")

	return config
}

// Loads the configurtion from the config file or from th command line
func (config *ConfigRecord) Load() (*mapper.MapperRecord, error) {
	var path string

	mRecord := new(mapper.MapperRecord)
	mStat, err := os.Stat(config.mapperPath)
	if !os.IsNotExist(err) {
		mFile, err := os.Open(config.mapperPath)
		if err != nil {
			return nil, err
		}
		defer mFile.Close()
		mData := make([]byte, mStat.Size())
		if _, err := mFile.Read(mData); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(mData, &mRecord); err != nil {
			return nil, err
		}
	}
	stat, err := os.Stat(config.configPath)
	if os.IsNotExist(err) {
		return mRecord, nil
	}
	file, err := os.Open(config.configPath)
	if err != nil {
		return mRecord, err
	}
	defer file.Close()
	data := make([]byte, stat.Size())
	if _, err := file.Read(data); err != nil {
		return mRecord, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return mRecord, err
	}
	flags := flag.NewFlagSet("whoisd", flag.ContinueOnError)
	// Begin ignored flags
	flags.StringVar(&path, "config", "", "")
	flags.StringVar(&path, "mapper", "", "")
	// End ignored flags
	flags.StringVar(&config.Host, "host", config.Host, "")
	flags.IntVar(&config.Port, "port", config.Port, "")
	flags.IntVar(&config.Workers, "work", config.Workers, "")
	flags.IntVar(&config.Connections, "conn", config.Connections, "")
	flags.StringVar(&config.Storage.StorageType, "storage", config.Storage.StorageType, "")
	flags.StringVar(&config.Storage.Host, "shost", config.Storage.Host, "")
	flags.IntVar(&config.Storage.Port, "sport", config.Storage.Port, "")
	flags.StringVar(&config.Storage.IndexBase, "base", config.Storage.IndexBase, "")
	flags.StringVar(&config.Storage.TypeTable, "table", config.Storage.TypeTable, "")
	flags.Parse(os.Args[1:])

	return mRecord, nil
}
