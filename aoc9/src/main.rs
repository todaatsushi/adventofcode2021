fn parse_input(test: bool) -> Vec<Vec<i32>> {
    let file_contents = match test {
        true => include_str!("../inputs/test.txt").lines(),
        false => include_str!("../inputs/puzzle.txt").lines(),
    };
    let mut points: Vec<Vec<i32>> = Vec::new();

    for line in file_contents {
        let line_points = line
            .split("")
            .filter(|&height| !height.is_empty())
            .map(|x| x.trim().parse::<i32>().unwrap())
            .collect();

        let _ = &points.push(line_points);
    }

    points
}

fn should_add_point(x: i32, y: i32, grid: &Vec<Vec<i32>>, moves: &[[i32; 2]; 4]) -> bool {
    let mut add = true;
    let num_rows = grid.len() as i32;
    let num_cols = grid[0].len() as i32;
    let current = grid[y as usize][x as usize];

    for alter in moves.iter() {
        let row = y + alter[0];
        let col = x + alter[1];

        let row_in_range = row < num_rows && row >= 0;
        let col_in_range = col < num_cols && col >= 0;

        if row_in_range && col_in_range {
            let adjacent_value = grid[row as usize][col as usize];
            if current >= adjacent_value {
                add = false;
            }
        }
    }
    add
}

fn main() {
    let grid = parse_input(true);
    let rows = grid.len() as i32;
    let cols = grid[0].len() as i32;
    let mut risk_score = 0;
    let directions = [[1, 0], [-1, 0], [0, -1], [0, 1]];

    for row_index in 0..rows {
        for col_index in 0..cols {
            let should_add = should_add_point(col_index, row_index, &grid, &directions);
            if should_add {
                risk_score += &grid[row_index as usize][col_index as usize] + 1;
            }
        }
    }
    println!("{}", risk_score);
}
