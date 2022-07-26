use std::io;
use std::cmp::Ordering;
use rand::Rng;

fn main() {
    println!("Guess the number!");
    
    let secret_num = rand::thread_rng().gen_range(1, 101); // [1, 100]
    println!("secret_num is: {}", secret_num);
    
    loop {
        println!("Please input your guess.");
    
        let mut guess = String::new();
    
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");
    
        let guess: u32 = guess.trim().parse() // overwrite string guess var
            .expect("Please type a number!"); 
    
        println!("You guessed: {}", guess);
    
        match guess.cmp(&secret_num) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
