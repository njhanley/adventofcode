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
	.split_once("\n\n")
	.unwrap()
	.pipe(|(initial, procedure)| {
		(
			initial
				.lines()
				.rev()
				.skip(1)
				.map(|line| line.chars().skip(1).step_by(4).collect::<Vec<_>>())
				.collect::<Vec<_>>(),
			procedure
				.lines()
				.map(|line| {
					line.split(' ')
						.filter_map(|field| field.parse::<u8>().ok())
						.collect::<Vec<_>>()
				})
				.map(|fields| {
					if let [a, b, c] = fields[..] {
						(a, b - 1, c - 1)
					} else {
						unreachable!()
					}
				})
				.collect::<Vec<_>>(),
		)
	})
	.pipe(|(rows, moves)| {
		let mut stacks = rows
			.first()
			.unwrap()
			.iter()
			.map(|_| Vec::new())
			.collect::<Vec<_>>();
		for (i, stack) in stacks.iter_mut().enumerate() {
			for row in rows.clone().into_iter() {
				if row[i] != ' ' {
					stack.push(row[i]);
				}
			}
		}
		(stacks, moves)
	})
	.pipe(|(mut stacks, moves)| {
		let mut tmp = Vec::new();
		for (n, from, to) in moves {
			for _ in 0..n {
				let c = stacks[from as usize].pop().unwrap();
				tmp.push(c);
			}
			for _ in 0..n {
				let c = tmp.pop().unwrap();
				stacks[to as usize].push(c);
			}
		}
		stacks
	})
	.iter()
	.map(|stack| stack.last().unwrap())
	.collect::<String>()
	.pipe(|answer| println!("{answer}"));
