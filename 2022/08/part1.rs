#!/usr/bin/env rust-script

use std::collections::{HashMap, HashSet};

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
		let mut visible = HashSet::new();
		for x in 0..width {
			let mut n = -1;
			for y in 0..height {
				if let Some(&m) = map.get(&(x, y)) {
					if m > n {
						visible.insert((x, y));
						n = m;
					}
				}
			}
			n = -1;
			for y in (0..height).rev() {
				if let Some(&m) = map.get(&(x, y)) {
					if m > n {
						visible.insert((x, y));
						n = m;
					}
				}
			}
		}
		for y in 0..height {
			let mut n = -1;
			for x in 0..width {
				if let Some(&m) = map.get(&(x, y)) {
					if m > n {
						visible.insert((x, y));
						n = m;
					}
				}
			}
			n = -1;
			for x in (0..width).rev() {
				if let Some(&m) = map.get(&(x, y)) {
					if m > n {
						visible.insert((x, y));
						n = m;
					}
				}
			}
		}
		visible
	})
	.len()
	.pipe(|answer| println!("{answer}"));
