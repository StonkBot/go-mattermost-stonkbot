package main

// https://freshman.tech/snippets/go/check-if-slice-contains-element/
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
