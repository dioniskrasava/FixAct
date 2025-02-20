package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTableInDB(db *sql.DB) {

	// Создание таблицы, если она не существует
	createTableSQL := `CREATE TABLE IF NOT EXISTS activities (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	type TEXT,
	start_time TEXT,
	end_time TEXT,
	total_time TEXT,
	comment TEXT
);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func createDB() (db *sql.DB) {
	// Подключение к базе данных SQLite
	// если пути к базе данных нет, то создаём базу данных в той же директории

	var err error

	if pathFileDB == "" {
		db, err = sql.Open("sqlite3", "./activities.db")
		if err != nil {
			log.Fatal(err)
		}

		// если пользователь выбрал путь базы данных, то создаем по этому пути
	} else if pathFileDB != "" {
		db, err = sql.Open("sqlite3", pathFileDB)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db

}

func addAct(w Widgets, db *sql.DB) {
	activity := Activity{
		Type:      w.activityType.Selected,
		StartTime: w.startTime.Text,
		EndTime:   w.endTime.Text,
		TotalTime: w.totalTime.Text,
		Comment:   w.comment.Text,
	}

	// Вставка данных в базу данных
	insertSQL := `INSERT INTO activities (type, start_time, end_time, total_time, comment) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(insertSQL, activity.Type, activity.StartTime, activity.EndTime, activity.TotalTime, activity.Comment)
	if err != nil {
		log.Fatal(err)
	}

	// Очистка полей после добавления
	w.activityType.SetSelected("")
	w.startTime.SetText("")
	w.endTime.SetText("")
	w.totalTime.SetText("")
	w.comment.SetText("")

	fmt.Println("Активность добавлена!")
}
