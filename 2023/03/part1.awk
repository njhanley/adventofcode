#!/usr/bin/awk -f

BEGIN { FS = ""; delete values }

{
	for (i = 1; i <= NF; i++)
		if ($i !~ /[0-9.]/)
			symbols[NR, i] = $i
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
		for (y = p[1] - 1; y <= p[1] + 1; y++)
			for (x = p[2] - 1; x <= p[2] + 1; x++)
				if ((y, x) in numbers)
					parts[y, x] = numbers[y, x]
	}
	for (coord in parts) {
		if ((k = parts[coord]) in values) {
			sum += values[k]
			delete values[k]
		}
	}
}

END { print sum }
