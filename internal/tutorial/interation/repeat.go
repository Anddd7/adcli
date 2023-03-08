package iteration

const REPEAT_TIMES = 5

func Repeat(s string) string {
	var repeated string
	for i := 0; i < REPEAT_TIMES; i++ {
		repeated += s
	}
	return repeated
}
