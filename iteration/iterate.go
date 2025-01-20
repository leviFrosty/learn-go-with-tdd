package iteration

func Repeat(thing string, times int) string {
	str := ""
	for i := 0; i < times; i++ {
		str += thing
	}
	return str
}
