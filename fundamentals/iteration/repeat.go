package iteration

func Repeat(a string, n int) string {
	var r string
	for i := 0; i < n; i++ {
		r += a
	}

	return r
}
