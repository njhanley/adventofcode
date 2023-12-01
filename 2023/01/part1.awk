#!/usr/bin/awk -f

{
	gsub(/[a-z]/, "")
	count += substr($0, 1, 1) substr($0, length(), 1)
}

END { print count }
