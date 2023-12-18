#!/usr/bin/awk -f

BEGIN { FS = "[^A-Z]+" }

NR == 1 { split($0, instructions, ""); next }

{ L[$1] = $2; R[$1] = $3 }

END {
	node = "AAA"
	while (node != "ZZZ") {
		if (instructions[step++ % length(instructions) + 1] == "L")
			node = L[node]
		else    node = R[node]
	}
	print step
}
