package go_contract_modules

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
