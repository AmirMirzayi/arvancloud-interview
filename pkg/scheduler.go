package pkg

import "time"

func GetNextMonth() time.Time {
	nextMonth := time.Now().AddDate(0, 1, 0)
	nextMonth = time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, nextMonth.Location())
	return nextMonth
}

func GetNextMinute() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()+1, 0, 0, time.Now().Location())
}

func CronJob(timeOffset func() time.Time, job func()) {
	duration := timeOffset().Sub(time.Now())
	ticker := time.NewTicker(duration)
	for range ticker.C {
		job()
		duration = timeOffset().Sub(time.Now())
		ticker.Reset(duration)
	}
}
