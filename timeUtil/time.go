package timeUtil

import (
	"time"

	"github.com/dromara/carbon/v2"
)

var (
	currLayout = carbon.Shanghai
	// currLayout = carbon.LosAngeles
)
var (
	Time    = _time{}
	TimePtr = &Time
)

type _time struct {
}

func (t _time) GetCurrLayout() string {
	return currLayout
}

func (t _time) NowCarbon() carbon.Carbon {
	return carbon.Now(currLayout)
}

func (t _time) NowTime() time.Time {
	return t.NowCarbon().StdTime()
}

func (t _time) NowCarbonPtr() *carbon.Carbon {
	now := t.NowCarbon()
	return &now
}

///////////

// 字符串转时间
func (t _time) StrToCarbon(strTime string) carbon.Carbon {

	return carbon.Parse(strTime, currLayout)

}

// 字符串转时间
func (t _time) StrToCarbonPtr(strTime string) *carbon.Carbon {

	dstTime := t.StrToCarbon(strTime)
	if dstTime.Error != nil {
		return nil
	}
	return &dstTime

}

// 字符串转时间
func (t _time) StrPtrToCarbonPtr(strTime *string) *carbon.Carbon {

	if strTime == nil {
		return nil
	}
	return t.StrToCarbonPtr(*strTime)

}

func (t _time) StartOfDay(strTime *string) *carbon.Carbon {

	dstTime := t.StrPtrToCarbonPtr(strTime)
	if dstTime == nil {
		return nil
	}

	startAt := dstTime.StartOfDay()
	return &startAt

}

func (t _time) EndOfDay(strTime *string) *carbon.Carbon {

	dstTime := t.StrPtrToCarbonPtr(strTime)
	if dstTime == nil {
		return nil
	}

	endAt := dstTime.EndOfDay()
	return &endAt
}

// ////////
func (t _time) NowFormatTime(layout string) string {
	return t.NowCarbon().StdTime().Format(layout)
}

// func FormatTime(time carbon.Carbon) string {
// 	return time.StdTime().Format(currLayout)
// }

func (t _time) CarbonToDateTime(time carbon.Carbon) string {
	return time.ToDateTimeString(currLayout)
}

func (t _time) CarbonPtrToDateTimePtr(at *carbon.Carbon) *string {

	if at == nil {
		return nil
	}
	str := t.CarbonToDateTime(*at)
	return &str

}

func (t _time) CarbonToDate(time carbon.Carbon) string {
	return time.ToDateString(currLayout)
}

func (t _time) CarbonPtrToDate(at *carbon.Carbon) string {

	if at == nil {
		return ""
	}
	return t.CarbonToDate(*at)

}

func (t _time) AddDayToDate(at *carbon.Carbon, days int) string {

	if at == nil {
		return ""
	}
	return t.CarbonToDate(at.AddDays(days))

}

// func (t _time) FormatPointerTime(time *carbon.Carbon) *string {
// 	if time == nil {
// 		return nil
// 	}
// 	str := time.ToDateTimeString(currLayout) //.ToDateString(currLayout)
// 	return &str
// }

func (t _time) PStrToPCarbonDate(str *string) *carbon.Carbon {

	time := t.StrPtrToCarbonPtr(str)
	if time == nil {
		return nil
	}
	startAt_tmp := time.StartOfDay()
	return &startAt_tmp

}

func (t _time) NowUnixNano() int64 {
	return t.NowCarbon().StdTime().UnixNano()
}

func (t _time) NowUnixMilli() int64 {
	return t.NowCarbon().StdTime().UnixMilli()
}

func (t _time) NowUnix() int64 {
	return t.NowCarbon().StdTime().Unix()
}

func (t _time) NowNanosecond() int {
	return t.NowCarbon().StdTime().Nanosecond()
}

func (t _time) NowAddSeconds(d int) time.Time {
	return t.NowCarbon().AddSeconds(d).StdTime()
}

func (t _time) NowAddMinutes(d int) time.Time {
	return t.NowCarbon().AddMinutes(d).StdTime()
}

func (t _time) NowAddHours(d int) time.Time {
	return t.NowCarbon().AddHours(d).StdTime()
}

func (t _time) NowAddDays(d int) time.Time {
	return t.NowCarbon().AddDays(d).StdTime()
}

func (t _time) NowAddMonths(d int) time.Time {
	return t.NowCarbon().AddMonths(d).StdTime()
}

func (t _time) TodyRange() (carbon.Carbon, carbon.Carbon) {
	nowCarbon := t.NowCarbon()
	start := nowCarbon.StartOfDay()
	end := nowCarbon.EndOfDay()
	return start, end
}

// 当天开始结束
func (t _time) CurrDayStartEnd() (startAt *carbon.Carbon, endAt *carbon.Carbon) {
	start, end := t.TodyRange()
	return &start, &end
}

// 昨日范围
func (t _time) YdayRange() (carbon.Carbon, carbon.Carbon) {
	yesterday := t.NowCarbon().SubDay()
	start := yesterday.StartOfDay()
	end := yesterday.EndOfDay()
	return start, end
}

// 前月范围
func (t _time) PreMonthRange() (startAt *carbon.Carbon, endAt *carbon.Carbon) {

	end := t.NowCarbon().EndOfDay()
	start := end.SubMonth().StartOfDay()

	startAt = &start
	endAt = &end

	return

}

// ***********************************************************
// 获取时间戳-毫秒级别
// ***********************************************************
func (t _time) GetTimeStampMilsecd() int64 {
	return t.NowUnixNano() / 1e6
}
