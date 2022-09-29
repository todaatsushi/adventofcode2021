use crate::utils::parse_input;
use std::collections::{HashSet, VecDeque};

fn build_basin(
    row: i32,
    col: i32,
    visited_points: &mut HashSet<[i32; 2]>,
    grid: &Vec<Vec<i32>>,
) -> i32 {
    let moves = [[1, 0], [-1, 0], [0, 1], [0, -1]];
    let mut to_visit: VecDeque<[i32; 2]> = VecDeque::new();
    to_visit.push_back([row, col]);

    let mut sub_total = 0;

    while to_visit.len() > 0 {
        let [row, col] = to_visit.pop_front().unwrap();

        if grid[row as usize][col as usize] == 9 {
            continue;
        }

        for alter in moves {
            let new_row = row + alter[0];
            let new_col = col + alter[1];
            let row_in_range = new_row < grid.len().try_into().unwrap() && new_row >= 0;
            let col_in_range = new_col < grid[0].len().try_into().unwrap() && new_col >= 0;

            if row_in_range
                && col_in_range
                && visited_points.contains(&[new_row, new_col]) == false
                && grid[new_row as usize][new_col as usize] != 9
            {
                let _ = &to_visit.push_back([new_row, new_col]);
                visited_points.insert([new_row, new_col]);
                sub_total += 1;
            }
        }
    }
    sub_total
}

pub fn solve_part_2(is_test: bool) {
    let grid = parse_input(is_test);

    println!("{:?}", grid);
}
