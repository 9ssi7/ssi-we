package helper

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErrWithMsg(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}