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
	.map(|line| line.split_at(line.len() / 2))
	.map(|(a, b)| {
		let i = a.find(|c| b.contains(c)).unwrap();
		a.chars().nth(i).unwrap()
	})
	.map(|c| match c {
		'a'..='z' => 1 + c as u8 - b'a',
		'A'..='Z' => 27 + c as u8 - b'A',
		_ => unreachable!(),
	} as u32)
	.sum::<u32>()
	.pipe(|answer| println!("{answer}"));
