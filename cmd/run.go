package cmd

import (
	"log"

	"github.com/ayn2op/discordo/internal/config"
	"github.com/ayn2op/discordo/internal/logger"
	"github.com/ayn2op/discordo/internal/ui"
	"github.com/rivo/tview"
)

var (
	discordState *State

	cfg      *config.Config
	app      = tview.NewApplication()
	mainFlex *MainFlex
  popup    *tview.Modal
  pages    *tview.Pages
)

func Run(token string) error {
	var err error
	cfg, err = config.Load()
	if err != nil {
		return err
	}

	if err := logger.Load(); err != nil {
		return err
	}

	if token == "" {
		lf := ui.NewLoginForm(cfg)

		go func() {
			// mainFlex must be initialized before opening a new state.
			mainFlex = newMainFlex()
      pages = tview.NewPages().AddPage("mainFlex", mainFlex, true, true)

			token := <-lf.Token
			if token.Error != nil {
				app.Stop()
				log.Fatal(token.Error)
			}

			if err := openState(token.Value); err != nil {
				app.Stop()
				log.Fatal(err)
			}

			app.QueueUpdateDraw(func() {
				app.SetRoot(pages, true)
			})
		}()

		app.SetRoot(lf, true)
	} else {
		mainFlex = newMainFlex()
    pages = tview.NewPages().AddPage("mainFlex", mainFlex, true, true)
		if err := openState(token); err != nil {
			return err
		}

		app.SetRoot(pages, true)
	}

	app.EnableMouse(cfg.Mouse)
	return app.Run()
}
