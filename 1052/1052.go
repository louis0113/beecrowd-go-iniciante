package main

import (
	"fmt"
)

func main() {

	var month int
	var month_name string

	fmt.Scanln(&month)
	switch month {
	case 1:
		month_name = "January"
	case 2:
		month_name = "February"
	case 3:
		month_name = "March"
	case 4:
		month_name = "April"
	case 5:
		month_name = "May"
	case 6:
		month_name = "June"
	case 7:
		month_name = "July"
	case 8:
		month_name = "August"
	case 9:
		month_name = "September"
	case 10:
		month_name = "October"
	case 11:
		month_name = "November"
	case 12:
		month_name = "December"
	}

	fmt.Println(month_name)
}
