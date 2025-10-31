package convertors

import "encoding/json"

func MapToString(mapData map[string]string) string {
	str, _ := json.Marshal(mapData)
	return string(str)
}

func StringToMap(stringData string) map[string]string {
	var data map[string]string
	_ = json.Unmarshal([]byte(stringData), &data)
	return data
}

func UrlValuesToString(mapData map[string][]string) string {
	str, _ := json.Marshal(mapData)
	return string(str)
}

func StringToUrlValues(stringData string) map[string][]string {
	var data map[string][]string
	_ = json.Unmarshal([]byte(stringData), &data)
	return data
}
