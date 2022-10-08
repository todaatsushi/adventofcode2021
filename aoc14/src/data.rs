use std::str::FromStr;
use std::string::ParseError;

pub struct Rule {
    pattern: String,
    insert: String,
}

impl FromStr for Rule {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (pattern, insert) = s.trim().split_once(" -> ").unwrap();
        Ok(Rule {
            pattern: pattern.to_string(),
            insert: insert.to_string(),
        })
    }
}
