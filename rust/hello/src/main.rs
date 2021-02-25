fn add(a: i32, b: i32) -> i32 {
    return a + b;
}

fn test_if() {
    let xx = 5;
    if xx < 5 { // 条件可以没有小括号, int不能自动转换成bool
        println!("test_if/条件为 true");
    } else if xx == 5 {
        println!("test_if/刚好");
    } else {
        println!("test_if/条件为 false");
    }

    let number = if xx > 0 { 1 } else { -1 };
    assert!(number == 1);
}

fn test_loop() {
    let mut number = 1;
    while number != 4 {
        println!("test_loop/{}", number);
        number += 1;
    }

    let a = [10, 20, 30, 40, 50];
    for i in a.iter() {
        println!("test_loop/值为 : {}", i);
    }

    for i in 0..5 {
        println!("test_loop/a[{}] = {}", i, a[i]);
    }

    let s = ['R', 'U', 'S', 'T'];
    let mut i = 0;
    loop {
        let ch = s[i];
        if ch == 'S' {
            break;
        }
        println!("test_loop/\'{}\'", ch);
        i += 1;
    }
    let location = loop {
        let ch = s[i];
        if ch == 'T' {
            break i;
        }
        i += 1;
    };
    println!("test_loop/\'T\' 的索引为 {}", location);

    return;
}

fn _test_mem_helper1(str: String) -> String {
    assert!(str == "hello");
    return str;
}
fn _test_mem_helper2(p: i32) {
    assert!(p == 1);
}
fn _test_mem_helper3(str: &String) -> usize {
    assert!(str == "hello");
    return str.len();
}
fn test_mem() {
    // move
    let s1 = String::from("hello");
    let s2 = s1; // now s1 un-use-able
    
    // clone
    let s3 = s2.clone(); // s2 & s3 are both use-able
    assert!(s3 == "hello");
    
    // scope
    let s4 = _test_mem_helper1(s3);
    // now s3 is un-use-able because it's moved
    let x = 1;
    _test_mem_helper2(x);
    // x1 use-able
    assert!(x == 1);

    let s5 = &s4; // 引用不会获得所有权，因此不能修改变量的值，只能租借一份引用给别人
    println!("s4={}, s5={}", s4, s5);
    let len = _test_mem_helper3(&s5);
    assert!(len == 5);
}

/// cargo doc
/// hello
fn main() {
    println!("Hello, world!");

    let a = 12; // read-only var
    println!("a is {}", a);
    let a = 123; // re-bind var a
    println!("a is {0}, a again is {0}/{1}", a, 2);

    println!("{{}}");

    let mut a = 123; // read-write var
    println!("The value of a={}", a); // 12
    a = 456;

    const C: i32 = 123; // int 32bit
    let _b: u64 = 123; // unsigned int 64bit

    let x = 5;
    let x = x + 1; // Shadowing 重影
    let x = x * 2;
    println!("The value of a={},x={},C={}", a, x, C); // 12
    assert!(x == 12);

    let mut s = "123";
    assert!(s.len() == 3);
    s = "AAA";
    assert!(s == "AAA");

    /*
    位长度	有符号	无符号
    8-bit	i8	u8
    16-bit	i16	u16
    32-bit	i32	u32
    64-bit	i64	u64
    128-bit	i128	u128
    arch	isize	usize

    进制	例
    十进制	98_222
    十六进制	0xff
    八进制	0o77
    二进制	0b1111_0000
    字节(只能表示 u8 型)	b'A'
    */

    let _f1 = 2.0; // f64
    let _f2: f32 = 3.0; // f32

    let _sum = 5 + 10; // 加
    let _difference = 95.5 - 4.3; // 减
    let _product = 4 * 30; // 乘
    let _quotient = 56.7 / 32.2; // 除
    let _remainder = 43 % 5; // 求余

    let tup: (i32, f64, u8) = (100, 3.14, 1);
    assert!(tup.0 == 100);
    assert!(tup.1 == 3.14);
    assert!(tup.2 == 1);
    let (_x, y, _z) = tup;
    assert!(y == 3.14);

    let aa = [1, 2, 3, 4, 5];
    let _ab = ["January", "February", "March"];
    let _ac: [i32; 5] = [1, 2, 3, 4, 5];
    let _ad = [3; 5]; // 等同于 let d = [3, 3, 3, 3, 3];
    let first = aa[0];
    let second = aa[1];
    assert!(first == 1);
    assert!(second == 2);
    
    let mut aa = [1, 2, 3];
    aa[0] = 4;

    let ret = add(1, 2);
    assert!(ret == 3);

    let xx = 5;
    let yy = {
        let xx = 3;
        xx + 1
    };
    assert!(xx == 5);
    assert!(yy == 4);

    /// 诡异用法: 函数体
    /// 函数体不能使用return
    fn five() -> i32 {
        5 // 没有分号，它将变成一条语句！
    }
    println!("five() 的值为: {}", five());
    assert!(five() == 5);

    test_if();

    test_loop();

    test_mem();
}
