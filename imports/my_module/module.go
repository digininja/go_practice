package module

func ExposedFunction() string {
	return "This function name has a capital letter"
}

func notExposedFunction() string {
	return "This is not exposed because of lower case function name so can't be imported"
}
