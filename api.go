package API

import (
	"fmt"
	"github.com/gojektech/heimdall/httpclient"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

// Generic get on the API.  Actually this would work on any url.
func Get(client *httpclient.Client, url string) ([]byte, error) {

	methodName := "bookwerx-common-go:api.go:get"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s :Error with http.NewRequest", methodName))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s :Error with client.Do", methodName))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s :Error with ioutil.ReadAll", methodName))
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s :Error with resp.Body.Close", methodName))
	}

	if resp.StatusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"%s :Status code error: expected= 200, received=%, body=%s", methodName, resp.StatusCode, string(body)))
	}

	return body, nil

}
