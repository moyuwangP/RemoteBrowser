package utils

func PanicIfNeeded(err error) {
	if err != nil {
		panic(err)
	}
}
