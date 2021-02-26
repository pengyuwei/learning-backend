#[derive(Debug)]

struct User {
    name: String,
    password: String,
    age: u32
}

impl User {
    fn token(&self) -> u32 {
        self.age * 2
    }
    fn create() -> User {
        User{name:String::from("default"), password:String::from("default"), age:0}
    }
}

pub fn test_struct() {
    let name = String::from("Tom");
    let Tom = User {
        name,
        password: String::from("123456"),
        age: 42
    };

    let Jerry = User {
        password: String::from("654321"),
        ..Tom
    }; // Tom moved

    struct Color(u8, u8, u8);
    struct Point(f64, f64);
    let black = Color(0, 0, 0);
    let origin = Point(0.0, 0.0);
    println!("black = ({}, {}, {})", black.0, black.1, black.2);
    println!("origin = ({}, {})", origin.0, origin.1);

    println!("Jerry is {:?}", Jerry); // because of #[derive(Debug)]
    println!("Jerry.token is {:?}", Jerry.token());
    println!("Jerry.token is {:?}", User::create());
}