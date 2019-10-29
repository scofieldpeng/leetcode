package dayOfTheWeek

func dayOfTheWeek(day int, month int, year int) string {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	var w int = 0

	if month == 1 || month == 2 {
		month = 12 + month
		year = year - 1
	}

	if (year < 1752) || (year == 1752 && month < 9) || (year == 1752 && month == 9 && day < 3) {
		w = (day + 2*month + 3*(month+1)/5 + year + year/4 + 5) % 7
	} else {
		w = (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400) % 7
	}

	return week[w]

}
