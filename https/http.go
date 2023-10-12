package https

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"code.byted.org/security/zti-jwt-helper-golang/helper"
)

// DoHttpGetWithBasicAuth GET
func DoHttpGetWithBasicAuth(url, appID, appSecret, namespace string, body []byte) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	token, err := helper.GetJwtSVID()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(appID, appSecret)
	req.Header.Set("ns", namespace)
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("X-App-Version", "v2")
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
