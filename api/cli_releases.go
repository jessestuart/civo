package api

import (
	"errors"
	"fmt"

	"github.com/jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

func makeGitHubCall(url string) (*gabs.Container, error) {
	request := gorequest.New()

	request = request.Get(url)
	resp, body, errs := request.End()
	if errs != nil {
		return nil, errs[0]
	}
	status := resp.StatusCode

	if !(status >= 200 && status <= 299) {
		return nil, HTTPErrorNew(fmt.Sprintf("Unable to make GitHub API call: %s", body), url, status)
	}
	if body != "" {
		jsonObject, err := gabs.ParseJSON([]byte(body))
		if err != nil {
			return nil, err
		}
		return jsonObject, nil
	}
	return nil, nil

}

func LatestRelease() (string, error) {
	result, err := makeGitHubCall("https://api.github.com/repos/absolutedevops/civo/releases/latest")
	if err != nil {
		return "", err
	}

	release, ok := result.Path("tag_name").Data().(string)
	if ok != true {
		return "", errors.New("Unable to find tag_name in github latest releases feed")
	}

	return release, nil
}
