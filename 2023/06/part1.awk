#!/usr/bin/awk -f

NR == 1 { for (i = 2; i <= NF; i++) times[i - 1] = $i }
NR == 2 { for (i = 2; i <= NF; i++) dists[i - 1] = $i }

END {
	total = 1
	for (i = 1; i <= length(times); i++) {
		ways = 0
		for (hold = 1; hold < times[i]; hold++)
			if (hold * (times[i] - hold) > dists[i]) ways++
		total *= ways
	}
	print total
}
