#!/usr/bin/env rust-script

use std::collections::VecDeque;

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
	.split("\n\n")
	.map(|record| {
		let mut lines = record.lines().skip(1);
		let items = lines
			.next()
			.unwrap()
			.strip_prefix("  Starting items: ")
			.unwrap()
			.split(", ")
			.map(|n| n.parse::<u32>().unwrap())
			.collect::<VecDeque<_>>();
		let operation = lines
			.next()
			.unwrap()
			.strip_prefix("  Operation: new = old ")
			.unwrap()
			.split_once(' ')
			.unwrap()
			.pipe(|(op, arg)| -> Box<dyn Fn(u32) -> u32> {
				let op = match op {
					"+" => |a, b| a + b,
					"*" => |a, b| a * b,
					_ => unreachable!(),
				};
				if arg == "old" {
					Box::new(move |item| op(item, item))
				} else {
					let arg = arg.parse::<u32>().unwrap();
					Box::new(move |item| op(item, arg))
				}
			});
		let divisor = lines
			.next()
			.unwrap()
			.strip_prefix("  Test: divisible by ")
			.unwrap()
			.parse::<u32>()
			.unwrap();
		let yes = lines
			.next()
			.unwrap()
			.strip_prefix("    If true: throw to monkey ")
			.unwrap()
			.parse::<usize>()
			.unwrap();
		let no = lines
			.next()
			.unwrap()
			.strip_prefix("    If false: throw to monkey ")
			.unwrap()
			.parse::<usize>()
			.unwrap();
		let test = move |item| if item % divisor == 0 { yes } else { no };
		(items, operation, test)
	})
	.fold(
		(Vec::new(), Vec::new()),
		|mut v, (items, operation, test)| {
			v.0.push(items);
			v.1.push((operation, test));
			v
		},
	)
	.pipe(|(mut items, monkeys)| {
		let mut inspected = vec![0; monkeys.len()];
		for _ in 0..20 {
			for (i, (operation, test)) in monkeys.iter().enumerate() {
				while let Some(item) = items[i].pop_front() {
					inspected[i] += 1;
					let item = operation(item) / 3;
					items[test(item)].push_back(item);
				}
			}
		}
		inspected.sort_by(|a, b| b.cmp(a));
		inspected.truncate(2);
		inspected.iter().product::<u32>()
	})
	.pipe(|answer| println!("{answer}"));
