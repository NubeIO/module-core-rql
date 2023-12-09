package apirules

import (
	"errors"
	"fmt"
	"github.com/NubeIO/module-core-rql/helpers/float"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"time"
)

func (inst *RQL) PointHistByDate(hostIDName, pointUUID, startTime, endTime string) any {
	parsedStartTime, err := inst.parseDateTime(startTime)
	if err != nil {
		//stats.Err = err
		return err
	}
	parsedEndTime, err := inst.parseDateTime(endTime)
	if err != nil {
		//stats.Err = err
		return err
	}
	resp, err := cli.GetPointHistoriesWithInterval(hostIDName, pointUUID, parsedStartTime.Format(time.RFC3339), parsedEndTime.Format(time.RFC3339))
	if err != nil {
		return err
	}
	return resp
}

func (inst *RQL) PointHist(hostIDName, pointUUID string) any {
	resp, err := cli.GetPointHistories(hostIDName, pointUUID)
	if err != nil {
		return err
	}
	return resp
}

// CheckPresentValueChange is a function that checks if the present value change of data points within a specific time range exceeds a threshold.
// It takes an array of Hist data points, a start time, an end time, and a threshold as parameters.
// It returns a boolean value indicating whether the present value change exceeds the threshold.
// The function starts by parsing the start and end times into time.Time objects.
// It then initializes a variable prevValue to store the previous present value.
// The function iterates through each data point in the dataPoints array.
// For each data point, it parses the timestamp into a time.Time object.
// If the parsed timestamp is after the start time and before the end time,
// the function checks if the prevValue is 0.0 (indicating the first data point in the time range),
// if so, it assigns the present value of the data point to prevValue and continues to the next iteration.
// Otherwise, it calculates the absolute difference between the present value of the data point and prevValue,
// and checks if it exceeds the threshold.
// If it does, the function returns true.
// If not, it updates prevValue with the present value of the data point.
// After iterating through all the data points, if no present value change exceeds the threshold, the function returns false.
func (inst *RQL) CheckPresentValueChange(dataPoints []model.PointHistory, startTime string, endTime string, threshold float64) bool {
	parsedStartTime, _ := time.Parse(time.RFC3339, startTime)
	parsedEndTime, _ := time.Parse(time.RFC3339, endTime)

	var prevValue float64
	for _, point := range dataPoints {
		parsedPointTime, _ := time.Parse(time.RFC3339, fmt.Sprint(point.Timestamp))
		if parsedPointTime.After(parsedStartTime) && parsedPointTime.Before(parsedEndTime) {
			pointValue := float.NonNil(point.Value)
			if prevValue == 0.0 {
				prevValue = pointValue
				continue
			}
			if abs(pointValue-prevValue) > threshold {
				return true
			}
			prevValue = pointValue
		}
	}
	return false
}

// Stats holds the statistical results
type Stats struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Avg  float64 `json:"avg"`
	Diff float64 `json:"diff"`
	Err  error
}

// PointHistStats checks if the present value for a point changes over a given date-time range by a given threshold and also calculates statistics
func (inst *RQL) PointHistStats(dataPoints []model.PointHistory, startTime string, endTime string) Stats {
	stats := Stats{}
	if len(dataPoints) <= 0 {
		return Stats{
			Err: errors.New("no data was passed in"),
		}
	}
	parsedStartTime, err := inst.parseDateTime(startTime)
	if err != nil {
		stats.Err = err
		return stats
	}
	parsedEndTime, err := inst.parseDateTime(endTime)
	if err != nil {
		stats.Err = err
		return stats
	}
	var prevValue float64
	stats = Stats{Min: float.NonNil(dataPoints[0].Value), Max: float.NonNil(dataPoints[0].Value), Diff: 0.0}
	totVal, cntVal := 0.0, 0.0 // to calculate average
	for _, point := range dataPoints {
		ts := point.Timestamp.String()
		parsedPointTime, err := inst.parseDateTime(ts)
		if err != nil {
			stats.Err = err
			return stats
		}
		if parsedPointTime.After(parsedStartTime) && parsedPointTime.Before(parsedEndTime) {
			pointValue := float.NonNil(point.Value)
			if prevValue == 0.0 {
				prevValue = pointValue
				continue
			}
			prevValue = pointValue
			if pointValue < stats.Min {
				stats.Min = pointValue
			}
			if pointValue > stats.Max {
				stats.Max = pointValue
			}
			totVal += pointValue
			cntVal++
		}
	}
	stats.Diff = stats.Max - stats.Min
	stats.Avg = totVal / cntVal
	return stats
}

// main code stays same as the previous example

// Helper function to calculate the absolute of a float64
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

//
//// Filters data points to those within a given date-time range
//func FilterDataByDateRange(dataPoints []DataPoint, startTime string, endTime string) []DataPoint {
//	parsedStartTime, _ := time.Parse(time.RFC3339, startTime)
//	parsedEndTime, _ := time.Parse(time.RFC3339, endTime)
//
//	var filteredDataPoints []DataPoint
//	for _, point := range dataPoints {
//		parsedPointTime, _ := time.Parse(time.RFC3339, point.Timestamp)
//		if parsedPointTime.After(parsedStartTime) && parsedPointTime.Before(parsedEndTime) {
//			filteredDataPoints = append(filteredDataPoints, point)
//		}
//	}
//	return filteredDataPoints
//}
