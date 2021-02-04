package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/layout"

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

func makeToolbarTab(w fyne.Window) fyne.CanvasObject {
	t := widget.NewToolbar(
		widget.NewToolbarAction(theme.MailComposeIcon(), func() { fmt.Println("New") }),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() { fmt.Println("Cut") }),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() { fmt.Println("Copy") }),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() { fmt.Println("Paste") }),
	)

	return container.NewBorder(t, nil, nil, nil)
}

func makeListTab(w fyne.Window) fyne.CanvasObject {
	filesName, mapFile, err := gaihosts.GetConfigFilesName()
	if err != nil {
		panic(err)
	}

	entryLoremIpsum := widget.NewMultiLineEntry()
	entryLoremIpsum.Wrapping = fyne.TextWrapWord

	prev := widget.NewButton("保存", func() {
		fmt.Println("保存")
	})
	next := widget.NewButton("刷新", func() {
		fmt.Println("刷新")
	})
	newFile := widget.NewButton("新建", func() {
		fmt.Println("新建")
	})

	usingStatus := "启用"
	startUsing := widget.NewButton(usingStatus, func() {
		fmt.Println(usingStatus)
		usingStatus = "停用"
	})

	clearHostsCache := widget.NewButton("清除缓存", func() {
		fmt.Println("清除缓存")
	})

	buttons := container.NewHBox(prev, next, newFile, startUsing, clearHostsCache)
	bar := container.NewBorder(nil, nil, buttons, nil)

	list := widget.NewList(
		func() int {
			return len(filesName)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()),
				widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			name := filesName[id]
			item.(*fyne.Container).
				Objects[1].(*widget.Label).
				SetText(name[:strings.LastIndex(name, ".")])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		_, s, err := gaihosts.GetFileBody(mapFile[filesName[id]])
		if err != nil {
			dialog.ShowInformation("提示", "打开文件出现错误", w)
		} else {
			entryLoremIpsum.SetText(s)
			entryLoremIpsum.Disable()
		}

	}
	list.OnUnselected = func(id widget.ListItemID) {
		entryLoremIpsum.SetText("选择一个进行查看")
	}
	list.Select(1)

	return container.NewHSplit(list, container.New(layout.NewBorderLayout(
		bar, nil, nil, nil), bar, entryLoremIpsum))
}
