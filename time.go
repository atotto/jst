package jst

import "time"

var (
	Location = time.FixedZone("JST", 9*60*60)
)

func Now() time.Time {
	return time.Now().In(Location)
}

func JST(t time.Time) time.Time {
	return t.In(Location)
}
