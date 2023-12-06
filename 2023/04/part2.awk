#!/usr/bin/awk -f

BEGIN { FS = "|" }

function splitset(s, a, fs, __a) {
	delete a
	split(s, __a, fs)
	for (i in __a) a[__a[i]] = 1
}

{
	sub(/^Card [0-9]+:[ ]*/, "")

	splitset($1, winners, " ")
	splitset($2, have, " ")

	matches = 0
	for (n in have) if (n in winners) matches++

	copies[NR]++
	for (i = 1; i <= matches; i++) copies[NR + i] += copies[NR]
}

END {
	for (i in copies) total += copies[i]
	print total
}
