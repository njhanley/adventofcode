#!/usr/bin/awk -f

function all(a, x, __i) {
	for (__i in a) if (a[__i] != x) return 0
	return 1
}

{
	split($0, a)
	do {
		for (i = 1; i < length(a); i++) a[i] = a[i + 1] - a[i]
		total += a[i]
		delete a[i]
	} while (!all(a, 0))
}

END { print total }
