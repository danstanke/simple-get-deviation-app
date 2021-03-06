package model

import (
	"math"
	"sort"
)

//data for deviation counting
type dataset []int

type Deviation struct {
	Stddev float64 `json:"stddev"`
	Data   dataset `json:"data"`
}

type DeviationsResult struct {
	Deviations  []Deviation `json:"deviations"`
	AllData     dataset     `json:"allData"`
	StddevOfSum float64     `json:"stddevOfSum"`
}

//sets data for deviation
func (d *Deviation) SetData(data *[]int) {
	d.Data = *data
}

//counts deviation from dataset
func (data *dataset) countDeviaton() float64 {
	sort.Ints(*data)
	sum := 0
	for _, integer := range *data {
		sum += integer
	}

	mean := float64(sum) / float64(len(*data))
	var squares float64 = 0.0
	var lastSquare float64 = 0.0

	for i, integer := range *data {
		//no need to count the same square
		if i > 0 {
			//so previous one is compared to current one
			if (*data)[i-1] != (*data)[i] {
				lastSquare = pow2(mean - float64(integer))
				squares += lastSquare
			} else {
				squares += lastSquare
			}
		} else {
			lastSquare = pow2(float64(integer) - mean)
			squares += lastSquare
		}
	}

	return math.Sqrt(squares / float64(len(*data)))

}

//pow function to avoid using math.Pow
func pow2(x float64) float64 {
	return x * x
}

//counts deviation for Deviation type
func (deviation *Deviation) Count() {
	*deviation = Deviation{
		Stddev: deviation.Data.countDeviaton(),
		Data:   deviation.Data,
	}
}

//counts all deviations and deviation of sum
func (deviationsResult *DeviationsResult) Count(deviations *[]Deviation) {
	var sumOfDatasets dataset
	for _, deviation := range *deviations {
		deviation.Stddev = deviation.Data.countDeviaton()
		sumOfDatasets = append(sumOfDatasets, deviation.Data...)
	}

	*deviationsResult = DeviationsResult{
		Deviations:  *deviations,
		AllData:     sumOfDatasets,
		StddevOfSum: sumOfDatasets.countDeviaton(),
	}
}
