package curp

import (
	"curp-scraper/api"
	"encoding/json"
	"log"
	"net/http"
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		curp := r.URL.Query().Get("curp")
		if len(curp) != 18 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(api.Response{
				Error: &api.Error{
					Code:    "INVALID_CURP",
					Message: "CURP must be 18 characters long",
				},
			})
			return
		}

		if cachedRes, ok := GetFromCache(curp); ok {
			log.Printf("Returning cached data for CURP: %s", curp)
			if cachedRes.Data != nil {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
			json.NewEncoder(w).Encode(cachedRes)
			return
		}

		model, err := Fetch(curp)
		res := api.Response{}
		if err != nil {
			code := http.StatusInternalServerError
			if e, ok := err.(*Error); ok && e.Code == "CURP_NOT_FOUND" {
				code = http.StatusNotFound
			}
			w.WriteHeader(code)

			res.Error = &api.Error{
				Code:    err.(*Error).Code,
				Message: err.(*Error).Message,
			}

			json.NewEncoder(w).Encode(res)
			return
		}

		res.Data = model
		SetToCache(curp, res)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	})
}
