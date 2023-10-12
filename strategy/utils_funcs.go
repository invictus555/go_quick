package strategy

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func ruleGroups2JSON(rules []*RuleGroupInfo) (string, error) {
	body, err := json.Marshal(rules)
	return string(body), err
}

func doHttpPost(url, jwt, env string, body []byte) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-tt-env", env)
	req.Header.Set("x-devsre-authorization", jwt)
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
