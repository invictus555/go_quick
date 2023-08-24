package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"code.byted.org/security/zti-jwt-helper-golang/helper"
	"github.com/pkg/errors"
)

// DoHttpGet GET
func DoHttpGet(url string, timeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	if resp.StatusCode != 200 {
		return body, errors.New("resp is not 200")
	}
	return body, err
}

// DoHttpGetWithBasicAuth GET
func DoHttpGetWithBasicAuth(url, appID, appSecret string, body []byte) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	var tokenString string
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(appID, appSecret)
	req.Header.Set("X-App-Version", "v2")
	tokenString, err = helper.GetJwtSVID()
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", tokenString)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}
	if resp.StatusCode != 200 {
		return ret, fmt.Errorf("resp status is %d, %s. url is %s, app_id=%s, app_secret=%s", resp.StatusCode, url, string(body), appID, appSecret)
	}
	return ret, err
}

// DoHttpPostWithBasicAuth POST
func DoHttpPostWithBasicAuth(url, appID, appSecret string, body []byte) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	var tokenString string
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(appID, appSecret)
	req.Header.Set("X-App-Version", "v2")
	tokenString, err = helper.GetJwtSVID()
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", tokenString)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resBody, err
	}
	if resp.StatusCode != 200 {
		return resBody, errors.New("resp is not 200")
	}
	return resBody, err
}
