# encoding=utf8
# python 3.10.12 版本

# 冒泡排序
def bubble_sort(collection:list[int])->list[int]:
    for i in reversed(range(len(collection))):
        swapped = False
        for j in range(i):
            if collection[j]>collection[j+1]:
                swapped = True
                collection[j], collection[j+1] = collection[j+1], collection[j]
        if not swapped:
            break
    return collection

# 插入排序
def insert_sort(collection:list[int])->list[int]:
    for i in range(1, len(collection)):
        cur_val = collection[i]
        j = i-1
        while j >=0 and cur_val<collection[j]:
            collection[j+1] = collection[j]
            j = j - 1
        collection[j+1] = cur_val
        
    return collection
                
# 选择排序
def select_sort(collection:list[int])->list[int]:
    for rest_begin_idx in range(len(collection)-1):
        min_idx = rest_begin_idx
        for cmp_idx in range(rest_begin_idx+1, len(collection)):
            if collection[cmp_idx] < collection[min_idx]:
                min_idx = cmp_idx
        if min_idx != rest_begin_idx:
            collection[min_idx],collection[rest_begin_idx] = collection[rest_begin_idx],collection[min_idx]
    return collection

# 希尔排序
def shell_sort(collection:list[int])->list[int]:
    step = len(collection)//2 # TODO
    while step>=1:
        for i in range(step, len(collection)):
            cur_val = collection[i]
            j = i-step
            while j >=0 and cur_val< collection[j]:
                collection[j+step] = collection[j]
                j = j-step
            collection[j+step] = cur_val
        if step <= 1:
            break
        step = step // 2
    return collection

# 合并排序
def merge_sort(collection:list[int])->list[int]:
    def merge(left:list, right:list)->list:
        result = []
        while left and right:
            result.append(left.pop(0) if left[0]<=right[0] else right.pop(0))
        result.extend(left)
        result.extend(right)
        return result
    if len(collection) <=1:
        return collection
    mid = len(collection) // 2
    return merge(merge_sort(collection[:mid]), merge_sort(collection[mid:]))
    

if __name__ == "__main__":
    cases = [
        ([],[]),
        ([1],[1]),
        ([1,2],[1,2]),
        ([1,3,2], [1,2,3]),
        ([1,2,2,3,4,5,6], [1,2,2,3,4,5,6])
    ]
    for input,expect_res in cases:
        #shell_sort(input)
        #bubble_sort(input)
        #insert_sort(input)
        #select_sort(input)
        input = merge_sort(input)
        print(input, expect_res)
        assert input == expect_res