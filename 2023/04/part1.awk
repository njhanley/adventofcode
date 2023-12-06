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

	points = 0
	for (n in have) if (n in winners) points = points ? 2 * points : 1
	total += points
}

END { print total }
