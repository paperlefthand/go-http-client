package client

import "net/http"

func GetStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	return code, nil
}