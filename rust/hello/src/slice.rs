pub fn slice1() {
    let s = String::from("slice");

    /*
    ..y 等价于 0..y
    x.. 等价于位置 x 到数据结束
    .. 等价于位置 0 到结束
     */
    let part1 = &s[0..3]; // read-only
    let part2 = &s[2..5];
    assert!(part1 == "sli");
    assert!(part2 == "ice");

    println!("slice/{}={}+{}", s, part1, part2);
}