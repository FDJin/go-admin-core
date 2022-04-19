package pkg

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetDateAfterMonths(num int) string {
	currentTime := time.Now().AddDate(0, num, 0)
	return currentTime.String()
}

func DataTimeToDate(e string) string{
	if find := strings.Contains(e, "T00"); find{
		e = e[:10]
	}

	return e
}
