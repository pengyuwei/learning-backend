pub fn test_num() {
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

    let _i32 = 18_000; // 18,000
    let i: i64 = 18_000; // 18,000
    let _f1 = 2.0; // f64
    let _f2: f32 = 3.0; // f32

    let _hex = 0xff;
    let _octal = 0o77;
    let _binary = 0b1111_0000;
    let _decimal = 98_222;

    println!("testNum:{0},{1}", _i32, i);
}