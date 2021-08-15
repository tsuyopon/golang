# 概要
Go言語ではinterface{}型とinterface は異なるのでちゃんと理解すること

# 概要
Go言語のinterfaceの役割は以下の２つです。
- 1) どんな型の値でも入れておける入れ物(TypeScript の any型 や C言語の void型 のようなもの) や 型アサーションやswitch内部での型判定を利用する
  - interface{} 型として使われる
- 2) 構造体がメソッド(関数)を持つことを保証するための型の定義
  - type xxxx interfaceとして使われる

上記を確認するために以下のコードを確認してください。
- 1) sample1\_\*.go
- 2) sample2\_\*.go


