package mbiz

import (
	"errors"
	"fmt"
	"time"
)

func case1(arr []int) (res int) {

	for _, v := range arr {
		res = res + v
	}
	return res
}

func case2(date string) (res string) {
	date2, err := time.Parse("Mon 2 Jan 2006", date)
	if err != nil {
		return errors.New("invalid date format").Error()
	}
	res = date2.Format("2006-01-02")
	return
}

func case3(num int) (res int) {
	res = num
	for i := num; i-1 > 0; i-- {
		res = res * (i - 1)

	}
	return
}

func case4(h int, m int) (res string) {
	nums := [30]string{"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen", "seventeen",
		"eighteen", "nineteen", "twenty", "twenty-one",
		"twenty-two", "twenty-three", "twenty-four",
		"twenty-five", "twenty-six", "twenty-seven",
		"twenty-eight", "twenty-nine"}

	if m == 60 {
		return nums[h+1] + " o'clock"
	} else if m == 0 {
		return nums[h] + " o'clock"
	} else if m == 1 {
		return "one minute past " + nums[h]
	} else if m == 59 {
		return "one minute to " + nums[(h%12)+1]
	} else if m == 15 {
		return "quarter past " + nums[h]
	} else if m == 30 {
		return "half past " + nums[h]
	} else if m == 45 {
		return "quarter to " +
			nums[(h%12)+1]
	} else if m <= 30 {
		return nums[m] + " minutes past " + nums[h]
	} else if m > 30 && m <= 59 {
		return nums[60-m] + " minutes to " + nums[(h%12)+1]
	} else {
		return "invalid time format"
	}
}

func case5(arr []int) (res int) {
	data := make(map[int]int)
	for _, v := range arr {
		data[v]++
	}
	maxKey := 0
	maxValue := 0
	for key, value := range data {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}
	fmt.Println("ini max key", maxKey)

	return len(arr) - maxValue
}
