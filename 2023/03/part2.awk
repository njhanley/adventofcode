#!/usr/bin/awk -f

BEGIN { FS = ""; delete values }

{
	for (i = 1; i <= NF; i++)
		if ($i == "*")
			symbols[NR, i] = 1
	for (i = 0; match($0, /[0-9]+/); i += RSTART + RLENGTH - 1) {
		values[k = length(values) + 1] = substr($0, RSTART, RLENGTH)
		for (j = 0; j < RLENGTH; j++)
			numbers[NR, i + RSTART + j] = k
		$0 = substr($0, RSTART + RLENGTH)
	}
}

END {
	for (coord in symbols) {
		split(coord, p, SUBSEP)
		for (y = p[1] - 1; y <= p[1] + 1; y++) {
			for (x = p[2] - 1; x <= p[2] + 1; x++) {
				if ((y, x) in numbers && !(numbers[y, x] in counted)) {
					symbols[coord] *= values[k = numbers[y, x]]
					connections[coord]++
					counted[k]
				}
			}
		}
		delete counted
	}
	for (coord in connections)
		if (connections[coord] == 2)
			sum += symbols[coord]
}

END { print sum }
