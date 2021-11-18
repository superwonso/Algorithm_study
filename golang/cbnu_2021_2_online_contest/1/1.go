package main

import "fmt"

func main() {
	var year, month, days int
	var input string
	fmt.Scanln(&input)
	if input[7] == ' ' {
		year = int(input[0]-'0')*100 + int(input[1]-'0')*10 + int(input[2]-'0')
		month = int(input[3]-'0')*10 + int(input[4]-'0')
		days = int(input[5]-'0')*10 + int(input[6]-'0')
	} else if input[6] == ' ' {
		year = int(input[0]-'0')*10 + int(input[1]-'0')
		month = int(input[2]-'0')*10 + int(input[3]-'0')
		days = int(input[4]-'0')*10 + int(input[5]-'0')
	} else if input[5] == ' ' {
		year = int(input[0] - '0')
		month = int(input[1]-'0')*10 + int(input[2]-'0')
		days = int(input[3]-'0')*10 + int(input[4]-'0')
	} else if input[7] != ' ' {
		year = int(input[0]-'0')*1000 + int(input[1]-'0')*100 + int(input[2]-'0')*10 + int(input[3]-'0')
		month = int(input[4]-'0')*10 + int(input[5]-'0')
		days = int(input[6]-'0')*10 + int(input[7]-'0')
	}

	if year >= 1000 && year < 10000 {
		if validYearMonthDays(year, month, days) == true {
			if month < 10 && days < 10 {
				fmt.Printf("%d/0%d/0%d", year, month, days)
			} else if month < 10 && days >= 10 {
				fmt.Printf("%d/0%d/%d", year, month, days)
			} else if month >= 10 && days < 10 {
				fmt.Printf("%d/%d/0%d", year, month, days)
			} else {
				fmt.Printf("%d/%d/%d", year, month, days)
			}
		} else if validYearMonthDays(year, month, days) == false {
			fmt.Printf("-1")
		}
	} else if year < 1000 && year > 99 {
		if validYearMonthDays(year, month, days) == true {
			if month < 10 && days < 10 {
				fmt.Printf("0%d/0%d/0%d", year, month, days)
			} else if month < 10 && days >= 10 {
				fmt.Printf("0%d/0%d/%d", year, month, days)
			} else if month >= 10 && days < 10 {
				fmt.Printf("0%d/%d/0%d", year, month, days)
			} else {
				fmt.Printf("0%d/%d/%d", year, month, days)
			}
		} else if validYearMonthDays(year, month, days) == false {
			fmt.Printf("-1")
		}
	} else if year < 100 && year > 9 {
		if validYearMonthDays(year, month, days) == true {
			if month < 10 && days < 10 {
				fmt.Printf("00%d/0%d/0%d", year, month, days)
			} else if month < 10 && days >= 10 {
				fmt.Printf("00%d/0%d/%d", year, month, days)
			} else if month >= 10 && days < 10 {
				fmt.Printf("00%d/%d/0%d", year, month, days)
			} else {
				fmt.Printf("00%d/%d/%d", year, month, days)
			}
		} else if validYearMonthDays(year, month, days) == false {
			fmt.Printf("-1")
		}
	} else if year < 10 && year > 0 {
		if validYearMonthDays(year, month, days) == true {
			if month < 10 && days < 10 {
				fmt.Printf("000%d/0%d/0%d", year, month, days)
			} else if month < 10 && days >= 10 {
				fmt.Printf("000%d/0%d/%d", year, month, days)
			} else if month >= 10 && days < 10 {
				fmt.Printf("000%d/%d/0%d", year, month, days)
			} else {
				fmt.Printf("000%d/%d/%d", year, month, days)
			}
		} else if validYearMonthDays(year, month, days) == false {
			fmt.Printf("-1")
		}
	} else if year == 0 {
		if validYearMonthDays(year, month, days) == true {
			if month < 10 && days < 10 {
				fmt.Printf("0000/0%d/0%d", month, days)
			} else if month < 10 && days >= 10 {
				fmt.Printf("0000/0%d/%d", month, days)
			} else if month >= 10 && days < 10 {
				fmt.Printf("0000/%d/0%d", month, days)
			} else {
				fmt.Printf("0000/%d/%d", month, days)
			}
		} else if validYearMonthDays(year, month, days) == false {
			fmt.Printf("-1")
		}
	} else {
		fmt.Printf("-1")
	}
}

func validYearMonthDays(year int, month int, days int) bool {
	if year < 0 {
		return false
	} else {
		if month < 1 || month > 12 {
			return false
		} else if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
			if days < 1 || days > 31 {
				return false
			}
		} else if month == 2 {
			if days < 1 || days > 29 {
				return false
			}
		} else if month == 4 || month == 6 || month == 9 || month == 11 {
			if days < 1 || days > 30 {
				return false
			}
		}
	}
	return true
}
