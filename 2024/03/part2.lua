#!/usr/bin/env lua5.4

local input = io.read("a")

local total = 0
local enabled = true
for i = 1, #input do
	if input:match("^do%(%)", i) then
		enabled = true
	elseif input:match("^don't%(%)", i) then
		enabled = false
	elseif enabled then
		local x, y = input:match("^mul%((%d+),(%d+)%)", i)
		if x and y then
			total = total + x * y
		end
	end
end
print(total)
