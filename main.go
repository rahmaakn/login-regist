package main

import (
	"Final-Project-Engineering-50/Backend/api"
	"Final-Project-Engineering-50/Backend/repository"
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
)

func main(){
	db, err := sql.Open("sqlite3", "./beasiswa.db")
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	aimproveRepo := repository.NewAimproveRepository(db)
	mainApi := api.NewApi(*userRepo, *aimproveRepo)
	mainApi.Start() 
}