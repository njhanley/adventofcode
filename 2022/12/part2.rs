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

// A* implementation curtesy of ChatGPT
use std::cmp::Ordering;
use std::collections::{BinaryHeap, HashMap, HashSet};

#[derive(Eq, PartialEq, Debug)]
struct Node {
	cost: u32,
	position: (i32, i32),
}

impl Ord for Node {
	fn cmp(&self, other: &Node) -> Ordering {
		other.cost.cmp(&self.cost)
	}
}

impl PartialOrd for Node {
	fn partial_cmp(&self, other: &Node) -> Option<Ordering> {
		Some(self.cmp(other))
	}
}

fn a_star(
	start: (i32, i32),
	goal: (i32, i32),
	heuristic: impl Fn((i32, i32)) -> u32,
	neighbors: impl Fn((i32, i32)) -> Vec<(i32, i32)>,
) -> Option<(Vec<(i32, i32)>, u32)> {
	let mut open = BinaryHeap::new();
	let mut closed = HashSet::new();
	let mut came_from = HashMap::new();
	let mut cost_so_far = HashMap::new();

	open.push(Node {
		cost: 0,
		position: start,
	});
	cost_so_far.insert(start, 0);

	while let Some(current) = open.pop() {
		if current.position == goal {
			let mut path = Vec::new();
			let mut current = current.position;
			path.push(current);
			while let Some(next) = came_from.get(&current) {
				current = *next;
				path.push(current);
			}
			path.reverse();
			return Some((path, *cost_so_far.get(&goal).unwrap()));
		}

		closed.insert(current.position);

		for position in neighbors(current.position) {
			let new_cost = cost_so_far[&current.position] + 1;
			if !cost_so_far.contains_key(&position) || new_cost < cost_so_far[&position] {
				cost_so_far.insert(position, new_cost);
				let priority = new_cost + heuristic(position);
				open.push(Node {
					cost: priority,
					position,
				});
				came_from.insert(position, current.position);
			}
		}
	}

	None
}
// End of ChatGPT generated code

std::fs::read_to_string("input.txt")
	.unwrap()
	.pipe(|input| {
		let mut heightmap = HashMap::new();
		let (mut x, mut y) = (0, 0);
		let mut end = None;
		for mut c in input.bytes() {
			if c == b'\n' {
				(x, y) = (0, y + 1);
				continue;
			}
			if c == b'S' {
				c = b'a';
			} else if c == b'E' {
				end = Some((x, y));
				c = b'z';
			}
			heightmap.insert((x, y), c);
			x += 1;
		}
		let starts = heightmap
			.iter()
			.filter_map(|(&p, &c)| if c == b'a' { Some(p) } else { None })
			.collect::<Vec<_>>();
		(starts, end.unwrap(), heightmap)
	})
	.pipe(|(starts, end, heightmap)| {
		starts
			.iter()
			.filter_map(|&start| {
				a_star(
					start,
					end,
					|_| 0, // unused
					|(x, y)| {
						let c = heightmap[&(x, y)];
						[(0, 1), (0, -1), (1, 0), (-1, 0)]
							.iter()
							.map(|(dx, dy)| (x + dx, y + dy))
							.filter(|p| heightmap.get(&p).filter(|&&d| c + 1 >= d).is_some())
							.collect()
					},
				)
				.map(|(_, cost)| cost)
			})
			.collect::<Vec<_>>()
	})
	.iter()
	.min()
	.unwrap()
	.pipe(|x| println!("{x:?}"));
