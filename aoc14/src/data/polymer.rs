use crate::data::rule;
use std::collections::HashMap;

#[derive(Debug)]
pub struct Polymer {
    pub sequence: String,
    pub rules: HashMap<String, char>,
    pub tally: HashMap<String, u64>,
}

impl Polymer {
    pub fn new(sequence: String, rules: Vec<rule::Rule>) -> Self {
        let rules = rule::rules_to_hashmap(rules);

        let mut tally: HashMap<String, u64> = HashMap::new();
        for index in 0..sequence.len() - 1 {
            let substr = &sequence[index..index + 2].to_string();
            let count = tally.entry(substr.to_string()).or_insert(0);
            *count += 1;
        }

        Self {
            sequence,
            rules,
            tally,
        }
    }

    fn update_tally(self: &mut Self, new_tally: HashMap<String, u64>) {
        for (pair, count) in new_tally {
            self.tally
                .entry(pair)
                .and_modify(|c| *c += count)
                .or_insert(count);
        }
    }

    fn step(self: &mut Self) {
        let mut new_tally_items: HashMap<String, u64> = HashMap::new();

        let elems_to_split = self
            .tally
            .iter_mut()
            .filter(|(_, count)| *count > &mut 0)
            .filter(|(pair, _)| self.rules.contains_key(*pair));

        for (pair, count) in elems_to_split {
            let to_insert = self.rules.get(pair).unwrap();

            // Create the new pairs (ie. pair[0] + to_insert, to_insert + pair[1]) and
            // add to new tally.
            let mut prepended: String = to_insert.clone().to_string();
            prepended.push(pair.chars().last().unwrap());

            let mut appended: String = pair.chars().next().unwrap().to_string();
            appended.push(to_insert.clone());

            for new_item in [appended, prepended] {
                let new_count = new_tally_items.entry(new_item).or_insert(0);
                *new_count += *count;
            }

            *count = 0;
        }

        self.update_tally(new_tally_items);
    }

    pub fn run_steps(self: &mut Self, num_steps: u64) {
        for _ in 0..num_steps {
            self.step();
        }
    }

    pub fn count_chars(self: &Self) {
        let mut counts: HashMap<String, u64> = HashMap::new();
        for (pair, count) in self.tally.iter().filter(|(_, count)| count > &&0) {
            let c = pair.chars().last().unwrap();
            let current = counts.entry(c.to_string()).or_insert(0);
            *current += count;
        }
        let c = self.sequence.chars().next().unwrap();
        let current = counts.entry(c.to_string()).or_insert(0);
        *current += 1;

        println!("Char counts: {:?}", counts);

        let mut max: u64 = *counts.values().next().unwrap();
        let mut min: u64 = *counts.values().next().unwrap();

        for count in counts.values() {
            if max < *count {
                max = *count;
            }

            if min > *count {
                min = *count;
            }
        }
        println!(
            "Difference between max and min: {} ({}, {})",
            max - min,
            max,
            min
        );
    }
}
