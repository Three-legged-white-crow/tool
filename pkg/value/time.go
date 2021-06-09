package value

import (
	"strconv"
)

const (
	timeUnitHour   = "h"
	timeUnitMinute = "m"
	timeUnitSecond = "s"
)

type TimeRes struct {
	Hour   int64
	Minute int64
	Second int64
}

// TimeValueFormat return hunman readable string format of time value with `second` offset.
func TimeValueFormat(timeValue int64) string {
	var hourStr, minuteStr, secondStr string

	res := TimeValueHandle(timeValue)

	if res.Hour > 0 {
		hourStr = strconv.FormatInt(res.Hour, 10) + timeUnitHour
	}

	if res.Minute > 0 {
		minuteStr = strconv.FormatInt(res.Minute, 10) + timeUnitMinute
	}

	secondStr = strconv.FormatInt(res.Second, 10) + timeUnitSecond

	timeValueStr := hourStr + minuteStr + secondStr

	return timeValueStr
}

// TimeValueHandle handle time value with `second` unit, return `hour` offset, `minute` offset, `second` offset.
func TimeValueHandle(timeValue int64) TimeRes {
	var hourValue, minValue, secValue int64

	switch {
	case timeValue < 0:
		secValue = 0

	case timeValue < 60:
		secValue = timeValue

	case timeValue < 3600:
		minValue = timeValue / 60
		remainTime := timeValue - 60*minValue
		secValue = TimeValueHandle(remainTime).Second

	default:
		hourValue = timeValue / 3600
		remainTime := timeValue - 3600*hourValue
		minValue = TimeValueHandle(remainTime).Minute
		if minValue > 0 {
			remainTime = remainTime - minValue*60
		}
		secValue = TimeValueHandle(remainTime).Second

	}

	return TimeRes{
		Hour:   hourValue,
		Minute: minValue,
		Second: secValue,
	}
}
