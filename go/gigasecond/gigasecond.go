package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

var gigasecond, _ = time.ParseDuration("1000000000s")

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
