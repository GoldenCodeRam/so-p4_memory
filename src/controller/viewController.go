package controller

import (
	"so-p4_memory/src/model"
	"so-p4_memory/src/view"
	"sync"

	"github.com/gotk3/gotk3/gtk"
)

type viewController struct {
	MainWindow *view.MainWindow
	Processor  *model.Processor
	Storage    *model.Storage
}

var viewControllerInstance *viewController
var lock = &sync.Mutex{}

func GetViewControllerInstance() *viewController {
	if viewControllerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if viewControllerInstance == nil {
			viewControllerInstance = &viewController{}
			viewControllerInstance.Processor = &model.Processor{
				LogListeners: viewControllerInstance,
			}
		}
	}
	return viewControllerInstance
}

func (v *viewController) StartApplication() {
	gtk.Init(nil)

	v.MainWindow = view.CreateMainWindow(v)
	v.MainWindow.Window.SetDefaultSize(800, 600)
	v.MainWindow.Window.ShowAll()

	gtk.Main()
}
