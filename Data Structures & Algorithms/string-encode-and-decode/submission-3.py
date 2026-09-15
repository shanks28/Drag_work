class Solution:

    def encode(self, strs: List[str]) -> str:
        res=""
        for w in strs:
            res+=str(len(w))
            res+="#"
            res+=w
        return res

    def decode(self, s: str) -> List[str]:
        i=0
        res=[]
        while i < len(s):
            j=i
            while j < len(s) and s[j]!="#":
                j+=1
            length=int(s[i:j])
            start=j+1
            end=start+length
            res.append(s[start:end])
            i=end
        return res