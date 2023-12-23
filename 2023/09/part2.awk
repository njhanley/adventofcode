#!/usr/bin/awk -f

function all(a, x, __i) {
	for (__i in a) if (a[__i] != x) return 0
	return 1
}

{
	split($0, a)
	delete d
	do {
		d[length(d) + 1] = a[1]
		for (i = 1; i < length(a); i++) a[i] = a[i + 1] - a[i]
		delete a[i]
	} while (!all(a, 0))
	for (i = length(d) - 1; i > 0; i--) d[i] -= d[i + 1]
	total += d[1]
}

END { print total }
