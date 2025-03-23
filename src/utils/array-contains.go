package utils

import "log"

func Contains(arr []string, elem string) (res int) {

	log.Println("These are your data", arr, elem)
	for index, value := range arr {
		if value == elem {
			return index
		}
	}

	return -1
}
