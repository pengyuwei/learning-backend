pub fn test_tuple() {
    let tup2 = (100, 3.14, 1);
    let (x, y, z) = tup2;
    println!("xyz is {},{},{}", x, y, z);
    
    let tx: (i32, f64, u8) = (100, 3.14, 1);
    let one = tx.0;
    let two = tx.1;
    println!("0-1 is {},{}", one, two);
}

pub fn test_array() {
    let a = [1,2,3,4];
    // let var: [array_type, count]
    let b: [i32; 5] = [1, 2, 3, 4, 5];
    println!("array[0] is {}", b[0]);
}