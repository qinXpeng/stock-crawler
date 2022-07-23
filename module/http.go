package module

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct{
	Client http.Client
	Header http.Header
}

var GlobalHttp *HttpClient

func AssertHttp(){
	if GlobalHttp == nil {
		client := http.Client{}
		header := http.Header{}
		header.Add("Accept","application/json, text/plain, */*")
		//header.Add("Accept-Encoding", "gzip, deflate, br")
		header.Add("Connection","keep-alive")
		header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
		header.Add("Cookie", GlobalRegister.Headers["Cookie"])
		GlobalHttp = &HttpClient{
			Client: client,
			Header: header,
		}
	}
}


func HttpGet(url string)([]byte, error){
	AssertHttp()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = GlobalHttp.Header
	res, err := GlobalHttp.Client.Do(req)
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
