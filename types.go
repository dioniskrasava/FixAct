package main

import "fyne.io/fyne/v2/widget"

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
