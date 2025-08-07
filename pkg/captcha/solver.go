package captcha

import (
	"fmt"

	"os"

	capsolver_go "github.com/capsolver/capsolver-go"
)

var capSolver = capsolver_go.CapSolver{
	ApiKey: os.Getenv("CAPSOLVER_API_KEY"),
}

type CaptchaSolution struct {
	Token     string
	UserAgent string
	Cookie    string
}

type CaptchaAction struct {
	Name   string
	Reload string
	Anchor string
}

func Solve(action CaptchaAction, cookie string) (*CaptchaSolution, error) {

	task := map[string]any{
		"type":       "ReCaptchaV3EnterpriseTaskProxyless",
		"websiteURL": "https://www.gob.mx",
		"websiteKey": "6Lfi0jcpAAAAAPfBiQkGzQR3gv8mDRkqPDHAy8hS",
		"pageAction": action.Name,
		"reload":     action.Reload,
		"anchor":     action.Anchor,
	}

	if cookie != "" {
		task["cookie"] = cookie
	}

	res, err := capSolver.Solve(task)
	if err != nil {
		return nil, fmt.Errorf("failed to solve captcha: %w", err)
	}

	// parsedData, err := json.MarshalIndent(res, "", "  ")
	// if err != nil {
	// 	log.Fatalf("could not format JSON: %v", err)
	// }
	// log.Printf("Extracted data:\n%s", string(parsedData))

	return &CaptchaSolution{
		Token:     res.Solution.GRecaptchaResponse,
		UserAgent: res.Solution.UserAgent,
		Cookie:    res.Solution.Cookie,
	}, nil
}
