package randorg

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/danstanke/simple-get-deviation-app/server/app/config"
	"github.com/danstanke/simple-get-deviation-app/server/app/model"
)

//errors
const TimeoutReachedError string = "Timeout reached"

//gets request url
func getRequestURL(numberOfIntegers int, minimumValue int, maximumValue int) (*string, error) {
	query := model.NewQuery(numberOfIntegers, minimumValue, maximumValue)

	if urlQuery, err := query.GetQuery(); err != nil {
		return nil, err
	} else {
		url := config.RandomOrgURL + *urlQuery
		return &url, nil
	}

}

//parses response into int slice
func parseResponse(responseBody []byte) ([]int, error) {
	var numbers []int
	for _, str := range strings.Split(string(responseBody), "\n") {
		if str != "" {
			if i, err := strconv.Atoi(str); err != nil {
				return nil, err
			} else {
				numbers = append(numbers, i)
			}
		}
	}

	return numbers, nil
}

//gets numbers form RandomOrgURL with context timeout
func GetIntegers(numberOfIntegers *int) ([]int, error) {

	url, err := getRequestURL(*numberOfIntegers, config.MinimumValue, config.MaximumValue)
	if err != nil {
		return nil, err
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, *url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(TimeoutReachedError)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	numbers, err := parseResponse(body)
	if err != nil {
		return nil, err
	}

	return numbers, nil
}
