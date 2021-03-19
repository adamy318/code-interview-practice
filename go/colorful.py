#!/bin/bash

def colorful(A):

        def intToList(num):
            ret = []
            while num > 0:
                ret.insert(0, num % 10)
                num = num // 10
            return ret

        def listToInt(arr):
            n = len(arr)-1

            num = 0
            for i in arr:
                num += (i * (10 ** n))
                n -= 1
            return num

        def product(arr):
            total = 1
            for i in arr:
                total *= i
            return total

        sums = {}
        nums = intToList(A)
        n = len(nums)

        for k in range(1,n+1):
            print(k)
            for i in range(n-k+1):
                print("i: ",i)
                total = product(nums[i:i+k])
                if total in sums.keys():
                    print("hey")
                    return 0
                sums[total] = listToInt(nums[i:i+k])
        return 1


print("colorful: ",colorful(99))
