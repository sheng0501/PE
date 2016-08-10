# Use the fact, assuming p/q is the first quotient, then
# 1 + 1 / (1 + p/q) =  ( p + 2q ) / ( p + q )
import math

x = [1]
y = [1]

for i in range(1,1000):
	x += [x[i-1] + 2 * y[i-1]]
	y += [x[i-1] + y[i-1]]
	

count = 0

for i in range(0,1000):
	if len(str(x[i])) > len(str(y[i])):
		count+=1

print count