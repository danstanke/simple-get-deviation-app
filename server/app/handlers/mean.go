package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/danstanke/simple-get-deviation-app/server/app/config"
	"github.com/danstanke/simple-get-deviation-app/server/app/model"
	"github.com/danstanke/simple-get-deviation-app/server/app/randorg"
)

//errors
const (
	MissingParameterError        string = "missing parameter"
	IncorrectParameterValueError string = "incorrect parameter value"
	OverTheLimitError            string = "number of request is over the limit"
)

type DeviationResult struct {
	deviation *model.Deviation
	err       error
}

//handles random/mean get request
func RandomMeanHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	numberOfRequests, numberOfIntegers, err := checkAndGetParameters(&query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dataChan := make(chan DeviationResult)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < *numberOfRequests; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dataChan <- getDeviation(numberOfIntegers)
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	var deviations []model.Deviation

	for deviationResult := range dataChan {
		if deviationResult.err != nil {
			http.Error(w, deviationResult.err.Error(), http.StatusInternalServerError)
			return
		} else {
			deviations = append(deviations, *deviationResult.deviation)
		}
	}

	var deviationsResult model.DeviationsResult
	deviationsResult.Count(&deviations)

	deviationsJSON, err := json.Marshal(deviationsResult)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(deviationsJSON)
	return
}

//checks and converts parameters
func checkAndGetParameters(query *url.Values) (*int, *int, error) {
	length := query.Get("length")
	requests := query.Get("requests")

	if length == "" || requests == "" {
		return nil, nil, errors.New(MissingParameterError)
	}

	numberOfRequests, err := strconv.Atoi(requests)

	if err != nil {
		return nil, nil, errors.New(IncorrectParameterValueError)
	}

	if numberOfRequests < 1 || numberOfRequests > config.RequestsLimit {
		return nil, nil, errors.New(OverTheLimitError)
	}

	//numberOfIntegers is checked in model.Query
	numberOfIntegers, err := strconv.Atoi(length)

	if err != nil {
		return nil, nil, errors.New(IncorrectParameterValueError)
	}

	return &numberOfRequests, &numberOfIntegers, nil
}

//requests numbers and returns counted deviation
func getDeviation(numberOfIntegers *int) DeviationResult {
	integers, err := randorg.GetIntegers(numberOfIntegers)
	if err != nil {
		return DeviationResult{nil, err}
	}

	var deviation model.Deviation
	deviation.SetData(&integers)
	deviation.Count()
	return DeviationResult{&deviation, nil}
}
