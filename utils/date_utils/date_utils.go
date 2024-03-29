package date_utils

import "time"

const (
	formatDate = "2006-01-02T15:04:05Z"
	apiDBLayout = "2006-01-02 15:04:05"
)


func GetNow() time.Time {
	return time.Now().UTC()
}
func GetNowString() string {
	return GetNow().Format(formatDate)
}

func GetNowDB() string {
	return GetNow().Format(apiDBLayout)
}