#!/usr/bin/env rust-script

let mut counts = std::fs::read_to_string("input.txt")?
	.split("\n\n")
	.map(|record| record.lines()
		.map(|line| line.parse::<u32>().unwrap())
		.sum::<u32>())
	.collect::<Vec<_>>();
counts.sort();
let answer = counts.iter()
	.rev()
	.take(3)
	.sum::<u32>();
println!("{}", answer);
