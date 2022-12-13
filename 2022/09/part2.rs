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
		((0i32, 0i32), [(0, 0); 9], HashSet::new()),
		|state, (direction, steps)| {
			(0..steps).fold(state, |(head, mut knots, mut visited), _| {
				let head = match direction {
					"U" => (head.0, head.1 + 1),
					"D" => (head.0, head.1 - 1),
					"R" => (head.0 + 1, head.1),
					"L" => (head.0 - 1, head.1),
					_ => unreachable!(),
				};
				let mut prev = head;
				knots.iter_mut().for_each(|knot| {
					let pull = (prev.0 - knot.0, prev.1 - knot.1);
					if pull.0.abs() > 1 || pull.1.abs() > 1 {
						knot.0 += pull.0.signum();
						knot.1 += pull.1.signum();
					}
					prev = *knot;
				});
				visited.insert(prev);
				(head, knots, visited)
			})
		},
	)
	.pipe(|(_, _, visited)| visited.len())
	.pipe(|answer| println!("{answer}"));
