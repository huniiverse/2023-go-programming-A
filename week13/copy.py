import copy

#a = [3, -99, 17]
a = [3, [5, 1], 17]
b = a
c = a[:]  #shallow copy
d = list(a) #shallow copy
e = a.copy()  #shallow copy
f = copy.deepcopy(a)  #deepcopy
print(a, b, c, d, e, f)
#a[1] = 100
a[1][0] = 100
print(a, b, c, d, e, f)