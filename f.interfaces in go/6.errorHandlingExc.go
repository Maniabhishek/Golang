package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour, Minute, Seconds int
}

type TimerParseError struct {
	msg   string
	input string
}

func (tpe *TimerParseError) Error() string {
	return fmt.Sprintf("%v : %v", tpe.msg, tpe.input)
}

func TimerParser(time string) (Time, error) {
	components := strings.Split(time, ":")

	if len(components) != 3 {
		return Time{}, &TimerParseError{"Invalid  string", time}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimerParseError{fmt.Sprintf("error parsing hour %v ", err), time}
		}

		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimerParseError{fmt.Sprintf("error parsing minute %v ", err), time}
		}

		sec, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimerParseError{fmt.Sprintf("error parsing seconds %v ", err), time}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimerParseError{"hour out of range", fmt.Sprintf("%v", hour)}
		}

		return Time{hour, minute, sec}, nil
	}

}
