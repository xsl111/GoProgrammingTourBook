package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	//location, _ := time.LoadLocation("America/New_York")
	return time.Now().In(location)
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(duration), nil
}

const (
	Nansecond   time.Duration = 1
	Microsecond               = 1000 * Nansecond
	Millsecond                = 1000 * Microsecond
	Second                    = 1000 * Millsecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)
