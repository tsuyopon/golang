// special thanks to https://qiita.com/rock619/items/14eb2b32f189514b5c3c
package main

import (
    "fmt"
    "net/http"
)

func main() {

    // 値のデフォルトのフォーマットでの表現を出力する。
    fmt.Printf("\n===================== %%v basic sample ====================\n");
    fmt.Printf("%v\n", true)
    fmt.Printf("%v\n", 42)
    fmt.Printf("%v\n", uint(42))
    fmt.Printf("%v\n", 12.345)
    fmt.Printf("%v\n", 1-2i)
    fmt.Printf("%v\n", "寿司🍣Beer🍺")
    fmt.Printf("%v\n", make(chan bool))
    fmt.Printf("%v\n", new(int))

    // 各要素に対して再帰的に %v でのフォーマットをしたものが結果として出力される。
    fmt.Printf("\n===================== %%v composit sample ====================\n");
    fmt.Printf("%v\n", http.Client{})
    fmt.Printf("%v\n", &http.Client{})
    fmt.Printf("%v\n", [...]int{1, 2, 3})
    fmt.Printf("%v\n", &[...]int{1, 2, 3})
    fmt.Printf("%v\n", []int{1, 2, 3})
    fmt.Printf("%v\n", &[]int{1, 2, 3})
    fmt.Printf("%v\n", map[string]int{"寿司": 1000, "ビール": 500})
    fmt.Printf("%v\n", &map[string]int{"寿司": 1000, "ビール": 500})

    // 構造体の場合にフィールド名を出力す
    fmt.Printf("\n===================== %%+v sample ====================\n");
    fmt.Printf("%v\n", http.Client{})
    fmt.Printf("%+v\n", http.Client{})

    // 値のGoの文法での表現を出力する。
    fmt.Printf("\n===================== %%#v sample ====================\n");
    fmt.Printf("%#v\n", true)
    fmt.Printf("%#v\n", 42)
    fmt.Printf("%#v\n", uint(42))
    fmt.Printf("%#v\n", 12.345)
    fmt.Printf("%#v\n", 1-2i)
    fmt.Printf("%#v\n", "寿司🍣Beer🍺")
    fmt.Printf("%#v\n", make(chan bool))
    fmt.Printf("%#v\n", new(int))
    fmt.Printf("\n")
    fmt.Printf("%#v\n", http.Client{})
    fmt.Printf("%#v\n", &http.Client{})
    fmt.Printf("%#v\n", [...]int{1, 2, 3})
    fmt.Printf("%#v\n", &[...]int{1, 2, 3})
    fmt.Printf("%#v\n", []int{1, 2, 3})
    fmt.Printf("%#v\n", &[]int{1, 2, 3})
    fmt.Printf("%#v\n", map[string]int{"寿司": 1000, "ビール": 500})
    fmt.Printf("%#v\n", &map[string]int{"寿司": 1000, "ビール": 500})

    // 値の型のGoの文法での表現を出力する。型情報のみを表示する
    fmt.Printf("\n===================== %%#v sample ====================\n");
    fmt.Printf("%T\n", true)
    fmt.Printf("%T\n", 42)
    fmt.Printf("%T\n", uint(42))
    fmt.Printf("%T\n", 12.345)
    fmt.Printf("%T\n", 1-2i)
    fmt.Printf("%T\n", "寿司🍣Beer🍺")
    fmt.Printf("%T\n", make(chan bool))
    fmt.Printf("%T\n", new(int))
    fmt.Printf("\n")
    fmt.Printf("%T\n", http.Client{})
    fmt.Printf("%T\n", &http.Client{})
    fmt.Printf("%T\n", [...]int{1, 2, 3})
    fmt.Printf("%T\n", &[...]int{1, 2, 3})
    fmt.Printf("%T\n", []int{1, 2, 3})
    fmt.Printf("%T\n", &[]int{1, 2, 3})
    fmt.Printf("%T\n", map[string]int{"寿司": 1000, "ビール": 500})
    fmt.Printf("%T\n", &map[string]int{"寿司": 1000, "ビール": 500})

    // %自体を表示したい場合に利用する
    fmt.Printf("\n===================== %% sample ====================\n");
    fmt.Printf("100%%\n")

    // trueかfalseをそのまま文字列として表示したい場合に利用する
    fmt.Printf("\n===================== %% sample ====================\n");
    fmt.Printf("%t\n", true)
    fmt.Printf("%t\n", false)

    // 整数表現
    fmt.Printf("\n===================== %%d, %%b, %%o, %%x, %%X, %%c, %%q  ====================\n");
    answer := 42
    fmt.Printf("%b\n", answer)  // 2進数表現
    fmt.Printf("%c\n", answer)  // Unicodeコードポイントに対応する文字
    fmt.Printf("%d\n", answer)  // 10進数表現
    fmt.Printf("%o\n", answer)  // 8進数表現
    fmt.Printf("%q\n", answer)  // 対応する文字をシングルクォートで囲んだ文字列
    fmt.Printf("%x\n", answer)  // 16進数での表現(a-fは小文字)
    fmt.Printf("%X\n", answer)  // 16進数での表現(A-Fは大文字)
    fmt.Printf("%U\n", answer)  // 

    // 浮動小数点、複素数
    fmt.Printf("\n===================== %%b, %%e, %%E, %%f, %%F, %%g, %%G  ====================\n");
    f := 12.345
    fmt.Printf("%b\n", f)   // 小数点なしの指数表記 指数は2の累乗
    fmt.Printf("%e\n", f)   // 指数表記
    fmt.Printf("%E\n", f)   // %eのeがEで表記される
    fmt.Printf("%f\n", f)   // 指数表記なし
    fmt.Printf("%F\n", f)   // 指数表記なし
    fmt.Printf("%g\n", f)   // 指数が大きい場合は%eそうでなければ%f
    fmt.Printf("%G\n", f)   // 指数が大きい場合は%Eそうでなければ%F
    fmt.Printf("%g\n", 12345678.9)  // 指数が大きい場合は%eそうでなければ%f
    fmt.Printf("%G\n", 12345678.9)  // 指数が大きい場合は%Eそうでなければ%F

    // 浮動小数点、複素数
    fmt.Printf("\n===================== %%s, %%q, %%x, %%X ====================\n");
    s := "寿司🍣Beer🍺"
    fmt.Printf("%s\n", s)
    fmt.Printf("%q\n", s)
    fmt.Printf("%x\n", s)
    fmt.Printf("%X\n", s)

    // 幅調整
    fmt.Printf("\n===================== width, precision ====================\n");
    f2 := 12.345
    fmt.Printf("%f\n", f2)
    fmt.Printf("%12f\n", f2)
    fmt.Printf("%12.2f\n", f2)
    fmt.Printf("%.2f\n", f2)
    fmt.Printf("%12.f\n", f2)
    fmt.Printf("%e\n", f2)
    fmt.Printf("%#g\n", f2)
    fmt.Printf("%g\n", f2)
    fmt.Printf("%f", 1-2i)

    // flag
    fmt.Printf("\n===================== + ====================\n");
    fmt.Printf("%d\n", 42)
    fmt.Printf("%+d\n", 42)  // 符号+を表示
    fmt.Printf("%q\n", 945)  // ASCIIだけで出力する
    fmt.Printf("%+q\n", 945)
    fmt.Printf("%q\n", "寿司🍣Beer🍺")
    fmt.Printf("%+q\n", "寿司🍣Beer🍺")

    // 左詰
    fmt.Printf("\n===================== left aligned ====================\n");
    fmt.Printf("%5d\n", 42)
    fmt.Printf("%-5d\n", 42)
    fmt.Printf("%10s\n", "寿司🍣Beer🍺")
    fmt.Printf("%-10s\n", "寿司🍣Beer🍺")

    // # デフォルトとは異なるフォーマットにする
    fmt.Printf("\n===================== not default format ====================\n");
    fmt.Printf("%o\n", answer)
    fmt.Printf("%#o\n", answer)  // 8進数の場合(%#o): 先頭に0を付ける
    fmt.Printf("%x\n", answer)
    fmt.Printf("%#x\n", answer)  // 16進数の場合(%#x): 先頭に0xを付ける
    fmt.Printf("%X\n", answer)
    fmt.Printf("%#X\n", answer)  // 16進数(大文字)の場合(%#X): 先頭に0Xを付ける
    fmt.Printf("%p\n", &answer)
    fmt.Printf("%#p\n", &answer) // ポインタの場合(%#p): 先頭の0xを付けない
    fmt.Printf("%q\n", "go")
    fmt.Printf("%#q\n", "go")    // %qの場合: strconv.CanBackquoteがtrueを返すならraw文字列を出力する
    fmt.Printf("%q\n", "`go`")
    fmt.Printf("%#q\n", "`go`")  // %qの場合: strconv.CanBackquoteがtrueを返すならraw文字列を出力する
    fmt.Printf("%.f\n", 12.345)
    fmt.Printf("%#.f\n", 12.345) // %e, %E, %f, %F, %g, %Gの場合: 必ず小数点を付ける
    fmt.Printf("%g\n", 12.345)
    fmt.Printf("%#g\n", 12.345)  // %g, %Gの場合: 末尾の0を省略しない
    fmt.Printf("%U\n", answer)
    fmt.Printf("%#U\n", answer)  // %Uの場合: U+0078 'x'の形式で出力する

    // 0埋め
    fmt.Printf("\n===================== 0 padding ====================\n");
    fmt.Printf("%10s\n", "寿司🍣Beer🍺")
    fmt.Printf("%010s\n", "寿司🍣Beer🍺")
    fmt.Printf("%10.3f\n", -12.345)
    fmt.Printf("%010.3f\n", -12.345)

}
