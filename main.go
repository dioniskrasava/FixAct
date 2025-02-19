package main

import (
	"FixActApp/pos"
	"database/sql"
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

const (
	WIDTH  float32 = 400
	HEIGHT float32 = 300
)

// newApp
type Activity struct {
	Type      string
	StartTime string
	EndTime   string
	TotalTime string
	Comment   string
}

// структура виджетов микроприложения
type Widgets struct {
	activityType *widget.Select
	startTime    *widget.Entry
	endTime      *widget.Entry
	totalTime    *widget.Entry
	comment      *widget.Entry
	addButton    *widget.Button
}

var widgtsApp Widgets = Widgets{}

func main() {
	newApp()
}

func newApp() {
	a := app.New()
	w := a.NewWindow("FixAct")
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.SetFixedSize(true)

	db := createDB()
	defer db.Close()

	createTableInDB(db)
	content := createInterfaceApp(db)

	// Установка контента в окно
	w.SetContent(content)
	w.ShowAndRun()
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
	db, err := sql.Open("sqlite3", "./activities.db")
	if err != nil {
		log.Fatal(err)
	}

	return db

}

func createInterfaceApp(db *sql.DB) (content *fyne.Container) {
	// Элементы интерфейса
	activityType := widget.NewSelect([]string{"Книга", "Код", "Видео"}, func(value string) {})
	activityType.PlaceHolder = "Выбери активность"

	startTime := widget.NewEntry()
	h, m, s := time.Now().Clock()
	timeStr := fmt.Sprintf("%d:%d:%d", h, m, s)
	startTime.SetText(timeStr)

	endTime := widget.NewEntry()
	totalTime := widget.NewEntry()
	comment := widget.NewMultiLineEntry()
	addButton := widget.NewButton("Добавить активность", func() { addAct(widgtsApp, db) })

	widgtsApp = Widgets{
		activityType: activityType,
		startTime:    startTime,
		endTime:      endTime,
		totalTime:    totalTime,
		comment:      comment,
		addButton:    addButton,
	}

	btnSupp1 := widget.NewButton("!", func() {})

	globContainer := container.NewWithoutLayout()

	h1 := float32(10)
	h2 := float32(50)
	h3 := float32(90)
	h4 := float32(130)
	h5 := float32(170)
	h6 := float32(250)

	pos.AddRow(globContainer, WIDTH, h1, widget.NewLabel("Тип активности:"), activityType)
	pos.AddRow(globContainer, WIDTH, h2, widget.NewLabel("Время начала:"), startTime, btnSupp1)
	pos.AddRow(globContainer, WIDTH, h3, widget.NewLabel("Время окончания:"), endTime)
	pos.AddRow(globContainer, WIDTH, h4, widget.NewLabel("Общее время:"), totalTime)
	pos.AddRow(globContainer, WIDTH, h5, widget.NewLabel("Комментарий:"), comment)
	pos.AddRow(globContainer, WIDTH, h6, widget.NewLabel(""), addButton)

	return globContainer
}
