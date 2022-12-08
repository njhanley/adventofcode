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

const LENGTH: usize = 4;
std::fs::read_to_string("input.txt")
	.unwrap()
	.chars()
	.collect::<Vec<_>>()
	.windows(LENGTH)
	.position(|window| {
		window
			.iter()
			.fold(0u128, |set, &c| set | 1 << c as u8)
			.count_ones() == LENGTH as u32
	})
	.unwrap()
	.pipe(|index| index + LENGTH)
	.pipe(|answer| println!("{answer}"));
