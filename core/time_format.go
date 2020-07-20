package core

import (
	"time"
)

type MTime time.Time

const TimeFormat = "2006-01-02 15:04:05"

func (t MTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(TimeFormat) + `"`), nil
}

func (t *MTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = MTime(now)
	return
}

func (t MTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

//func (t *MTime) String() string {
//	return t.Format(TimeFormat)
//}
