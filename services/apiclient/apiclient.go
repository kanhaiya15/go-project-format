package apiclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
)

// GetService GetService
func GetService(url string, RequestParams ...map[string]string) (responseData interface{}, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return responseData, err
	}
	if len(RequestParams) > 0 {
		for k, v := range RequestParams[0] {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return responseData, err
	}

	json.NewDecoder(resp.Body).Decode(&responseData)
	return responseData, err
}

// PostService PostService
func PostService(url string, payload interface{}, RequestParams ...map[string]string) (responseData interface{}, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	newData, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(newData))
	if err != nil {
		return responseData, err
	}
	if len(RequestParams) > 0 {
		for k, v := range RequestParams[0] {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return responseData, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&responseData)
	return responseData, err
}

// PutService PutService
func PutService(url string, payload interface{}, RequestParams ...map[string]string) (responseData interface{}, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	newData, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(newData))
	if err != nil {
		return responseData, err
	}
	if len(RequestParams) > 0 {
		for k, v := range RequestParams[0] {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return responseData, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&responseData)
	return responseData, err
}

// DeleteService DeleteService
func DeleteService(url string, RequestParams ...map[string]string) (responseData interface{}, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return responseData, err
	}
	if len(RequestParams) > 0 {
		for k, v := range RequestParams[0] {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return responseData, err
	}

	json.NewDecoder(resp.Body).Decode(&responseData)
	return responseData, err
}
