# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        rootNode = ListNode(0)
        rNode = rootNode
        carry = 0
        while l1 or l2 or carry:
            v = carry
            if l1:
                v += l1.val
            if l2:
                v += l2.val

            if v >= 10:
                v = v - 10
                carry = 1
            else:
                carry = 0
            
            rNode.next = ListNode(v)
            rNode = rNode.next
            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next

        return rootNode.next

