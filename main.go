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

// oldApp const
const (
	WIDTH  float32 = 300
	HEIGHT float32 = 250
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
	//oldApp()
	newApp()
}

func newApp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("FixAct")

	db := createDB()
	defer db.Close()

	createTableInDB(db)
	content := createInterfaceApp(db)

	// Установка контента в окно
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
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

	row1 := container.NewGridWithColumns(2, widget.NewLabel("Тип активности:"), activityType)
	row2 := container.NewGridWithColumns(2, widget.NewLabel("Время начала:"), startTime)
	row3 := container.NewGridWithColumns(2, widget.NewLabel("Время окончания:"), endTime)
	row4 := container.NewGridWithColumns(2, widget.NewLabel("Общее время:"), totalTime)
	row5 := container.NewGridWithColumns(2, widget.NewLabel("Комментарий:"), comment)

	// Создание контейнера с элементами интерфейса
	cont1 := container.NewGridWithRows(4,
		row1,
		row2,
		row3,
		row4,
	)

	contSuppBut := container.NewGridWithRows(6, widget.NewLabel(""), btnSupp1)

	contentLeft := container.NewVBox(cont1, row5, addButton)
	content = container.NewHBox(contentLeft, contSuppBut)

	return content
}

// -------------------------------------------
// ---- OLD APP --------
// ------------------------
// -----------------------
func oldApp() {
	a := app.New()
	w := a.NewWindow("Fix activity application")
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.SetFixedSize(true)

	w.SetContent(oldWidgetsForTest())
	w.ShowAndRun()
}
func oldWidgetsForTest() *fyne.Container {
	lbl := widget.NewLabel("Begin time :")
	lbl.TextStyle = fyne.TextStyle{Bold: true}

	ent := widget.NewEntry()
	ent.Resize(fyne.NewSize(ent.MinSize().Width, ent.MinSize().Height))

	btn := widget.NewButton("Click", func() {})
	btn.Resize(fyne.NewSize(btn.MinSize().Width, btn.MinSize().Height))

	//----------------------------------------
	lbl2 := widget.NewLabel("end :")
	lbl2.TextStyle = fyne.TextStyle{Bold: true}

	ent2 := widget.NewEntry()

	btn2 := widget.NewButton("!", func() {})

	// высоты строк
	h1 := float32(10)
	h2 := float32(50)

	globContainer := container.NewWithoutLayout()

	pos.AwesomeShit(globContainer, WIDTH, h1, lbl, ent, btn)
	pos.AwesomeShit(globContainer, WIDTH, h2, lbl2, ent2, btn2)
	return globContainer
}
