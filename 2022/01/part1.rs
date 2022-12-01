#!/usr/bin/env rust-script

let answer = std::fs::read_to_string("input.txt")?
	.split("\n\n")
	.map(|record| record.lines()
		.map(|line| line.parse::<u32>().unwrap())
		.sum::<u32>())
	.max().unwrap();
println!("{}", answer);
