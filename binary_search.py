# -*- coding:utf-8 -*-

def binary_search(list, item):
    low = 0
    high = len(list)-1

    while low <= high:
        mid = (low + high) // 2     # 注意使用整除符号
        guess = list[mid]
        if guess == item:
            return mid
        if guess > item:
            high = mid-1
        else:
            low = mid+1
    return None

my_list = [1,3,5,6,7]
print(binary_search(my_list, 3)) 
