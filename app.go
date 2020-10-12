package glap

// App TODO
type App struct {
	// Name TODO
	name string

	// Author TODO
	author string

	// About TODO
	about string

	// LongAbout TODO
	longAbout string

	// AfterHelp TODO
	afterHelp string

	// BeforeHelp TODO
	beforeHelp string

	// Version TODO
	version string

	// LongVersion TODO
	longVersion string

	// Usage TODO
	usage string

	// Help TODO
	help string
}

// NewApp TODO
func NewApp() *App {
	return &App{}
}

// Name TODO
func (a *App) Name(name string) *App {
	a.name = name
	return a
}

// Author TODO
func (a *App) Author(author string) *App {
	a.author = author
	return a
}
