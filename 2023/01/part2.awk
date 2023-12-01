#!/usr/bin/awk -f

{
	for (digits = ""; $0; $0 = substr($0, 2)) {
		if ($0 ~ /^one/)   { digits = digits "1"; continue }
		if ($0 ~ /^two/)   { digits = digits "2"; continue }
		if ($0 ~ /^three/) { digits = digits "3"; continue }
		if ($0 ~ /^four/)  { digits = digits "4"; continue }
		if ($0 ~ /^five/)  { digits = digits "5"; continue }
		if ($0 ~ /^six/)   { digits = digits "6"; continue }
		if ($0 ~ /^seven/) { digits = digits "7"; continue }
		if ($0 ~ /^eight/) { digits = digits "8"; continue }
		if ($0 ~ /^nine/)  { digits = digits "9"; continue }
		if ($0 ~ /^[1-9]/)   digits = digits substr($0, 1, 1)
	}
	count += substr(digits, 1, 1) substr(digits, length(digits), 1)
}

END { print count }
