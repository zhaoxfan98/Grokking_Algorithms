package golang

func binary_search(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

var my_list []int = []int{1, 3, 5, 6, 7}

// func main() {
// 	fmt.Println(binary_search(my_list, 3))
// }
