package arr

func NumsArrSlice(arr []int, start int, end int) (newArr []int) {

	for i := start; i < end; i++ {
		newArr = append(newArr, arr[i])
	}

	return newArr

}

func StrArrSlice(arr []string, start int, end int) (newArr []string) {

	for i := start; i < end; i++ {
		newArr = append(newArr, arr[i])
	}

	return newArr

}

func NumsArrInclude(arr []int, element int) (isInclude bool) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			return true
		}
	}

	return false

}

func StrArrInclude(arr []string, element string) (isInclude bool) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			return true
		}
	}

	return false

}
