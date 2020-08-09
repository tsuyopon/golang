# 概要
GO111MODULE=onでモジュール分割を実行させるサンプル

以下の点に着目してください。
- go.modに"module testmod"が定義されていること
- main.goでは"import "testmod/hogefuga"として"testmod/hogefuga"が定義されていること

# 参考
- https://pod.hatenablog.com/entry/2018/12/26/074944
