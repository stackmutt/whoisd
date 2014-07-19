package config

import (
	"encoding/json"
	"errors"
	"flag"
	"os"

	"github.com/takama/whoisd/mapper"
)

// Default values: path to config file, host, port, etc
const (
	defaultConfigPath  = "/etc/whoisd/whoisd.conf"
	defaultMappingPath = "/etc/whoisd/conf.d/mapping.json"
	defaultHost        = "0.0.0.0"
	defaultPort        = 43
	defaultWorkers     = 1000
	defaultConnections = 1000
	defaultStorageType = "Dummy"
	defaultStorageHost = "localhost"
	defaultStoragePort = 9200
	defaultIndexBase   = "whois"
	defaultTypeTable   = "domain"
)

type ConfigRecord struct {
	ConfigPath  string
	MappingPath string

	ShowVersion bool
	TestMode    bool
	TestQuery   string

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

// returns the config initialized with default values
func New() *ConfigRecord {
	config := new(ConfigRecord)
	flag.BoolVar(&config.ShowVersion, "version", false, "show version")
	flag.BoolVar(&config.ShowVersion, "v", false, "show version")
	flag.BoolVar(&config.TestMode, "t", false, "test mode")
	flag.BoolVar(&config.TestMode, "test", false, "test mode")
	flag.StringVar(&config.ConfigPath, "config", defaultConfigPath, "path to configuration file")
	flag.StringVar(&config.MappingPath, "mapping", defaultMappingPath, "path to mapping file")
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

// Loads settings from config file or from sh command line
func (config *ConfigRecord) Load() (*mapper.MapperRecord, error) {
	var path string
	var err error
	mRecord := new(mapper.MapperRecord)

	if err = config.LoadConfigFile(config.ConfigPath); err != nil {
		return nil, err
	}
	if mRecord, err = LoadMappingFile(config.MappingPath); err != nil {
		return nil, err
	}

	// overwrite config from file by cmd flags
	flags := flag.NewFlagSet("whoisd", flag.ContinueOnError)
	// Begin ignored flags
	flags.StringVar(&path, "config", "", "")
	flags.StringVar(&path, "mapping", "", "")
	// End ignored flags
	flags.BoolVar(&config.TestMode, "t", config.TestMode, "")
	flags.BoolVar(&config.TestMode, "test", config.TestMode, "")
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

// loads congig file into config record
func (config *ConfigRecord) LoadConfigFile(path string) error {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	data := make([]byte, stat.Size())
	if _, err := file.Read(data); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}

// loads mapper records and returns it
func LoadMappingFile(path string) (*mapper.MapperRecord, error) {
	record := new(mapper.MapperRecord)
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, errors.New("Mapping file not found, please load it through -mapping option or install command")
	}
	mFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer mFile.Close()
	data := make([]byte, stat.Size())
	if _, err := mFile.Read(data); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	return record, nil
}
