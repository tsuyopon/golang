# 概要
POSTリクエストの使い分け

以下の3つの記述方法があります

- http.PostForm(url, data)
  - http.PostForm関数は、http.Post関数と比較し、より単純化された機能です
- http.Post(method, url, data)
  - タイムアウト値やクッキー値などを指定したい場合は、こちらを使用するとよい
- http.NewRequestを使う
  - 細かな制御を行いたい場合には

# 詳細

