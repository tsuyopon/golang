package main

import (
    "testing"
)

/*
 * ### ソースコードの処理詳細
 * trace_test.goは以下の処理で処理されます。
 *   1. Test_Func1が実行され、処理が完了します。
 *   2. 次に、Test_Func2の実行に移りますが、t.Parallel()メソッドを呼び出したところで一時停止します。
 *   3. Test_Func2の実行が停止した状態で、Test_Func3が実行され、処理が完了します。
 *   4. 次に、Test_Func4の実行に移りますが、t.Parallel()メソッドが呼び出されたとことで一時停止します。
 *   5. Test_Fun4の実行が停止した状態で、Test_Func5が実行され、処理が完了します。
 *   6. t.Parallel()メソッドを呼び出していないTest_Func1、Test_Func3、Test_Func5が順にすべて実行されると、t.Parallel()メソッドを呼び出しているTest_Func2とTest_Func4の処理が並列に再開して、処理が完了します。
 */
func Test_Func1(t *testing.T) {
    defer trace("Test_Func1")()

    // ...
}

func Test_Func2(t *testing.T) {
    defer trace("Test_Func2")()
    t.Parallel()

    // ...
}

func Test_Func3(t *testing.T) {
    defer trace("Test_Func3")()

    // ...
}

func Test_Func4(t *testing.T) {
    defer trace("Test_Func4")()
    t.Parallel()

    // ...
}

func Test_Func5(t *testing.T) {
    defer trace("Test_Func5")()

    // ...
}

