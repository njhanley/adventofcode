#!/usr/bin/awk -f

{
	digits = ""
	while ($0) {
		     if ($0 ~ /^one/)   digits = digits "1"
		else if ($0 ~ /^two/)   digits = digits "2"
		else if ($0 ~ /^three/) digits = digits "3"
		else if ($0 ~ /^four/)  digits = digits "4"
		else if ($0 ~ /^five/)  digits = digits "5"
		else if ($0 ~ /^six/)   digits = digits "6"
		else if ($0 ~ /^seven/) digits = digits "7"
		else if ($0 ~ /^eight/) digits = digits "8"
		else if ($0 ~ /^nine/)  digits = digits "9"
		else if ($0 ~ /^[1-9]/) digits = digits substr($0, 1, 1)
		$0 = substr($0, 2)
	}
	count += substr(digits, 1, 1) substr(digits, length(digits), 1)
}

END { print count }
