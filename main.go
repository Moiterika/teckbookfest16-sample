package main

import (
	"database/sql"
	"fmt"
	"os"
	"techbookfest16-sample/apl"

	_ "github.com/lib/pq"
)

//go:generate go run ./cmd/code_gen/ -i ./doc/テーブル定義書.xlsx
func main() {
	fmt.Println("開始")

	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOSTNAME"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	mhs := apl.NewMyHttpServer(db, 8080)
	mhs.EntryUsecaseWithNo("/単位/", mhs.UseCase単位)
	mhs.ListenAndServe()
	fmt.Println("終了")
}
