# -*- coding:utf-8 -*-

def findSmallest(arr):
    smallest = arr[0]
    smallest_index = 0
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i
    return smallest_index

def selectionSort(arr):
    newArr = []
    for _ in range(len(arr)):
        smallest = findSmallest(arr)    # 找到数组中最小的元素
        newArr.append(arr.pop(smallest))
    
    return newArr