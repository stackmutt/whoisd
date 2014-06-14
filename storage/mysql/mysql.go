package mysql

import (
	"errors"
)

type MysqlRecord struct {
	Host     string
	Port     int
	DataBase string
	Table    string
}

// TODO - Mysql storage is not released
func (mysql *MysqlRecord) Search(name string, query string) (map[string][]string, error) {
	result := make(map[string][]string)

	return result, errors.New("Mysql driver not released")
}

// TODO - Mysql storage is not released
func (mysql *MysqlRecord) SearchRelated(typeTable string, name string, query string) (map[string][]string, error) {
	result := make(map[string][]string)

	return result, errors.New("Mysql driver not released")
}

// TODO - Mysql storage is not released
func (mysql *MysqlRecord) SearchMultiple(typeTable string, name string, query string) (map[string][]string, error) {
	result := make(map[string][]string)

	return result, errors.New("Mysql driver not released")
}
