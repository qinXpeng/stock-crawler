package module

import (
	"io/ioutil"
	"net/http"
)

func NetGet(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range GlobalConfig.HttpHeader {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
