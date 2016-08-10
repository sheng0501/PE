def digit_sum(n):
	s = 0
	while n:
		s += n % 10
		n /= 10
	return s


result = 0

for x in range(1,100):
	for y in range (1,100):
		z = x**y
		s = digit_sum(z)
		if s > result:
			result = s

print result


