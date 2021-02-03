package main

import (
	"fyne.io/fyne/v2/dialog"
	"github.com/getbuguai/gaihosts"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var topWindow fyne.Window

func main() {
	a := app.NewWithID("buguai.get.")
	a.SetIcon(pngIconResource)
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("改 Hosts")

	//systray.Run(tutorials.OnReady, func() {
	//	a.Quit()
	//})

	topWindow = w

	w.SetContent(makeListTab(w))

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func makeListTab(w fyne.Window) fyne.CanvasObject {
	filesName, mapFile, err := gaihosts.GetConfigFilesName()
	if err != nil {
		panic(err)
	}

	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.Wrapping = fyne.TextWrapWord

	list := widget.NewList(
		func() int {
			return len(filesName)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()),
				widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(filesName[id])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		_, s, err := gaihosts.GetFileBody(mapFile[filesName[id]])
		if err != nil {
			dialog.ShowInformation("提示", "打开文件出现错误", w)
		} else {
			entryLoremIpsum.SetText(s)
		}
		//entryLoremIpsum.SetText("-------------" + string(id))
	}
	list.OnUnselected = func(id widget.ListItemID) {
		entryLoremIpsum.SetText("选择一个进行查看")
	}
	list.Select(1)

	return container.NewHSplit(list, entryLoremIpsum)
}
