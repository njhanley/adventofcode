#!/usr/bin/env lua5.4

local total = 0
for x, y in io.read("a"):gmatch("mul%((%d+),(%d+)%)") do
	total = total + x * y
end
print(total)
