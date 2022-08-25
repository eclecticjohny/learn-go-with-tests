package iteration

func Repeat(c string, r int) (repeated string) {
	for i := 0; i < r; i++ {
		repeated += c
	}
	return
}
