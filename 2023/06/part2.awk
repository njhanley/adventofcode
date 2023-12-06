#!/usr/bin/awk -f

{ sub(/^.*:/, ""); gsub(/ /, "") }
NR == 1 { time = 0 + $0 }
NR == 2 { dist = 0 + $0 }

END {
	ways = 0
	for (hold = 1; hold < time; hold++)
		if (hold * (time - hold) > dist) ways++
	print ways
}
