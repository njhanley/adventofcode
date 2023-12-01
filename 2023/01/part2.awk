#!/usr/bin/awk -f

BEGIN { split("one two three four five six seven eight nine", words) }

function parse(s) {
	if (s ~ /^[1-9]/) return substr(s, 1, 1)
	for (i in words) if (index(s, words[i]) == 1) return i
	return ""
}

{
	for (digits = ""; $0; $0 = substr($0, 2))
		digits = digits parse($0)
	count += substr(digits, 1, 1) substr(digits, length(digits), 1)
}

END { print count }
