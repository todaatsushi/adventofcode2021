use std::str::FromStr;
use std::string::ParseError;

#[derive(Debug, Clone, Copy)]
enum Axis {
    Horizontal,
    Vertical,
}

#[derive(Debug, Clone, Copy)]
pub struct Fold {
    axis: Axis,
    point: usize,
}

impl FromStr for Fold {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let binding = s.trim().replace("fold along ", "");
        let fold = binding.split("=").collect::<Vec<&str>>();

        let axis = match fold[0] {
            "y" => Axis::Horizontal,
            "x" => Axis::Vertical,
            _ => panic!("Axis is not supported."),
        };

        Ok(Fold {
            axis,
            point: fold[1].parse::<usize>().unwrap(),
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
    fn get_board_size(points: &Vec<Point>) -> (u32, u32) {
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
        dimensions
    }

    fn create_empty_board(x: u32, y: u32) -> Vec<Vec<u32>> {
        let mut board: Vec<Vec<u32>> = Vec::new();

        for _ in 0..x {
            let mut row: Vec<u32> = Vec::new();
            for _ in 0..y {
                row.push(0);
            }
            board.push(row);
        }
        board
    }

    fn get_num_rows(self: &Self) -> usize {
        self.board.len()
    }

    fn get_num_cols(self: &Self) -> usize {
        self.board[0].len()
    }

    pub fn new(points: Vec<Point>) -> Self {
        let dimensions = Self::get_board_size(&points);
        let mut board = Self::create_empty_board(dimensions.0, dimensions.1);

        for point in points {
            board[point.y as usize][point.x as usize] += 1
        }

        Map { board }
    }

    pub fn display(self: &Self, show_board: bool, show_visible_points: bool) {
        if show_board {
            for row in &self.board {
                for col in row {
                    print!("{} ", col);
                }
                println!()
            }
        }

        if show_visible_points {
            let mut total_visible_points = 0;
            for row in &self.board {
                for point in row {
                    if *point == 1 {
                        total_visible_points += 1;
                    }
                }
            }

            println!("Total visible points: {}", total_visible_points);
        }
    }

    fn split_along(self: &Self, fold: &Fold) -> (Vec<Vec<u32>>, Vec<Vec<u32>>) {
        // TODO: Flip around to_fold & fold_into so that fold_into is the larger one
        match fold.axis {
            Axis::Horizontal => {
                let fold_into: Vec<Vec<u32>> = self.board[0..fold.point].into();
                let mut to_fold: Vec<Vec<u32>> =
                    self.board[fold.point as usize + 1..self.get_num_rows()].into();
                to_fold.reverse();

                (fold_into, to_fold)
            }
            Axis::Vertical => {
                let mut fold_into: Vec<Vec<u32>> = Vec::new();
                let mut to_fold: Vec<Vec<u32>> = Vec::new();

                for row in &self.board {
                    let before_split: Vec<u32> = row[0..fold.point].into();

                    let mut after_split: Vec<u32> =
                        row[fold.point as usize + 1..self.get_num_cols()].into();
                    after_split.reverse();

                    to_fold.push(after_split);
                    fold_into.push(before_split);
                }
                (fold_into, to_fold)
            }
        }
    }

    pub fn fold(self: &mut Self, fold: &Fold) {
        let (fold_into, to_fold) = self.split_along(&fold);
        self.board = fold_into;

        match fold.axis {
            Axis::Horizontal => {
                let start_point = self.get_num_rows() - to_fold.len();
                for to_fold_index in 0..to_fold.len() {
                    let fold_into_index = start_point + to_fold_index;

                    let fold_into = &mut self.board[fold_into_index];
                    let to_fold = &to_fold[to_fold_index];

                    for index in 0..fold_into.len() {
                        if fold_into[index] == 1 || to_fold[index] == 1 {
                            fold_into[index] = 1;
                        }
                    }
                }
            }
            Axis::Vertical => {
                for row_num in 0..self.get_num_rows() {
                    let to_fold_row = &to_fold[row_num];
                    let start_point = self.get_num_cols() - to_fold_row.len();
                    let fold_into_row = &mut self.board[row_num];

                    for to_fold_index in 0..to_fold_row.len() {
                        let fold_into_index = start_point + to_fold_index;

                        if fold_into_row[fold_into_index] == 1 || to_fold_row[to_fold_index] == 1 {
                            fold_into_row[fold_into_index] = 1;
                        }
                    }
                }
            }
        };
    }
}
