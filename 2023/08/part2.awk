#!/usr/bin/awk -f

BEGIN { FS = "[^A-Z0-9]+" }

NR == 1 { split($0, instructions, ""); next }

{ L[$1] = $2; R[$1] = $3 }

$1 ~ /A$/ { nodes[NR] = $1 }

function gcd(a, b) { return b == 0 ? a : gcd(b, a % b) }
function lcm(a, b) { return a * (b / gcd(a, b)) }

END {
	while (length(nodes) > 0) {
		for (i in nodes) {
			if (instructions[step % length(instructions) + 1] == "L")
				nodes[i] = L[nodes[i]]
			else    nodes[i] = R[nodes[i]]

			if (nodes[i] ~ /Z$/) {
				cycles[i] = step + 1
				delete nodes[i]
			}
		}
		step++
	}
	steps = 1
	for (i in cycles) steps = lcm(steps, cycles[i])
	print steps
}
