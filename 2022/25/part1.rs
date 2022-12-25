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
	.map(|line| {
		let mut number = line
			.chars()
			.map(|c| match c {
				'2' => 2,
				'1' => 1,
				'0' => 0,
				'-' => -1,
				'=' => -2,
				_ => unreachable!(),
			})
			.collect::<Vec<i8>>();
		number.reverse();
		number
	})
	.reduce(|a, b| {
		let (mut a, mut b, mut c) = (a.into_iter(), b.into_iter(), Vec::new());
		loop {
			match (a.next(), b.next()) {
				(Some(a), Some(b)) => c.push(a + b),
				(Some(n), None) | (None, Some(n)) => c.push(n),
				(None, None) => break c,
			}
		}
	})
	.unwrap()
	.into_iter()
	.scan(0, |n, mut x| {
		(*n, x) = (0, x + *n);
		while x > 2 {
			(*n, x) = (*n + 1, x - 5);
		}
		while x < -2 {
			(*n, x) = (*n - 1, x + 5);
		}
		Some(x)
	})
	.collect::<Vec<_>>()
	.pipe(|mut v| {
		v.reverse();
		v.into_iter()
	})
	.map(|x| match x {
		2 => '2',
		1 => '1',
		0 => '0',
		-1 => '-',
		-2 => '=',
		_ => unreachable!(),
	})
	.collect::<String>()
	.pipe(|answer| println!("{answer}"));
