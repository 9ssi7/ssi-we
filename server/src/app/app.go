package app

type App struct {}

func New() *App {
	return &App{}
}

func (a *App) Prepare() {}

func (a *App) Listen() {}