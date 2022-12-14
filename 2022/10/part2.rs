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

const WIDTH: i32 = 40;
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
	.pipe(|instructions| {
		let mut cycle = 0;
		let mut draw = |x| {
			let pixel = (cycle % WIDTH) + 1;
			if x - 1 <= pixel && pixel <= x + 1 {
				print!("#");
			} else {
				print!(".");
			}
			if pixel == WIDTH {
				println!();
			}
			cycle += 1;
		};
		let mut x = 2;
		for op in instructions {
			match op {
				Addx(v) => {
					draw(x);
					draw(x);
					x += v;
				}
				Noop => {
					draw(x);
				}
			}
		}
	});
