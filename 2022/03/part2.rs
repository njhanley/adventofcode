#!/usr/bin/env rust-script

use std::collections::{hash_map::RandomState, HashSet};
use std::iter::FromIterator;

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
	.collect::<Vec<_>>()
	.chunks(3)
	.map(|group| {
		*group
			.iter()
			.map(|rucksack| HashSet::from_iter(rucksack.chars()))
			.reduce(|a: HashSet<char, RandomState>, b| &a & &b)
			.iter()
			.next()
			.unwrap()
			.iter()
			.next()
			.unwrap()
	})
	.map(|c| match c {
		'a'..='z' => 1 + c as u8 - b'a',
		'A'..='Z' => 27 + c as u8 - b'A',
		_ => unreachable!(),
	} as u32)
	.sum::<u32>()
	.pipe(|answer| println!("{answer}"));
