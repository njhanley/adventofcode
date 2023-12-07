#!/usr/bin/awk -f

BEGIN { RS = "" }

{ sub(/^.*:[ \n]/, "") }

NR == 1 { split($0, seeds); next }

{
	split($0, maps, "\n")
	for (i in maps) {
		split(maps[i], m)
		dst = m[1]; src = m[2]; len = m[3]
		for (k in seeds) if (src <= seeds[k] && seeds[k] < src + len)
			locations[k] = seeds[k] - src + dst
		if (!(k in locations))
			locations[k] = seeds[k]
	}
	for (k in locations) seeds[k] = locations[k]
	delete locations
}

END {
	lowest = seeds[1]
	for (i in seeds) if (seeds[i] < lowest) lowest = seeds[i]
	print lowest
}
