# cara menjalankan
## 1. pastikan file di folder config -> evn.go diisi dengan nama root folder untuk meload file .env
```go
// for more information, see --> https://github.com/joho/godotenv/issues/43#issuecomment-503183127
const projectDirName = "test-golang" // ini adalah root directori project
```
## 2. isi file .env dan sesuaikan dengan database yang diinginkan
## 3. pada terminal directory saat ini, ketikan
```bash
Make run
```
## atau dapat langsung menjalankan progam dengan
```bash
go run main.go
```
## 4. Jika sesuai, maka akan muncul
```bash
2023/06/30 23:50:06 Postgres server listening on port 5432...
2023/06/30 23:50:06 migration success
2023/06/30 23:50:06 your jwt for auth : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoicmlraWFuZmFpc2FsIiwiZXhwIjoxNjkxNzQzODA2fQ.DM2MyWLl1P4tl09BNJxMm1x4Whge9Y4BYKkmoA6fJts
2023/06/30 23:50:06 Server running at 9090
```
### Success create connection to Postgres server at localhost://postgres:****:2345/your_db -> berhasil connect
### migration success -> migrasi database book dengan default 5 data, setiap restart akan migrasi, data yang tersimpan akan terganti default.
### your jwt for auth bearer -> pastikan masukkan kode ini di header untuk setiap request. Expired 1 tahun