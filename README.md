# gin による Web アプリケーション構築のチュートリアル

## 参考資料
https://qiita.com/advent-calendar/2018/gopherdojo

## 学んだこと

### gin による簡単なAPIの作成

localhost：8080/ にアクセスすると Hello World を表示するアプリケーションは以下の通り。

```go
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})
	r.Run()
}
```

### postgreSQL に接続

```
$ psql -h 0.0.0.0 -U user gorm -p 5433
```
`gorm`はデータベース名を指す。