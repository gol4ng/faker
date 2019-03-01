package faker

import "time"

func TimeBetween(t1 time.Time, t2 time.Time) time.Time {
	t := IntBetween(int(t1.UnixNano()), int(t2.UnixNano()))
	return time.Unix(int64(t/1e9), int64(t%1e9))
}

func TimePast() time.Time {
	t := Intn64(time.Now().UnixNano())
	return time.Unix(t/1e9, t%1e9)
}

func TimeFuture() time.Time {
	t := Intn64(int64(time.Hour*24*365*10)) + time.Now().UnixNano()
	return time.Unix(t/1e9, t%1e9)
}
