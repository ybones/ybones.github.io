class Solution:
    
    def isValid(self, s: str) -> bool:
        itemTurn = {
            ')':'(', '}':'{', ']':'[',
        }
        stackList = []
        for _s in s:
            if _s == " ":
                continue
            _v = itemTurn.get(_s, None)
            if _v:
                if len(stackList)> 0 and _v == stackList.pop(-1):
                    continue
                else:
                    return False
            
            stackList.append(_s)
        
        return len(stackList) == 0
