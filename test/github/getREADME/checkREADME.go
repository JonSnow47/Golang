package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/fengyfei/gu/libs/logger"
)

func main() {
	e := new(string)
	fmt.Println(getReadme(e))
}

func getReadme(reponame *string) string {
	const (
		head     = "https://raw.githubusercontent.com/master/"
		emptyStr = ""
	)

	var (
		suffix   = []string{"README.md", "Readme.md", "README"}
		bodyByte []byte
	)

	if reponame == nil {
		return emptyStr
	}

	for _, v := range suffix {
		fullURL := head + *reponame + v

		resp, err := http.Get(fullURL)
		if err != nil {
			logger.Debug("getReadme http.Get returned error:", err)
			goto finish
		}

		if resp.StatusCode != http.StatusOK {
			logger.Error(resp.StatusCode)
			goto finish
		}

		bodyByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Debug("getReadme ReadAll returned error:", err)
			goto finish
		}

		resp.Body.Close()
		return string(bodyByte)

	finish:
		resp.Body.Close()
	}
	return emptyStr
}
