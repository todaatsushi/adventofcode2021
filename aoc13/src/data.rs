use std::str::FromStr;
use std::string::ParseError;

#[derive(Debug)]
enum Axis {
    Horizontal,
    Vertical,
}

#[derive(Debug)]
pub struct Fold {
    axis: Axis,
    point: u32,
}

impl FromStr for Fold {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let binding = s.trim().replace("fold along ", "");
        let fold = binding.split("=").collect::<Vec<&str>>();

        let axis = match fold[0] {
            "x" => Axis::Horizontal,
            "y" => Axis::Vertical,
            _ => panic!("Axis is not supported."),
        };

        Ok(Fold {
            axis: axis,
            point: fold[1].parse::<u32>().unwrap(),
        })
    }
}
