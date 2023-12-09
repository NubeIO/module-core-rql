package apirules

import (
	"fmt"
	"github.com/NubeIO/module-core-rql/helpers"

	"time"
)

// Package documentation for code containing date and time formats.
// These constants represent commonly used formats for date and time representation.
// The formats include the full date and time, time with milliseconds, time only, date only, year only,
// day of the week only, and the full date and time with the day of the week.
const (
	FullFormat    = "2006-01-02 15:04:05"
	TimeWithMS    = "15:04:05.000"
	TimeFormat    = "15:04:05"
	DateFormat    = "2006.01.02"
	YearFormat    = "2006"
	DayFormat     = "Monday"
	FullDayFormat = "2006-01-02 15:04:05 Monday"
)

// SubtractYears returns a new `time.Time` value that is `years` years before the given `t`.
// Example usage:
//   result := inst.SubtractYears(t, years)
func (inst *RQL) SubtractYears(t time.Time, years int) time.Time {
	return t.AddDate(-years, 0, 0)
}

// SubtractDays subtracts the specified number of days from the given time and returns the resulting time.
// Example usage:
//   result := inst.SubtractDays(t, days)
//
// Parameters:
// - t: The original time from which days are subtracted.
// - days: The number of days to subtract.
//
// Returns:
// The resulting time after subtracting the specified number of days.
func (inst *RQL) SubtractDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, -days)
}

// SubtractHours subtracts the specified number of hours from the given time and returns the resulting time.
//
// Example usage:
//   result := inst.SubtractHours(t, hours)
//
// Parameters:
// - t: the original time from which hours are subtracted.
// - hours: the number of hours to subtract.
func (inst *RQL) SubtractHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(-hours) * time.Hour)
}

// SubtractMinutes subtracts the specified number of minutes from the given time and returns the resulting time.
// Example usage:
//   result := inst.SubtractMinutes(t, minutes)
// where:
// - t is the original time from which minutes are subtracted.
// - minutes is the number of minutes to subtract.
func (inst *RQL) SubtractMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(-minutes) * time.Minute)
}

// SubtractSeconds subtracts the specified number of seconds from the given time and returns the resulting time.
//
// Example usage:
//
//     result := inst.SubtractSeconds(t, seconds)
//
// where:
//
// - t is the original time from which seconds are subtracted.
// - seconds is the number of seconds to subtract.
func (inst *RQL) SubtractSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(-seconds) * time.Second)
}

// AddYears adds the specified number of years to the given time and returns the resulting time.
// Example usage:
//   result := inst.AddYears(t, years)
// where:
// - t is the original time to which years are added.
// - years is the number of years to add.
func (inst *RQL) AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// AddDays returns a new Time instance that is `days` days after `t`.
func (inst *RQL) AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddHours adds the specified number of hours to the given time and returns the resulting time.
// Example usage:
//   result := inst.AddHours(t, hours)
//
// Parameters:
// - t is the original time to which hours are added.
// - hours is the number of hours to add.
func (inst *RQL) AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddMinutes adds the specified number of minutes to the given time and returns the resulting time.
// Example usage:
//   result := inst.AddMinutes(t, minutes)
// where:
// - t is the original time to which minutes are added.
// - minutes is the number of minutes to add.
func (inst *RQL) AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddSeconds adds the specified number of seconds to the given time and returns the resulting time.
// Example usage:
//   result := inst.AddSeconds(t, seconds)
// where:
// - t is the original time to which seconds are added.
// - seconds is the number of seconds to add.
func (inst *RQL) AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// TimeSince calculates the time duration that has elapsed since the given time and returns it as a string representation of the time duration in a human-readable format.
// Example usage:
//   result := inst.TimeSince(t)
// where:
// - t is the original time for which the time duration is calculated.
// The returned string represents the time duration in the following format:
// - "now" if the time duration is zero or negative.
// - "%d seconds ago" if the time duration is less than a minute.
// - "%d minutes ago" if the time duration is less than an hour.
// - "%d hours ago" if the time duration is less than a day.
// - "%d days ago" if the time duration is more than a day.
// The function uses the time.Since() function from the time package to calculate the time duration since the given time.
// It then converts the time duration to seconds and checks against different thresholds to determine the appropriate string representation.
// The time duration is rounded down to the nearest whole unit, i.e., seconds, minutes, hours, or days.
func (inst *RQL) TimeSince(t time.Time) string {
	delta := time.Since(t)
	total := int(delta.Seconds())

	if total <= 0 {
		return "now"
	} else if total < 60 { // less than a minute
		return fmt.Sprintf("%d seconds ago", total)
	} else if total < 3600 { // less than an hour
		return fmt.Sprintf("%d minutes ago", total/60)
	} else if total < 86400 { // less than a day
		return fmt.Sprintf("%d hours ago", total/3600)
	} else { // more than a day
		return fmt.Sprintf("%d days ago", total/86400)
	}
}

// currentTime returns the current time in UTC format.
// Example usage:
//   result := inst.currentTime()
//
// This method can be used to get the current time in UTC format.
// Without any parameters, it will return the current time with
// precision up to nanoseconds.
//
// Returns the current time in UTC format.
func (inst *RQL) currentTime() time.Time {
	return time.Now().UTC()
}

// formattedCurrentTime returns the current time formatted according to the provided format string.
// The format string follows the same layout as the one used in the time.Format method.
// Example usage:
//   result := inst.formattedCurrentTime(format)
// where:
// - inst is an instance of the RQL struct.
// - format is the desired format string for the time representation.
//   Valid format strings are defined in the constants FullFormat, TimeWithMS, TimeFormat, DateFormat, YearFormat, DayFormat, and FullDayFormat.
//   These constants can be used as reference for common format requirements.
//   Example usage: FullFormat, TimeWithMS, TimeFormat, DateFormat, YearFormat, DayFormat, FullDayFormat.
func (inst *RQL) formattedCurrentTime(format string) string {
	return inst.currentTime().Format(format)
}

// DateTimeUTC returns the current time in UTC format.
// Example usage:
//   result := inst.TimeUTC()
// where:
// - inst is the receiver instance of type RQL.
// The returned time is in Coordinated Universal Time (UTC).
func (inst *RQL) DateTimeUTC() time.Time {
	return inst.currentTime()
}

// DateTime returns the current local time as a new `time.Time` value.
// Example usage:
//   result := inst.DateTime()
func (inst *RQL) DateTime() time.Time {
	return time.Now()
}

// TimeDate returns the current time in the format "2006-01-02 15:04:05".
// Example usage:
//   result := inst.TimeDate()
// No parameters needed.
// Returns the current time as a string in the specified format.
func (inst *RQL) TimeDate() string {
	return inst.formattedCurrentTime(FullFormat)
}

// TimeWithMilliseconds returns the current time in the format "15:04:05.000" with milliseconds.
func (inst *RQL) TimeWithMilliseconds() string {
	return inst.formattedCurrentTime(TimeWithMS)
}

// Time returns the current time in the specified format.
// It uses the formattedCurrentTime method to format the time.
// Example usage:
//   result := inst.Time()
// where:
// - result is the current time in string format.
// Without duplicating the example above.
// formattedCurrentTime is a private helper method used internally by the Time method.
// It takes the time format as a parameter and returns the formatted time.
// It uses the TimeFormat constant to format the time.
// Example usage:
//   result := inst.formattedCurrentTime(format)
// where:
// - format is the desired format of the time.
// - result is the formatted time in string format.
func (inst *RQL) Time() string {
	return inst.formattedCurrentTime(TimeFormat)
}

// Date returns the current date in the specified format ("2006.01.02") as a string.
func (inst *RQL) Date() string {
	return inst.formattedCurrentTime(DateFormat)
}

// Year returns the current year formatted as a string using the YearFormat constant.
// This method is a helper function that utilizes the formattedCurrentTime method.
//
// Example usage:
//   result := inst.Year()
//
// Return Value:
//   The current year formatted as a string.
//
// Note: YearFormat is a constant that represents the year format "2006".
func (inst *RQL) Year() string {
	return inst.formattedCurrentTime(YearFormat)
}

// Day returns the current day of the week in the format "Monday".
// It uses the formattedCurrentTime method to retrieve the formatted time.
// Example usage:
//   result := inst.Day()
//
// where:
// - inst is a pointer to an RQL instance.
//   It is used to access the formattedCurrentTime method.
//
// Return value:
// - string: the current day of the week.
func (inst *RQL) Day() string {
	return inst.formattedCurrentTime(DayFormat)
}

// TimeDateDay returns the formatted current time using the FullDayFormat constant.
// Example usage:
//   result := inst.TimeDateDay()
//
// FullDayFormat is a constant representing the format of the date and time:
//    "2006-01-02 15:04:05 Monday".
// The resulting time string will have the following format:
//    YYYY-MM-DD HH:MM:SS Weekday
// where:
//    - YYYY represents the 4-digit year
//    - MM represents the 2-digit month
//    - DD represents the 2-digit day
//    - HH represents the 2-digit hour in 24-hour format
//    - MM represents the 2-digit minute
//    - SS represents the 2-digit second
//    - Weekday represents the full name of the weekday (e.g., Monday, Tuesday, etc.)
// Without duplicating the example above, document the following code:
// ```go
// func (inst *RQL) TimeDateDay() string {
// 	return inst.formattedCurrentTime(FullDayFormat)
// }
// ```
func (inst *RQL) TimeDateDay() string {
	return inst.formattedCurrentTime(FullDayFormat)
}

// ParseDateTime parses the given date string and returns the corresponding time.Time value.
// If the date string is not in a valid format, it returns a zero value of type time.Time.
//
// Example usage:
//   result := inst.ParseDateTime(dateStr)
//
// Parameters:
//   - dateStr: the date string to be parsed.
//
// Returns:
//   - t: the time.Time value corresponding to the parsed date string.
//   - If the date string is not in a valid format, it returns a zero value of type time.Time.
func (inst *RQL) ParseDateTime(dateStr string) any {
	t, err := inst.parseDateTime(dateStr)
	if err != nil {
		return err
	}
	return t
}

// handle error
func (inst *RQL) parseDateTime(dateStr string) (time.Time, error) {
	tt, err := helpers.DateTimeParse(dateStr).TimeIn("UTC")
	if err != nil {
		return time.Time{}, err
	}
	return tt, nil
}

// TimeDiff Define a struct to hold the time difference in various units
type TimeDiff struct {
	seconds float64
	minutes float64
	hours   float64
	days    float64
	years   float64
}

// GetDifference Returns a struct containing time differences in various units
func (inst *RQL) GetDifference(time1, time2 time.Time) TimeDiff {
	diff := time2.Sub(time1)
	return TimeDiff{
		seconds: diff.Seconds(),
		minutes: diff.Minutes(),
		hours:   diff.Hours(),
		days:    diff.Hours() / 24,
		years:   diff.Hours() / (24 * 365.25),
	}
}
