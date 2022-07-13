package utils

import (
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

func ShowErrorDialog(err error) error {
	dialog := gtk.MessageDialogNew(
		nil,
		gtk.DIALOG_DESTROY_WITH_PARENT,
		gtk.MESSAGE_ERROR,
		gtk.BUTTONS_NONE,
		err.Error(),
	)
	addDialogClosebutton(dialog)
	dialog.Run()
	dialog.Destroy()
	return err
}

func ShowInfoDialog(message string) {
	dialog := gtk.MessageDialogNew(
		nil,
		gtk.DIALOG_DESTROY_WITH_PARENT,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_NONE,
		message,
	)
	addDialogClosebutton(dialog)
	dialog.Run()
	dialog.Destroy()
}

func addDialogClosebutton(dialog *gtk.MessageDialog) {
	button, err := dialog.AddButton(lang.CLOSE, 0)
	if err != nil {
		panic(err)
	}

	button.Connect("clicked", func() {
		dialog.Close()
	})

}
