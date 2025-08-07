package health

import (
	"curp-scraper/api"
	"curp-scraper/pkg/captcha"
	"encoding/json"
	"net/http"
)

// HealthHandler handles health check requests
func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		solver := captcha.Init() // Initialize the captcha solver

		json.NewEncoder(w).Encode(api.Response{
			Data: map[string]string{
				"status":  "healthy",
				"balance": solver.Balance(),
			},
		})
	})
}
