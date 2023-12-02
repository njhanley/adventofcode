#!/usr/bin/awk -f

BEGIN {
	split("red green blue", colors)
	FS = "; "
}

function max(a, b) { return a > b ? a : b }
function product(a, __x) { __x = 1; for (i in a) __x *= a[i]; return __x }

{
	match($0, /^Game [0-9]+: /)
	$0 = substr($0, RSTART + RLENGTH)

	for (i = 1; i <= NF; i++) {
		split($i, subset, ", ")
		for (j in subset) {
			split(subset[j], pair, " ")
			cubes[pair[2]] = pair[1]
		}
		for (j in colors) {
			c = colors[j]
			minimum[c] = max(minimum[c], cubes[c])
		}
		delete cubes
	}

	sum += product(minimum)
	delete minimum
}

END { print sum }
