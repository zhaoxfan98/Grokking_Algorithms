# -*- coding:utf-8 -*-

def quickSort(array):
    if len(array) < 2:
        return array
    else:
        pivot = array[0]    # 选择基准值
        # 由所有小于等于基准值的元素组成的子数组
        less = [i for i in array[1:] if i <= pivot]
        # 由所有大于基准值的元素组成的子数组
        greater = [i for i in array[1:] if i > pivot]
        return quickSort(less) + [pivot] + quickSort(greater)