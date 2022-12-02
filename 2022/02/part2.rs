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
	.map(|strategy| match strategy {
		"A X" => "A C",
		"A Y" => "A A",
		"A Z" => "A B",
		"B X" => "B A",
		"B Y" => "B B",
		"B Z" => "B C",
		"C X" => "C B",
		"C Y" => "C C",
		"C Z" => "C A",
		_ => unreachable!(),
	})
	.map(|round| {
		(match round {
			"A C" | "B A" | "C B" => 0,
			"A A" | "B B" | "C C" => 3,
			"A B" | "B C" | "C A" => 6,
			_ => unreachable!(),
		}) + (match round.chars().last().unwrap() {
			'A' => 1,
			'B' => 2,
			'C' => 3,
			_ => unreachable!(),
		})
	})
	.sum::<u32>()
	.pipe(|score| println!("{score}"));
