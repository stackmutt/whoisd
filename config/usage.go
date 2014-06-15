package config

var usage = `
whoisd - Whois Daemon

Usage:
  whoisd -option
  whoisd -h | --help
  whoisd -v | --version

Options:
  -h --help         Show this screen
  -v --version      Show version
  -config=<path>    Path to configuration file
  -mapper=<path>    Path to mapper file
  -host=<host/IP>   Host name or IP address
  -port=<port>      Port number
  -work=<number>    Number of active workers (default 1000)
  -conn=<number>    Number of active connections (default 1000)
  -storage=<type>   Type of storage (Elasticsearch, Mysql or Dummy for testing)
  -shost=<host/IP>  Storage host name or IP address
  -sport=<port>     Storage port number
  -base=<name>      Storage index or database name
  -table=<name>     Storage type or table name
`

func Usage() string {
	return usage
}
