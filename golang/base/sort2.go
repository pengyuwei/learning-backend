package main

import (
	"fmt"
	"sort"
)

type Data [][]string

func (p Data) Len() int {
    return len(p)
}

func (p Data) Less(i, j int) bool {
    // 以第五个元素为排序依据
    // 正序：<
    // 倒序：>
    return p[i][5] < p[j][5]
}

func (p Data) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func main() {
    var d Data
    d = [][]string{
        {"a", "a1", "a2", "a3", "a4", "5"},
        {"b", "b1", "b2", "b3", "b4", "4"},
        {"c", "c1", "c2", "c3", "c4", "3"},
        {"d", "d1", "d2", "d3", "d4", "2"},
        {"e", "e1", "e2", "e3", "e4", "1"},
    }

    // Slice方法不需要Len、Less、Swap三个方法
    sort.Slice(d, func(i, j int) bool {
        // 正序：>
        return d[i][5] < d[j][5]
    })
    fmt.Println("正序：", d)
    sort.Slice(d, func(i, j int) bool {
        // 倒序：>
        return d[i][5] > d[j][5]
    })
    fmt.Println("倒序：", d)

    // 依赖Len、Less、Swap三个方法的实现
    sort.Sort(d)
    fmt.Println("正序：", d)
}
