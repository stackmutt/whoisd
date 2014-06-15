package dummy

type DummyRecord struct {
	Host      string
	Port      int
	IndexBase string
	TypeTable string
}

func (dummy *DummyRecord) Search(name string, query string) (map[string][]string, error) {
	data := make(map[string][]string)

	result, err := dummy.SearchRaw(dummy.TypeTable, name, query)
	if err != nil {
		return data, err
	}
	if len(result) > 0 {
		return result[0], nil
	}

	return data, nil
}

func (dummy *DummyRecord) SearchRelated(typeTable string, name string, query string) (map[string][]string, error) {
	data := make(map[string][]string)

	result, err := dummy.SearchRaw(typeTable, name, query)
	if err != nil {
		return data, err
	}
	if len(result) > 0 {
		return result[0], nil
	}

	return data, nil
}

func (dummy *DummyRecord) SearchMultiple(typeTable string, name string, query string) (map[string][]string, error) {
	data := make(map[string][]string)

	result, err := dummy.SearchRaw(typeTable, name, query)
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

func (dummy *DummyRecord) SearchRaw(typeTable string, name string, query string) ([]map[string][]string, error) {
	data := make([]map[string][]string, 0)
	for _, result := range dummyData[typeTable] {
		for _, item := range result[name] {
			if item == query {
				data = append(data, result)
				break
			}
		}
	}

	return data, nil
}
