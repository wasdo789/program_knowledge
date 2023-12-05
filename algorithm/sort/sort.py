# encoding=utf8
# python 3.10.12 版本

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
# 插入排序


if __name__ == "__main__":
    cases = [
        ([],[]),
        ([1],[1]),
        ([1,2],[1,2]),
        ([1,3,2], [1,2,3])
    ]
    for input,expect_res in cases:
        select_sort(input)
        print(input, expect_res)
        assert input == expect_res