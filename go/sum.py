#!/bin/bash

class Node:
    def __init__(self, val = 0):
        self.val = val
        self.next = None



def sumLists(list1, list2):
    
    stack1 = []
    stack2 = []
    head = None
    n1 = list1
    n2 = list2

    while n1 is not None:
        stack1.append(n1.val)
        n1 = n1.next

    while n2 is not None:
        stack2.append(n2.val)
        n2 = n2.next

    carry = 0
    while stack1 != [] or stack2 != []:
        val1 = stack1.pop()
        val2 = stack2.pop()

        total = val1+val2+carry
        node = Node()
        if total > 9:
            carry = 1
            node.val = total % 10
        else:
            carry = 0
            node.val = total

        node.next = head
        head = node

    if carry == 1:
        node = Node(1)
        node.next = head
        head = node
    return head

def helper(carry, l1, l2, head):

    if l1 is not None and l2 is not None:
        node = Node()
        total = carry+l1.val+l2.val
        if total > 9:
            node.val = total % 10
            carry = 1
        else:
            node.val = total
            carry = 0



def recur(l1, l2):
    head = None
    carry = 0
    head = helper(carry, l1, l2, head)
    return head


n1 = Node(6)
n2 = Node(1)
n3 = Node(7)

m1 = Node(2)
m2 = Node(9)
m3 = Node(5)

n1.next = n2
n2.next = n3

m1.next = m2
m2.next = m3

summed = sumLists(n1,m1)
while summed is not None:
    print(summed.val)
    summed = summed.next




