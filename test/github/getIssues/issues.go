package getIssues

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const token = `9a7d1bda80e539f1ced7c6a1e0b7ffa7bbd64195`

var _url = "https://api.github.com/repos/ShiChao1996/myBlog/issues/events?access_token="

func GetIssues() (string, error) {
	resp, err := http.Get(_url + url.QueryEscape(token))
	if err != nil {
		log.Println(err)
		return "", err
	}

	s, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(s), nil
}
