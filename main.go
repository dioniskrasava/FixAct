package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "github.com/mattn/go-sqlite3"
)

// ВЕРСИЯ  V 1.1.001 - ДОБАВИТЬ ФИЧУ (ВЫБОР МЕСТА СОХРАНИЕНИЯ БАЗЫ ДАННЫХ)

func main() {
	newApp()
}

func newApp() {
	// Создаем новое приложение с уникальным идентификатором
	a := app.NewWithID("com.fixact.app1")
	w := a.NewWindow("FixAct")
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.SetFixedSize(true)

	// Получаем объект Preferences
	//prefs := a.Preferences()
	/*
			Preferences предоставляет методы для работы с различными типами данных:
		1. Строки

		    SetString(key string, value string): Сохраняет строку по ключу.
		    String(key string) string: Возвращает строку по ключу. Если ключ отсутствует, возвращает пустую строку.
		    StringWithFallback(key string, fallback string) string: Возвращает строку по ключу или значение fallback, если ключ отсутствует.

			6. Удаление данных

		    RemoveValue(key string): Удаляет значение по ключу.
	*/

	createDB()
	defer db.Close()

	createTableInDB(db)
	content := createInterfaceApp(db)

	// Установка контента в окно
	w.SetContent(content)
	w.SetMainMenu(createMenu(w))
	w.ShowAndRun()
}
