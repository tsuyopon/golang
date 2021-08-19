/*
 * Panicによるサンプル
 */
package main

import (
    "log"
)

func main(){
  div(2, 0)
}

func div(x, y float64) float64 {
  if y == 0 {
    log.Panicf("0での除算： %1.f/%1.f", x, y)
  }
  return x/y
}
