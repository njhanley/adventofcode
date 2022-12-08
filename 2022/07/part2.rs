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

fn join(path: &mut String, name: &str) {
	if !path.ends_with('/') {
		path.push('/');
	}
	path.push_str(name);
}

fn dirname(path: &mut String) {
	let index = path.rfind('/').unwrap();
	path.truncate(index);
	if path.is_empty() {
		path.push('/');
	}
}

const CAPACITY: u32 = 70000000;
const UPDATE_SIZE: u32 = 30000000;
std::fs::read_to_string("input.txt")
	.unwrap()
	.lines()
	.filter(|&line| !(line == "$ ls" || line.starts_with("dir")))
	.pipe(|lines| {
		let mut path = String::new();
		lines.filter_map(move |line| {
			if let Some(arg) = line.strip_prefix("$ cd ") {
				match arg {
					"/" => {
						path.clear();
						path.push('/');
					}
					".." => dirname(&mut path),
					_ => join(&mut path, arg),
				}
				None
			} else {
				let (size, name) = line.split_once(' ').unwrap();
				let mut path = path.clone();
				join(&mut path, name);
				Some((path, size.parse::<u32>().unwrap()))
			}
		})
	})
	.fold(HashMap::new(), |mut directories, (mut path, size)| {
		while path != "/" {
			dirname(&mut path);
			*directories.entry(path.clone()).or_insert(0) += size;
		}
		directories
	})
	.pipe(|directories| {
		let available = CAPACITY - directories["/"];
		let need = UPDATE_SIZE - available;
		directories
			.into_iter()
			.filter(move |(_, size)| *size >= need)
	})
	.min_by_key(|(_, size)| *size)
	.unwrap()
	.pipe(|(_, size)| size)
	.pipe(|answer| println!("{answer}"));
