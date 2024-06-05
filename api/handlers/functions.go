package handlers

import (
	"fmt"
	"strings"
	"time"
)

func CheckGmail(gmail string) bool {
	err := false
	for _, v := range gmail {
		if v == '@' {
			if strings.HasSuffix(gmail, ".com") {
				err = true
			}
			break
		}
	}

	return err

}
func CheckDate(dateStr string) (time.Time, bool) {
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please enter a date in YYYY-MM-DD format.")
		return date,false
	}

	if dateStr != date.Format(layout) {
		fmt.Println("Invalid date. Please enter a valid date.")
		return date,false
	}

	endDate, _ := time.Parse(layout, time.Now().Format("2006-01-02"))

	if date.After(endDate) {
		fmt.Println("Date is out of the acceptable range.")
		return date,false
	}

	return date,true
}
