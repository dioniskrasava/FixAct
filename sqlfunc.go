package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTableInDB() {

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

func createDB() {
	var err error

	if pathFileDB == "" {
		// Если путь к базе данных не указан, используем путь по умолчанию
		db, err = sql.Open("sqlite3", "./activities.db")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Если путь к базе данных указан, используем его
		db, err = sql.Open("sqlite3", pathFileDB)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Проверяем подключение к базе данных
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

}

func addAct(w Widgets) {
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
