package go_web

func Error(err error) {
	if err != nil {
		panic(err)
	}
}