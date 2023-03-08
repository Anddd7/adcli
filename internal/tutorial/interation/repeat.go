package iteration

func Repeat(s string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += s
	}
	return repeated
}

func RepeatFast(s string, times int) string {
	repeated := make([]byte, times*len(s))
	for i := 0; i < times; i++ {
		for j := 0; j < len(s); j++ {
			repeated[j+i*len(s)] = s[j]
		}
	}
	return string(repeated)
}
