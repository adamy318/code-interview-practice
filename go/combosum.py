#!/bin/bash


def comboSum(arr, target):
    def backtrack(arr, curr, t):
        if t <= 0:
            if t == 0:
                if sorted(curr) not in out:
                    out.append(curr[:])
            return

        for i in range(len(arr)):
            print(curr, i)
            if arr[i] > t:
                return
            if arr[i] <= t:
                curr.append(arr[i])

            backtrack(arr, curr, t-arr[i])
            curr.pop()
            print(curr, "popped")


    out = []

    real = sorted(arr)
    true = []
    for i in range(len(real)-1):
        if real[i] == real[i+1]:
            continue
        else:
            true.append(real[i])
    true.append(real[-1])

    backtrack(true, [], target)

    return out


a = [8, 10, 6, 11, 1, 16, 8]
b = 28 
c = [2,3,4,6,7]
d = 7
print(comboSum(c,d))
