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
            axis,
            point: fold[1].parse::<u32>().unwrap(),
        })
    }
}

#[derive(Debug)]
pub struct Point {
    x: u32,
    y: u32,
}

impl FromStr for Point {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let points = s
            .trim()
            .split(",")
            .map(|p| p.parse::<u32>().unwrap())
            .collect::<Vec<u32>>();
        Ok(Point {
            x: points[0],
            y: points[1],
        })
    }
}

#[derive(Debug)]
pub struct Map {
    board: Vec<Vec<u32>>,
}

impl Map {
    pub fn new(points: Vec<Point>) -> Self {
        let mut dimensions = points.iter().fold((0, 0), |mut dimensions, point| {
            if point.y > dimensions.0 {
                dimensions.0 = point.y;
            }

            if point.x > dimensions.1 {
                dimensions.1 = point.x;
            }
            dimensions
        });
        dimensions.0 += 1;
        dimensions.1 += 1;

        let mut board: Vec<Vec<u32>> = Vec::new();

        for _ in 0..dimensions.0 {
            let mut row: Vec<u32> = Vec::new();
            for _ in 0..dimensions.1 {
                row.push(0);
            }
            board.push(row);
        }

        for point in points {
            board[point.y as usize][point.x as usize] += 1
        }

        Map { board }
    }

    pub fn display(self: Self) {
        for row in self.board {
            for col in row {
                print!("{} ", col);
            }
            println!()
        }
    }
}
