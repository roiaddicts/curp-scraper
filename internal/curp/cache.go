package curp

import (
	"curp-scraper/api"
	"encoding/json"
	"os"
)

func GetFromCache(curp string) (api.Response, bool) {
	res, ok := fromCache[api.Response](curp + ".json")
	if !ok {
		return api.Response{}, false
	}
	return res, true
}

func SetToCache(curp string, res api.Response) {
	cache(curp+".json", res)
}

func cache(filename string, data any) {
	file, _ := os.Create("cache/" + filename)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
	file.Close()
}

func fromCache[T any](filename string) (T, bool) {
	var zero T

	if _, err := os.Stat(filename); err != nil {
		return zero, false
	}

	file, err := os.Open(filename)
	if err != nil {
		return zero, false
	}
	defer file.Close()

	var data T
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return zero, false
	}

	return data, true
}
