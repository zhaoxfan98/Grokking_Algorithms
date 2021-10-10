package golang

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := []int{}
	greater := []int{}
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less, greater = quickSort(less), quickSort(greater)
	myarr := append(append(less, pivot), greater...)
	return myarr
}

// func main() {
// 	my_list := []int{1, 3, 9, 8, 7, 1}
// 	fmt.Println(quickSort(my_list))
// }
