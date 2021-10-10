package golang

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i := range arr {
		if smallest > arr[i] {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectSort(arr []int) []int {
	newArr := []int{}
	n := len(arr)
	for i := 0; i < n; i++ {
		smallest_index := findSmallest(arr)
		newArr = append(newArr, arr[smallest_index])
		arr = append(arr[:smallest_index], arr[smallest_index+1:]...)
	}
	return newArr
}

//删除指定的索引元素
func DelIndex() {
	var DelIndex []int
	DelIndex = make([]int, 5)
	DelIndex[0] = 0
	DelIndex[1] = 1
	DelIndex[2] = 2
	DelIndex[3] = 3
	DelIndex[4] = 4
	fmt.Println("删除指定索引(下标)前>>:")
	fmt.Printf("len:%v\tDelIndex:%v\n", len(DelIndex), DelIndex)
	//删除元素 3 索引(下标) 3
	index := 3
	// 这里通过 append 方法 分成两个然后合并
	// append(切片名,追加的元素) 切片名这里我们进行切割一个新的切片DelIndex[:index] 追加的元素将索引后面的元素追加
	// DelIndex[index+1:]...) 为什么追加会有...三个点, 因为是一个切片 所以需要展开
	DelIndex = append(DelIndex[:index], DelIndex[index+1:]...)
	fmt.Printf("len:%v\tDelIndex:%v\n", len(DelIndex), DelIndex)
	fmt.Println("DelIndex:", DelIndex)
}

// func main() {
// 	my_list := []int{1, 3, 9, 8, 7}
// 	fmt.Println(selectSort(my_list))

// 	// DelIndex()
// }
