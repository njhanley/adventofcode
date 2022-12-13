#!/usr/bin/env rust-script

use std::collections::HashSet;

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
	.map(|line| line.split_once(' ').unwrap())
	.map(|(c, n)| (c, n.parse::<i32>().unwrap()))
	.fold(
		((0, 0), (0, 0), HashSet::new()),
		|state, (direction, steps)| {
			(0..steps).fold(state, |(head, tail, mut visited), _| {
				let head = match direction {
					"U" => (head.0, head.1 + 1),
					"D" => (head.0, head.1 - 1),
					"R" => (head.0 + 1, head.1),
					"L" => (head.0 - 1, head.1),
					_ => unreachable!(),
				};
				let tail = match direction {
					"U" => {
						if head.1 > tail.1 + 1 {
							(head.0, head.1 - 1)
						} else {
							tail
						}
					}
					"D" => {
						if head.1 < tail.1 - 1 {
							(head.0, head.1 + 1)
						} else {
							tail
						}
					}
					"R" => {
						if head.0 > tail.0 + 1 {
							(head.0 - 1, head.1)
						} else {
							tail
						}
					}
					"L" => {
						if head.0 < tail.0 - 1 {
							(head.0 + 1, head.1)
						} else {
							tail
						}
					}
					_ => unreachable!(),
				};
				visited.insert(tail);
				(head, tail, visited)
			})
		},
	)
	.pipe(|(_, _, visited)| visited.len())
	.pipe(|answer| println!("{answer}"));
