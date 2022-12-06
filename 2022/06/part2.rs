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
		let mut set = [0u8; 128];
		for &c in window {
			if set[c as usize] > 0 {
				return false;
			} else {
				set[c as usize] += 1;
			}
		}
		return true;
	})
	.unwrap()
	.pipe(|index| index + LENGTH)
	.pipe(|answer| println!("{answer}"));
