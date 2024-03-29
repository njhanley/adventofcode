#!/usr/bin/awk -f

BEGIN { FS = "; " }

{
	sub(/^Game [0-9]+: /, "")

	valid = 1
	for (i = 1; i <= NF; i++) {
		split($i, subset, ", ")
		for (j in subset) {
			split(subset[j], pair, " ")
			cubes[pair[2]] = pair[1]
		}
		if (cubes["red"] > 12 ||
		    cubes["green"] > 13 ||
		    cubes["blue"] > 14)
			valid = 0
		delete cubes
	}
	if (valid)
		sum += NR
}

END { print sum }
