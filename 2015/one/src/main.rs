use std::fs::File;
use std::io::BufReader;
use std::io::{self, BufRead};

const UP: char = '(';
const DOWN: char = ')';

fn main() -> io::Result<()> {
    let f = File::open("input.txt").expect("Error reading file");
    let reader = BufReader::new(f);

    let mut parens = String::new();
    let mut count = 0;

    for line in reader.lines() {
        parens = parens + &line?;
    }

    for (idx, item) in parens.char_indices() {
        if count == -1 {
            // The previous iteration would have set the count to -1. No need to increase idx position.
            println!(
                "Position of the char that caused santa to enter the basement: {}",
                idx
            );
            break;
        }

        match item {
            UP => count += 1,
            DOWN => count -= 1,
            _ => continue,
        }
    }

    // println!("Santa needs to go to floor {count}");

    Ok(())
}
