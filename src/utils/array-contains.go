package utils

func Contains(arr []string, elem string) (res int) {

	for index, value := range arr {
		if value == elem {
			return index
		}
	}

	return -1
}
