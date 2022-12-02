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
	.lines()
	.map(|round| {
		(match round {
			"A Z" | "B X" | "C Y" => 0,
			"A X" | "B Y" | "C Z" => 3,
			"A Y" | "B Z" | "C X" => 6,
			_ => 0,
		}) + (match round.chars().last().unwrap() {
			'X' => 1,
			'Y' => 2,
			'Z' => 3,
			_ => 0,
		})
	})
	.sum::<u32>()
	.pipe(|score| println!("{score}"));
