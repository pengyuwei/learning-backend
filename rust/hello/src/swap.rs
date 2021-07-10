pub fn swap1(_a: i32, _b: i32) -> i32 {
    let mut a = 1;
    let mut b = 2;
    let c;
    c = a;
    a = b;
    b = c;

    println!("{0} {1}", a, b); // 2 1
    return 0;
}

pub fn swap2(_a: i32, _b: i32) -> i32 {
    let mut a = 1;
    let mut b = 2;
    a = a + b; // a=3;b=2
    b = a - b; // b=1
    a = a - b; // a=2

    println!("{0} {1}", a, b); // 2 1
    return 0;
}