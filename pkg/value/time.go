package value

import (
	"strconv"
)

const (
	timeUnitHour        = "h"
	timeUnitMinute      = "m"
	timeUnitSecond      = "s"
	timeUnitMilliSecond = "ms"
	timeUnitMicroSecond = "us"
	timeUnitNanoSecond  = "ns"
)

type TimeRes struct {
	Hour        int64
	Minute      int64
	Second      int64
	Millisecond int64
	Microsecond int64
	NanoSecond  int64
}

// TimeValueFormat return hunman readable string format of time value with `nano second` offset.
func TimeValueFormat(timeValue int64) string {
	var hourStr, minuteStr, secondStr, millSecondStr, microSecondStr, nanoSecondStr string

	res := TimeValueHandle(timeValue)

	if res.Hour > 0 {
		hourStr = strconv.FormatInt(res.Hour, 10) + timeUnitHour
	}

	if res.Minute > 0 {
		minuteStr = strconv.FormatInt(res.Minute, 10) + timeUnitMinute
	}

	if res.Second > 0 {
		secondStr = strconv.FormatInt(res.Second, 10) + timeUnitSecond
	}

	if res.Millisecond > 0 {
		millSecondStr = strconv.FormatInt(res.Millisecond, 10) + timeUnitMilliSecond
	}

	if res.Microsecond > 0 {
		microSecondStr = strconv.FormatInt(res.Microsecond, 10) + timeUnitMicroSecond
	}

	nanoSecondStr = strconv.FormatInt(res.NanoSecond, 10) + timeUnitNanoSecond

	timeValueStr := hourStr + minuteStr + secondStr + millSecondStr + microSecondStr + nanoSecondStr

	return timeValueStr
}

// TimeValueHandle handle time value with `nano second` unit, return `hour`, `minute`, `second`, 'millsecond' 'microsecond', 'nanosecond' offset.
func TimeValueHandle(timeValue int64) TimeRes {
	var hourValue, minValue, secValue, milliSecValue, microSecValue, nanoSecValue, remainTime int64
	var resTime TimeRes

	switch {
	case timeValue < 0:
		nanoSecValue = 0

	case timeValue < 1e3:
		nanoSecValue = timeValue

	case timeValue < 1e6:
		microSecValue = timeValue / 1e3
		nanoSecValue = timeValue - microSecValue * 1e3

	case timeValue < 1e9:
		milliSecValue = timeValue / 1e6
		remainTime = timeValue - milliSecValue * 1e6
		resTime = TimeValueHandle(remainTime)
		microSecValue = resTime.Microsecond
		nanoSecValue = resTime.NanoSecond

	case timeValue < 60 * 1e9:
		secValue = timeValue / 1e9
		remainTime = timeValue - secValue * 1e9
		resTime = TimeValueHandle(remainTime)
		milliSecValue = resTime.Millisecond
		microSecValue = resTime.Microsecond
		nanoSecValue = resTime.NanoSecond

	case timeValue < 3600 * 1e9:
		minValue = timeValue / (60 * 1e9)
		remainTime = timeValue - minValue * 60 * 1e9
		resTime = TimeValueHandle(remainTime)
		secValue = resTime.Second
		milliSecValue = resTime.Millisecond
		microSecValue = resTime.Microsecond
		nanoSecValue = resTime.NanoSecond

	default:
		hourValue = timeValue / (3600 * 1e9)
		remainTime = timeValue - hourValue * 3600 * 1e9
		resTime = TimeValueHandle(remainTime)
		minValue = resTime.Minute
		secValue = resTime.Second
		milliSecValue = resTime.Millisecond
		microSecValue = resTime.Microsecond
		nanoSecValue = resTime.NanoSecond

	}

	return TimeRes{
		Hour:   hourValue,
		Minute: minValue,
		Second: secValue,
		Millisecond: milliSecValue,
		Microsecond: microSecValue,
		NanoSecond: nanoSecValue,
	}
}
