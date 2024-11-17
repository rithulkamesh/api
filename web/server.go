package web

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rithulkamesh/api/util"
	unkeygo "github.com/unkeyed/unkey-go"
)

func initializeAppState() (*AppState, error) {
	unkeyClient := unkeygo.New(
		unkeygo.WithSecurity(os.Getenv("UNKEY_ROOT_KEY")),
	)

	return &AppState{
		UnkeyClient: unkeyClient,
	}, nil
}

func InitServer(db *util.DB) {
	appState, err := initializeAppState()
	if err != nil {
		log.Fatalf("Failed to initialize app state: %v", err)
	}

	e := echo.New()
	registerPasteRoutes(e, db, appState, os.Getenv("BASE_URL"))

	e.Logger.Fatal(e.Start(":8080"))
}
