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

use Instruction::*;
#[derive(Clone, Copy, Debug)]
enum Instruction {
	Addx(i32),
	Noop,
}

std::fs::read_to_string("input.txt")
	.unwrap()
	.lines()
	.map(|line| {
		if let Some(("addx", v)) = line.split_once(' ') {
			Addx(v.parse().unwrap())
		} else if line == "noop" {
			Noop
		} else {
			unreachable!()
		}
	})
	.scan((0, 1), |state, op| {
		match op {
			Addx(v) => {
				state.0 += 2;
				state.1 += v;
			}
			Noop => {
				state.0 += 1;
			}
		};
		Some(*state)
	})
	.collect::<Vec<_>>()
	.pipe(|history| {
		[20, 60, 100, 140, 180, 220]
			.iter()
			.map(|&cycle| {
				history
					.iter()
					.rfind(|(i, _)| *i < cycle)
					.unwrap()
					.pipe(|(_, x)| cycle * x)
			})
			.sum::<i32>()
	})
	.pipe(|answer| println!("{answer}"));
