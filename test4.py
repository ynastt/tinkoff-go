def is_boring(arr):
    s = set(arr)
    if len(s) == 2:
        if arr.count(min(s)) == 1:
            return True
        if arr.count(max(s)) == 1:
            return True
    if arr.count(1) == len(arr):
        return True
    return False
 
 
n = int(input())
*a, = map(int, input().split())
ans = 1
d = {}
for i in range(n):
    key = a[i]
    d[key] = d.get(key, 0) + 1
    if is_boring(list(d.values())):
        ans = i + 1
print(ans)