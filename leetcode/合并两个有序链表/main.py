# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        rootNode = ListNode()
        tempNode = rootNode
        while l1 or l2:
            v1 = l1.val if l1 else 0
            v2 = l2.val if l2 else 0
            
            tempValue = 0
            if (l1 and not l2) or (l1 and l2 and v1 < v2):
                tempValue = v1
                l1 = l1.next
            else:
                tempValue = v2
                l2 = l2.next
            tempNode.next = ListNode(tempValue)
            tempNode = tempNode.next

        return rootNode.next
            