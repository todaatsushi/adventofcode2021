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
