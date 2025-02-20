package pos

import (
	"fyne.io/fyne/v2"
)

// название классное
// принимает глобальный кастомный контейнер container.NewWithoutLayout(), ширину окна, высоту позиции строки виджетов (пока еще недоработан этот момент)
// и пачку (ВРОДЕ 3 ОБЪЕКТА. ИНАЧЕ ОШИБКА) канвас объектов которые нужно расстрочить
func AddRow(globContainer *fyne.Container, widthWindow float32, heightWidg float32, objects ...fyne.CanvasObject) {

	// ratios - процентные соотношения, на которые делится окно
	ratios := []float32{0, 37, 91}

	// высчитываем 1 процент
	percent := widthWindow / 100

	// получаем позиции вертик линий
	positionsX := getVertLinePos(ratios, percent)
	// получаем конкретные координаты вертикальных линий по Х
	widthsX := getColumnWidths(positionsX, widthWindow)
	// стругаем контейнер с уже позиционированными виджетами
	addGlobContainer(globContainer, positionsX, widthsX, heightWidg, objects...)

}

// получить позиции вертикальных границ
// вспомогательная функция для awesomeShit
func getVertLinePos(ratios []float32, percent float32) []float32 {
	positionsX := []float32{} // 0, 250, 400
	// вычисляем позиции вертикальных рамок
	for _, v := range ratios {
		posX := percent * v
		positionsX = append(positionsX, posX)
	}

	//fmt.Println("Позиции вертикальных линий : ", positionsX)
	return positionsX
}

// получить позиции по оси Х
// вспомогательная функция для awesomeShit
func getColumnWidths(positionsX []float32, w float32) []float32 {
	widthsX := []float32{}
	// вычисляем ширины колонок
	for k, _ := range positionsX {
		// если последний элемент
		if k == len(positionsX)-1 {
			width := w - positionsX[k]
			widthsX = append(widthsX, width)
		} else {
			width := positionsX[k+1] - positionsX[k]
			widthsX = append(widthsX, width)
		}
	}
	//fmt.Println("Ширины колонок : ", widthsX)
	return widthsX
}

// добавление объектов на кастомную сетку
// вспомогательная функция для awesomeShit
func addGlobContainer(globContainer *fyne.Container, positionsX []float32, widthsX []float32, heightWidg float32, objects ...fyne.CanvasObject) {

	for k, v := range objects {
		v.Move(fyne.NewPos(positionsX[k], heightWidg))
		v.Resize(fyne.NewSize(widthsX[k]-5, v.MinSize().Height))
		globContainer.Add(v)
	}
}
