package main

import (
	"fmt"
	"time"
)

func main() {

	// THERE ARE 3 WAYS TO CREATE TIME IN GOLANG
	
	// BY AUTOMATIC
	time1 := time.Now()
	fmt.Println(time1)

	// BY MANUAL
	time2 := time.Date(2022, time.June, 29, 10, 30, 45, 10, time.Local)
	fmt.Println(time2)

	// BY PARSING
	format := "2006-01-02"	// 2006-01-02 is the mandatory date format, time3 will be inappropriate if it's not 2006-01-02
	time3, _ := time.Parse(format, "2022-07-29")
	fmt.Println(time3)
}

// OUTPUT
// 2022-07-25 01:14:33.545961808 +0700 WIB m=+0.000050077
// 2022-06-29 10:30:45.00000001 +0700 WIB
// 2022-07-29 00:00:00 +0000 UTC