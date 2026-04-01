package str

func DeleteElement(str string, elementIndex int) (newString string) {

	for i := 0; i < len(str); i++ {
		if i != elementIndex {
			newString += string(str[i])
		}
	}

	return newString

}

func ChangeElement(str string, position int, element string) (newString string) {

	for i := 0; i < len(str); i++ {
		if i != position {
			newString += string(str[i])
		} else {
			newString += element
		}
	}

	return newString

}

func ReverseString(str string) (reverseString string) {

	for i := len(str) - 1; i >= 0; i-- {
		reverseString += string(str[i])
	}

	return reverseString

}

func RepeatElementCounter(str string, element string) (counter int) {

	for i := 0; i < len(str); i++ {
		if string(str[i]) == element {
			counter++
		}
	}

	return counter

}
