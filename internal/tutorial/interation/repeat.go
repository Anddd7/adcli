package iteration

const REPEAT_TIMES = 5

func Repeat(s string) string {
	var repeated string
	for i := 0; i < REPEAT_TIMES; i++ {
		repeated += s
	}
	return repeated
}

func RepeatFast(s string) string {
	repeated := make([]byte, REPEAT_TIMES*len(s))
	for i := 0; i < REPEAT_TIMES; i++ {
		for j := 0; j < len(s); j++ {
			repeated[j+i*len(s)] = s[j]
		}
	}
	return string(repeated)
}
