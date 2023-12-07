#!/usr/bin/awk -f

function min(a, b) { return a < b ? a : b }
function max(a, b) { return a > b ? a : b }

function push(a, x) {
	"" in a # coerce into array
	a[length(a) + 1] = x
}

function pop(a, __x, __i) {
	__x = a[1]
	for (__i = 1; __i < length(a); __i++) a[__i] = a[__i + 1]
	delete a[__i]
	return __x
}

BEGIN { RS = "" }

{ sub(/^.*:[ \n]/, "") }

NR == 1 {
	for (i = 1; i < NF; i += 2) push(ranges, $i " " $i + $(i + 1))
	next
}

{
	split($0, maps, "\n")

	for (i in ranges) push(queue, ranges[i])
	delete ranges

	while (range = pop(queue)) {
		split(range, r)
		if (r[1] == r[2]) continue

		for (i in maps) {
			split(maps[i], m)
			m1 = m[2]; m2 = m[2] + m[3]; d = m[1] - m[2]

			if (disjoint = r[2] <= m1 || m2 <= r[1]) continue

			o1 = max(r[1], m1)
			o2 = min(r[2], m2)
			push(queue, r[1] " " o1)
			push(ranges, o1 + d " " o2 + d)
			push(queue, o2 " " r[2])
			break
		}
		if (disjoint)
			push(ranges, range)
	}
}

END {
	range = pop(ranges)
	split(range, r)
	lowest = r[1]
	while (range = pop(ranges)) {
		split(range, r)
		if (r[1] < lowest) lowest = r[1]
	}
	print lowest
}
