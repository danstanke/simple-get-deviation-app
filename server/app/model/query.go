package model

import (
	"errors"
	"strconv"
)

//errors
const (
	IncorrectMaximumValueError     string = "too large maximum value of integers"
	IncorrectMinimumValueError     string = "too large minimum value of integers"
	MinimumGreaterThanMaximumError string = "minimum value cannot be greater than maximum value"
	IncorrectNumberOfIntegersError string = "incorrect or too large number of integers"
)

//query limits
const (
	MaxNumberOfIntegers int = 100
	MaxMinimumValue     int = 1000
	MaxMaximumValue     int = 1000
	MinMinimumValue     int = -1000
	MinMaximumValue     int = -1000
)

//query for random.org url
type Query struct {
	numberOfIntegers int
	minimumValue     int
	maximumValue     int
}

//creates
func NewQuery(numberOfIntegers int, minimumValue int, maximumValue int) *Query {
	return &Query{numberOfIntegers: numberOfIntegers, minimumValue: minimumValue, maximumValue: maximumValue}
}

//validates if Query has proper values
func (q *Query) validate() error {
	if q.maximumValue > MaxMaximumValue || q.maximumValue < MinMaximumValue {
		return errors.New(IncorrectMaximumValueError)
	}

	if q.minimumValue > MaxMinimumValue || q.minimumValue < MinMinimumValue {
		return errors.New(IncorrectMinimumValueError)
	}

	if q.minimumValue > q.maximumValue {
		return errors.New(MinimumGreaterThanMaximumError)
	}

	if q.numberOfIntegers > MaxNumberOfIntegers || q.numberOfIntegers < 0 {
		return errors.New(IncorrectNumberOfIntegersError)
	}

	return nil
}

//gets query
func (q *Query) GetQuery() (*string, error) {
	if err := q.validate(); err == nil {
		urlQuery := "?num=" +
			strconv.Itoa(q.numberOfIntegers) +
			"&min=" + strconv.Itoa(q.minimumValue) +
			"&max=" + strconv.Itoa(q.maximumValue) +
			"&col=1&base=10&format=plain&rnd=new"
		return &urlQuery, nil
	} else {
		return nil, err
	}
}
