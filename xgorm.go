package main

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)
func main(){
db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: "host=localhost user=cahyono password=csteam512@ dbname=aksi port=5432 sslmode=disable TimeZone=Asia/Jakarta", // data source name, refer https://github.com/jackc/pgx
  PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
}), &gorm.Config{})

if err != nil {
panic("gagal kin3k dul")
}

fmt.Println("koneksi hasil")
asak, _ := db.DB()
asak.Stats()
}
