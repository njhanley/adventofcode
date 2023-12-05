#!/bin/sh

check() {
	printf '%s:%*s' "$1" "$((12 - ${#1}))"
	$1 -f "$2/part$3.awk" "$2/input.txt"
}

for day in *; do
	[ -d "$day" ] || continue
	for part in 1 2; do
		echo "=== Day $day Part $part ==="
		for awk in nawk gawk mawk 'busybox awk'; do
			check "$awk" "$day" "$part"
		done
		echo
	done
done
