package curp

import (
	"bytes"
	"curp-scraper/pkg/captcha"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Fetch(curp string) (*CurpModel, error) {

	consulta, err := captcha.Solve(captcha.Consulta, "")
	if err != nil {
		return nil, &Error{Code: "CAPTCHA_FAILURE", Message: fmt.Sprintf("failed to get captcha solution: %v", err)}
	}

	reqBody := map[string]string{
		"curp":         curp,
		"tipoBusqueda": "curp",
		"ip":           "127.0.0.1",
		"token":        consulta.Token,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, &Error{Code: "ENCODING_ERROR", Message: err.Error()}
	}

	req, err := http.NewRequest(http.MethodPost, "https://www.gob.mx/v1/renapoCURP/consulta", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, &Error{Code: "REQUEST_CREATION_ERROR", Message: fmt.Sprintf("failed to create request: %v", err)}
	}

	// Use Header map directly for non-canonical keys
	headers := map[string]string{
		"accept": "application/json",
		// "user-agent": consulta.UserAgent,
		// "cookie": cookie,

		// "accept-language":              "es-MX,es;q=0.9", // q stands for quality factor, indicating preference
		"access-control-allow-headers": "Origin, X-Requested-With, Content-Type, Accept",
		"access-control-allow-origin":  "*",
		// "cache-control":                "no-cache",
		// "connection":                   "keep-alive",
		"content-type": "application/json; charset=UTF-8;",
		"dnt":          "1",
		"host":         "www.gob.mx",
		"origin":       "https://www.gob.mx",
		"pragma":       "no-cache",
		"referer":      "https://www.gob.mx/",
		// "sec-fetch-dest":               "empty",
		// "sec-fetch-mode":               "cors",
		// "sec-fetch-site":               "same-origin",
		"x-requested-with": "XMLHttpRequest",
		// "sec-ch-ua":                    `"Not)A;Brand";v="8", "Chromium";v="138"`,
		// "sec-ch-ua-mobile":             "?0",
		// "sec-ch-ua-platform":           `"macOS"`,
	}

	for k, v := range headers {
		req.Header[k] = []string{v}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, &Error{Code: "RENAPO_HTTP_ERROR", Message: fmt.Sprintf("request failed with status %d: %s", resp.StatusCode, body)}
	}

	// Process the response body as needed

	var res curpResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, &Error{Code: "DECODE_ERROR", Message: err.Error()}
	}

	if res.Codigo == "180001" || len(res.Registros) == 0 {
		return nil, &Error{Code: "CURP_NOT_FOUND", Message: "No records found"}
	}

	r := res.Registros[0]
	return &CurpModel{
		Curp:                   r.Curp,
		Nombres:                r.Nombres,
		PrimerApellido:         r.PrimerApellido,
		SegundoApellido:        r.SegundoApellido,
		ClaveGenero:            r.Sexo,
		Genero:                 r.Sexo,
		FechaNacimiento:        r.FechaNacimiento,
		DiaNacimiento:          r.FechaNacimiento[0:2],
		MesNacimiento:          r.FechaNacimiento[3:5],
		AnioNacimiento:         r.FechaNacimiento[6:10],
		ClaveEntidadNacimiento: r.ClaveEntidad,
		EntidadNacimiento:      r.Entidad,
	}, nil
}
