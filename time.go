package jst

import "time"

var (
	JST = time.FixedZone("JST", 9*60*60)
)

func Now() time.Time {
	return time.Now().In(JST)
}
