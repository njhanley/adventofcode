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

const LENGTH: usize = 14;
std::fs::read_to_string("input.txt")
	.unwrap()
	.chars()
	.collect::<Vec<_>>()
	.windows(LENGTH)
	.position(|window| {
		[0u8; 128].pipe(|mut set| {
			window.iter().for_each(|&c| set[c as usize] += 1);
			!set.iter().any(|&n| n > 1)
		})
	})
	.unwrap()
	.pipe(|index| index + LENGTH)
	.pipe(|answer| println!("{answer}"));
