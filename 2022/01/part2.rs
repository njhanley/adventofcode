#!/usr/bin/env rust-script

impl<T> Pipe for T {}
trait Pipe: Sized {
	fn pipe<B, F>(self, f: F) -> B
	where
		F: FnOnce(Self) -> B,
	{
		f(self)
	}
}

std::fs::read_to_string("input.txt")
	.unwrap()
	.split("\n\n")
	.map(|record| {
		record
			.lines()
			.map(|line| line.parse::<u32>().unwrap())
			.sum::<u32>()
	})
	.collect::<Vec<_>>()
	.pipe(|mut calories| {
		calories.sort();
		calories
	})
	.iter()
	.rev()
	.take(3)
	.sum::<u32>()
	.pipe(|answer| println!("{answer}"));
