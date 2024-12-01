#!/usr/bin/env lua5.4

local l, r = {}, {}
for line in io.lines() do
	local a, b = line:match("^(%d+)%s+(%d+)$")
	table.insert(l, tonumber(a))
	table.insert(r, tonumber(b))
end

table.sort(l)
table.sort(r)

local total = 0
for i in ipairs(l) do
	total = total + math.abs(r[i] - l[i])
end
print(total)
