package main

// testingというパッケージを利用
import (
    "testing"
)

// 入力値（in）と期待値（out）を定義する構造体を作成（mapなどで代替しても良い）
type squareTest struct {
    in, out int
}

// テストをしたい入力値と期待値の一覧を作成
var squareTests = []squareTest {
    squareTest{1, 1},
    squareTest{2, 4},
    squareTest{5, 25},
    squareTest{-2, 4},
}

// テスト関数を作成する
// 関数名は「Test」をプレフィックスにつけ、引数には「t *testing.T」を渡す.
func TestSqure( t *testing.T) {

    // 入力値と期待値を1件ずつテストする.
    for _, st := range squareTests {
        v := Square(st.in)
        if v != st.out {
            // テストNGがあればエラーにする（エラー内容がわかるようにメッセージをちゃんと書く）
            t.Errorf("Square(%d) = %d, want %d.", st.in, v, st.out)
        }
    }
}
