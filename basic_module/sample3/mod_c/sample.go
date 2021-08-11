package mod_c

import (
    // fmtにfというエイリアスを付与することで、fmt.Printではなく、f.Printで呼び出すことができるようになる
    f "fmt"
)

func Test() string {
    f.Print("Hello world from mod_c!\n")
    return "This is mod_c Test."
}
