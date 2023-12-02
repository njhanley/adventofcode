#!/usr/bin/awk -f

BEGIN { FS = "; " }

function max(a, b) { return a > b ? a : b }

{
	match($0, /^Game [0-9]+: /)
	$0 = substr($0, RSTART + RLENGTH)

	for (i = 1; i <= NF; i++) {
		split($i, subset, /, /)
		for (j in subset) {
			split(subset[j], pair, " ")
			cubes[pair[2]] = pair[1]
		}
		minimum["red"] = max(minimum["red"], cubes["red"])
		minimum["green"] = max(minimum["green"], cubes["green"])
		minimum["blue"] = max(minimum["blue"], cubes["blue"])
		delete cubes
	}
	sum += minimum["red"] * minimum["green"] * minimum["blue"]
	delete minimum
}

END { print sum }
