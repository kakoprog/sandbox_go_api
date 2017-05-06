# sandbox_go_api
sandbox go api

#### 実行時の注意
カレントディレクトリにconfig.yamlがある状態で実行ファイルや go run を実行する
dbはsqliteを使っていて、配置はconfig.yamlに実行時のディレクトリからの相対パスで指定

#### db周りはxoで作成 + 手直し
```
xo 'file:sqlite/todo.db?loc=auto' -o repository/
```
