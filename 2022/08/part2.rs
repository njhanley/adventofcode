#!/usr/bin/env rust-script

use std::collections::HashMap;

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
	.enumerate()
	.pipe(|lines| {
		let mut map = HashMap::new();
		let mut width = 0;
		let mut height = 0;
		lines.for_each(|(y, line)| {
			line.chars().enumerate().for_each(|(x, c)| {
				map.insert((x, y), (c as u8 - b'0') as i8);
				width = x + 1;
			});
			height = y + 1;
		});
		let mut scores = HashMap::new();
		map.iter().for_each(|(&(x, y), &n)| {
			let mut score = [0; 4];
			for x in (0..x).rev() {
				score[0] += 1;
				if n <= map[&(x, y)] {
					break;
				}
			}
			for x in (x + 1)..width {
				score[1] += 1;
				if n <= map[&(x, y)] {
					break;
				}
			}
			for y in (0..y).rev() {
				score[2] += 1;
				if n <= map[&(x, y)] {
					break;
				}
			}
			for y in (y + 1)..height {
				score[3] += 1;
				if n <= map[&(x, y)] {
					break;
				}
			}
			scores.insert((x, y), score.iter().product::<u32>());
		});
		scores
	})
	.values()
	.max()
	.unwrap()
	.pipe(|answer| println!("{answer}"));
