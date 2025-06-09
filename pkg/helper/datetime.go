package helper

import "time"

func ListWeekdays(startDate, endDate string) []string {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)
	if start.After(end) {
		start, end = end, start // Ensure start is before or equal to end
	}

	listWeekdays := []string{}
	for d := start; !d.After(end); d = d.Add(24 * time.Hour) {
		weekday := d.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			listWeekdays = append(listWeekdays, d.Format("2006-01-02"))
		}
	}
	return listWeekdays
}

func DifferenceDate(a, b []string) (diff []string) {
	m := make(map[string]bool)
	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
