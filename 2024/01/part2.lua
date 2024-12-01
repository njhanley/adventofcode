#!/usr/bin/env lua5.4

local l, r = {}, setmetatable({}, { __index = function () return 0 end })
for line in io.lines() do
	local a, b = line:match("^(%d+)%s+(%d+)$")
	a, b = tonumber(a), tonumber(b)
	table.insert(l, a)
	r[b] = r[b] + 1
end

local total = 0
for _, x in ipairs(l) do
	total = total + (x * r[x])
end
print(total)
