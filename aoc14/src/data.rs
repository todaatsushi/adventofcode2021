use std::collections::HashMap;
use std::str::FromStr;
use std::string::ParseError;

#[derive(Debug)]
pub struct Rule {
    pattern: String,
    insert: char,
}

impl FromStr for Rule {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (pattern, insert) = s.trim().split_once(" -> ").unwrap();
        Ok(Rule {
            pattern: pattern.to_string(),
            insert: insert.chars().next().unwrap(),
        })
    }
}

pub fn rules_to_hashmap(rules: Vec<Rule>) -> HashMap<String, char> {
    let mut rule_map: HashMap<String, char> = HashMap::new();

    for rule in rules {
        rule_map.entry(rule.pattern).or_insert(rule.insert);
    }
    rule_map
}

#[derive(Debug)]
pub struct Polymer {
    pub sequence: String,
    pub rules: HashMap<String, char>,
    pub tally: HashMap<String, u64>,
}

impl Polymer {
    pub fn new(sequence: String, rules: Vec<Rule>) -> Self {
        let rules = rules_to_hashmap(rules);

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
}
