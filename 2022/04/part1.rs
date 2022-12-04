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
	.map(|line| line.split_once(',').unwrap())
	.map(|(a, b)| (a.split_once('-').unwrap(), b.split_once('-').unwrap()))
	.map(|((a, b), (c, d))| {
		(
			a.parse::<u8>().unwrap(),
			b.parse::<u8>().unwrap(),
			c.parse::<u8>().unwrap(),
			d.parse::<u8>().unwrap(),
		)
	})
	.filter(|(a, b, c, d)| a <= c && d <= b || c <= a && b <= d)
	.count()
	.pipe(|answer| println!("{answer}"));
