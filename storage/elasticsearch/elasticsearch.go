package elasticsearch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ElasticsearchRecord struct {
	Host  string
	Port  int
	Index string
	Type  string
}

// Search the whois data in the storage
func (elastic *ElasticsearchRecord) Search(name string, query string) (map[string][]string, error) {
	data := make(map[string][]string)
	result, err := elastic.SearchRaw(elastic.Type, name, query)
	if err != nil {
		return data, err
	}
	if len(result) > 0 {
		return result[0], nil
	}

	return data, nil
}

// Search the whois data in the storage from related type or table
func (elastic *ElasticsearchRecord) SearchRelated(typeTable string, name string, query string) (map[string][]string, error) {
	data := make(map[string][]string)
	result, err := elastic.SearchRaw(typeTable, name, query)
	if err != nil {
		return data, err
	}
	if len(result) > 0 {
		return result[0], nil
	}

	return data, nil
}

// Search multiple records from the whois data in the storage
func (elastic *ElasticsearchRecord) SearchMultiple(typeTable string, name string, query string) (map[string][]string, error) {
	data := make(map[string][]string)
	result, err := elastic.SearchRaw(typeTable, name, query)
	if err != nil {
		return data, err
	}
	if len(result) > 0 {
		for _, item := range result {
			for key, value := range item {
				data[key] = append(data[key], value...)
			}
		}
		return data, nil
	}

	return data, nil
}

func (elastic *ElasticsearchRecord) SearchRaw(typeTable string, name string, query string) ([]map[string][]string, error) {
	all := make([]map[string][]string, 0)
	url := "http://" + elastic.Host + ":" + strconv.Itoa(elastic.Port) + "/" + elastic.Index + "/" + typeTable
	request := url + "/_search?q=" + name + ":" + query + ""
	response, err := http.Get(request)
	if err != nil {
		return all, err
	}
	jsondata, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return all, err
	}

	type DataRecord struct {
		Took int
		Hits struct {
			Total int
			Hits  []map[string]interface{}
		}
	}

	result := new(DataRecord)
	if err := json.Unmarshal(jsondata, result); err != nil {
		return all, err
	}

	if result.Hits.Total > 0 {
		for _, record := range result.Hits.Hits {
			element := transformData(record)
			all = append(all, element)
		}
	}

	return all, nil
}

// Transformation the data to requested format
func transformData(record map[string]interface{}) map[string][]string {
	element := make(map[string][]string)

	// Check the data on simple and array fields
	// Convert all data to []string
	for index, value := range record["_source"].(map[string]interface{}) {

		// Check for array field
		if arrayFields, ok := value.(map[string]interface{}); ok {

			// Array field must be converted to <name.key> format
			for key, val := range arrayFields {
				var item []string
				if array, ok := val.([]interface{}); ok {
					for _, str := range array {
						if str != nil {
							item = append(item, str.(string))
						}
					}
				} else {
					if val != nil {
						item = []string{val.(string)}
					}
				}
				element[index+"."+key] = item
			}

		} else {

			// Simple field used as <name> and converted to []string
			var item []string
			if str, ok := value.(string); ok {
				item = []string{str}
			}
			element[index] = item
		}
	}

	return element
}
