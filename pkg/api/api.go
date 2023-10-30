package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Fetches the data with method GET and runs it through a json unmarshaler
func Get(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("GetJson: failed to create request %s: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GetJson: response did not return 'OK': %w", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("GetJson: failed to read response body: %w", err)
	}

	err = json.Unmarshal(data, &out)
	if err != nil {
		return fmt.Errorf("GetJson: failed to unmarshal response body: %w", err)
	}

	return nil
}

// This method only returns the raw response body bytes
func GetRaw(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Get: failed to create request %s: %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get: response did not return 'OK': %w", err)
	}

	bin, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Get: failed reading body: %w", err)
	}

	return bin, nil
}
