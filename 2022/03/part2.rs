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
	.collect::<Vec<_>>()
	.chunks(3)
	.map(|group| {
		let mut set = [0u8; 128];
		group
			.iter()
			.enumerate()
			.for_each(|(i, &line)| line.chars().for_each(|c| set[c as usize] |= 1 << i));
		set.iter().position(|&n| n == 0b111).unwrap() as u8
	})
	.map(|c| match c {
		b'a'..=b'z' => 1 + c - b'a',
		b'A'..=b'Z' => 27 + c - b'A',
		_ => unreachable!(),
	} as u32)
	.sum::<u32>()
	.pipe(|answer| println!("{answer}"));
