package main

import (
	"FixActApp/pos"
	"log"

	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func createInterfaceApp() (content *fyne.Container) {
	// Элементы интерфейса
	activityType := widget.NewSelect([]string{"Книга", "Код", "Видео"}, func(value string) {})
	activityType.PlaceHolder = "Выбери активность"

	startTime := widget.NewEntry()
	startTime.SetText(getNow())

	endTime := widget.NewEntry()
	totalTime := widget.NewEntry()
	comment := widget.NewMultiLineEntry()
	addButton := widget.NewButton("Добавить активность", func() { addAct(widgtsApp) })

	widgtsApp = Widgets{
		activityType: activityType,
		startTime:    startTime,
		endTime:      endTime,
		totalTime:    totalTime,
		comment:      comment,
		addButton:    addButton,
	}

	btnSupp1 := widget.NewButton("*", func() { startTime.SetText(getNow()) })
	btnSupp2 := widget.NewButton("*", func() { endTime.SetText(getNow()) })
	btnSupp3 := widget.NewButton("*", func() { totalTime.SetText(getActTime(endTime.Text, startTime.Text)) })

	globContainer := container.NewWithoutLayout()

	h1 := float32(10)
	h2 := float32(50)
	h3 := float32(90)
	h4 := float32(130)
	h5 := float32(170)
	h6 := float32(250)

	pos.AddRow(globContainer, WIDTH, h1, widget.NewLabel("Тип активности:"), activityType)
	pos.AddRow(globContainer, WIDTH, h2, widget.NewLabel("Время начала:"), startTime, btnSupp1)
	pos.AddRow(globContainer, WIDTH, h3, widget.NewLabel("Время окончания:"), endTime, btnSupp2)
	pos.AddRow(globContainer, WIDTH, h4, widget.NewLabel("Общее время:"), totalTime, btnSupp3)
	pos.AddRow(globContainer, WIDTH, h5, widget.NewLabel("Комментарий:"), comment)
	pos.AddRow(globContainer, WIDTH, h6, widget.NewLabel(""), addButton)

	return globContainer
}

func createMenu(w fyne.Window) *fyne.MainMenu {
	file_item1 := fyne.NewMenuItem("Сохранить", func() {})
	file_item2 := fyne.NewMenuItem("Открыть", func() {
		openFileWindow(w)
		// ВОЗМОЖНО ТУТ СТОИТ ОБДУМАТЬ !!!!
		// ПОТОМУ ЧТО КАЖЕТСЯ БД НЕ ОБНОВЛЯЕТСЯ ИЛИ ПРОСТО
		// ПРОГРАММА ЖДЕТ ПОКА ПОЛЬЗОВАТЕЛЬ ВЫБЕРЕТ ПУТЬ А ЭТОТ КОД УЖЕ ОТРАБОТАЛ
	})

	file_menu := fyne.NewMenu("Файл", file_item1, file_item2)

	main_menu := fyne.NewMainMenu(file_menu)

	return main_menu
}

func openFileWindow(w fyne.Window) {
	// Увеличиваем размер окна
	w.Resize(fyne.NewSize(550, 400)) // Установите желаемый размер окна

	// Показываем диалог открытия файла
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		// Восстанавливаем исходный размер окна после закрытия диалога
		defer w.Resize(fyne.NewSize(WIDTH, HEIGHT))

		if err != nil {
			log.Println("Error opening file:", err)
			return
		}
		if reader == nil {
			log.Println("No file selected")
			return
		}
		defer reader.Close()

		checkOpenFileDB(reader.URI().Path(), w)

		// Обрабатываем выбранный файл
		log.Println("Selected file:", reader.URI().Path())

	}, w)
}

func checkOpenFileDB(pathFileFromUser string, w fyne.Window) {
	// ЕСЛИ ВЫБРАННЫЙ ФАЙЛ КОНЧАЕТСЯ НА .ДБ
	if strings.HasSuffix(pathFileFromUser, ".db") {
		pathFileDB = pathFileFromUser // Обновляем глобальную переменную

		// ОТКЛЮЧАЕМСЯ ОТ СТАРОЙ БАЗЫ ДАННЫХ
		if db != nil {
			log.Println("the old database was disabled // отключили старую базу данных")
			db.Close()
		}

		// Обновляем подключение к базе данных
		createDB()
		if db == nil {
			log.Println("Failed to create or connect to the database // Не удалось создать базу данных или подключиться к ней")
			return
		}

		// Теперь можно использовать новую базу данных
		log.Println("Database connection updated successfully // Подключение к базе данных успешно обновлено")
		dialog.ShowInformation("Информация", "Всё заебок! Подключились к новой базе данных!", w)
	} else {
		//ТЫ ДУРАЧОК!
		dialog.ShowInformation("Предупреждение", "Ты дурачок! Выбрал не тот тип файла!", w)
	}
}
