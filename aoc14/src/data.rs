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
}

impl Polymer {
    fn step(self: &mut Self) {
        let (mut left, mut right) = (0 as usize, 2 as usize);

        while right <= self.sequence.len() && left < right {
            let subpattern = self.sequence[left..right].to_string();

            match self.rules.get(&subpattern) {
                Some(c) => {
                    self.sequence.insert(left + 1 as usize, *c);
                    left += 2;
                    right = left + 2;
                }
                None => {
                    left += 1;
                    right += 1;
                }
            };
        }
    }

    pub fn run_steps(self: &mut Self, num_steps: i32) {
        for _ in 0..num_steps {
            self.step();
        }
    }
}
