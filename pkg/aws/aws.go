package aws

// Can be imported in other packages
func HelloWord() string {
	return "Hello World"
}

// Cannot be imported/used in other packages
func HelloWordCannotImport() string {
	return "Hello World"
}
