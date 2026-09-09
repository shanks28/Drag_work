class Solution:
    def calPoints(self, operations: List[str]) -> int:
        stack=[]
        def is_integer(n):
            try:
                int(n)
                return True
            except:
                return False
        for i,v in enumerate(operations):
            if is_integer(v):
                print(v)
                stack.append(v)
            elif v == "+":
                n1=int(stack[-1])
                n2=int(stack[-2])
                n3=n1+n2
                stack.append(str(n3))
            elif v == "D":
                print(stack)
                n1=int(stack[-1])
                n2=2*n1
                stack.append(str(n2))
            elif v =="C":
                stack.pop()
        func=lambda x : int(x)
        result=list(map(func,stack))
        print(result)
        return sum(result)