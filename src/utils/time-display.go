package utils

import (
	"fmt"
	"time"
)


func FormatTimestamp(timestamp time.Time) (res string) {
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return timestamp.In(loc).Format("January 2, 2006 (03:04 PM)")
}