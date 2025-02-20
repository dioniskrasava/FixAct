package main

// ФАЙЛ С ФУНКЦИЯМИ ДЛЯ РАБОТЫ С ВРЕМЕНЕМ

import (
	"fmt"
	"time"
)

// ф-ии работы со временем определить в другой файл!
func getNow() string {
	h, m, s := time.Now().Clock()
	timeStr := fmt.Sprintf("%02d:%02d:%02d", h, m, s) // Добавляем ведущие нули
	return timeStr
}

// ф-ии работы со временем определить в другой файл!
func getActTime(time1 string, time2 string) string {
	// Парсинг строк в объекты time.Time
	layout := "15:04:05"
	t1, err := time.Parse(layout, time1)
	if err != nil {
		fmt.Println("Ошибка парсинга времени:", err)
	}

	t2, err := time.Parse(layout, time2)
	if err != nil {
		fmt.Println("Ошибка парсинга времени:", err)
	}

	// Вычисление разницы между временами
	diff := t1.Sub(t2)

	// Вывод результата
	fmt.Printf("Разница между %s и %s составляет %v\n", time1, time2, diff)

	return formatDuration(diff)
}

// ф-ии работы со временем определить в другой файл!

// Функция для преобразования time.Duration в формат 00:00:00
func formatDuration(d time.Duration) string {
	// Получаем общее количество секунд
	totalSeconds := int(d.Seconds())

	// Вычисляем часы, минуты и секунды
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	// Форматируем с ведущими нулями
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
