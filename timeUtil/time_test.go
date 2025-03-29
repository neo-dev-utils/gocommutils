package timeUtil

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	// {
	// 	fmt.Println(Time.GetCurrLayout())
	// 	fmt.Println(Time.NowCarbon())
	// 	fmt.Println(Time.NowTime())
	// 	fmt.Println(Time.NowCarbonPtr())
	// }

	// {
	// 	fmt.Println(Time.StrToCarbon("2024-03-19 21:26:03"))
	// 	fmt.Println(Time.StrToCarbonPtr("2024-03-19 21:26:03"))
	// 	strTime := "20240319"
	// 	fmt.Println(Time.StrPtrToCarbonPtr(&strTime))
	// }

	// {
	// 	strTime := "2024-03-19 21:26:03"
	// 	fmt.Println(Time.StartOfDay(&strTime))
	// 	fmt.Println(Time.EndOfDay(&strTime))
	// }

	// {
	// 	// strTime := "2024-03-19 21:26:03"
	// 	fmt.Println(Time.NowFormatTime(carbon.DateTimeLayout))
	// 	// fmt.Println(Time.NowFormatTime("ssss"))

	// 	t := Time.StrToCarbon("2023-01-11 21:26:03")
	// 	fmt.Println(Time.CarbonToDateTimeStr(t))

	// }
	// {
	// 	t := Time.StrToCarbon("2023-01-11 21:26:03")
	// 	fmt.Println(Time.CarbonToDate(t))
	// 	fmt.Println(Time.CarbonPtrToDate(&t))
	// }
	// {
	// 	t := Time.StrToCarbon("2027-01-11 21:26:03")
	// 	fmt.Println(Time.CarbonToDate(t))
	// 	fmt.Println(Time.AddDayToDate(&t, 100))
	// 	// fmt.Println(Time.CarbonPtrToDate(&t))
	// }
	{
		strTime := "2024-03-19 21:26:03"
		fmt.Println(Time.PStrToPCarbonDate(&strTime))
	}

	{
		// strTime := "2024-03-19 21:26:03"
		fmt.Println(Time.TodyRange())
		fmt.Println(Time.CurrDayStartEnd())
		// fmt.Println(Time.CurrNight())
		fmt.Println(Time.YdayRange())
		fmt.Println(Time.PreMonthRange())
	}

}
