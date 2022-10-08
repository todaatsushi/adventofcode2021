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

#[derive(Debug)]
pub struct Polymer {
    pub sequence: String,
    pub rules: Vec<Rule>,
}

impl Polymer {
    fn step(self: &mut Self) {}

    pub fn run_steps(self: &mut Self, num_steps: i32) {
        for _ in 0..num_steps {
            self.step();
        }
    }
}
