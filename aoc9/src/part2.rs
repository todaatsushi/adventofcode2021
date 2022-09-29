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
    let rows = grid.len() as i32;
    let cols = grid[0].len() as i32;

    let mut visited_points: HashSet<[i32; 2]> = HashSet::new();
    let mut sub_total: i32;
    let mut totals: Vec<i32> = Vec::new();

    for row in 0..rows {
        for col in 0..cols {
            sub_total = build_basin(row, col, &mut visited_points, &grid);
            if sub_total > 0 {
                totals.push(sub_total);
            }
        }
    }
    totals.sort();
    totals.reverse();

    let mut total = 1;

    // No guarantee of 3 elems usually, but the AOC spec says all non 9 squares are in a basin,
    // so make this unreasonable assumption.
    for i in 0..3 {
        total *= totals[i];
    }
    println!("Part 2 answer: {} (Test: {})", total, is_test);
}
