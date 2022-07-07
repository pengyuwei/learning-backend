use std::io;
use rand::Rng;

fn main() {
    println!("Guess the number!");
    
    let secret_num = rand::thread_rng().gen_range(1, 101); // [1, 100]
    println!("secret_num is: {}", secret_num);
    
    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {}", guess);
}
