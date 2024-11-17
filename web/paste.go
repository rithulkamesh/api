package web

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rithulkamesh/api/util"
	"github.com/teris-io/shortid"
)

type PasteResponse struct {
	ID        string    `json:"id"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreatePasteRequest struct {
	Content string `json:"content"`
}

func registerPasteRoutes(router *echo.Echo, db *util.DB, appState *AppState, baseURL string) error {
	r := router.Group("/api/paste")
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)

	if err != nil {
		return err
	}

	r.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid paste ID")
		}

		paste, err := db.GetOne(id)
		if err != nil {
			if err.Error() == "paste not found" {
				return echo.NewHTTPError(http.StatusNotFound, "Paste not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve paste")
		}

		c.Response().Header().Set("Content-Type", http.DetectContentType(paste.Content))

		return c.Blob(http.StatusOK, "", paste.Content)
	})

	r.Use(unkeyMiddleware(appState))

	r.POST("/", func(c echo.Context) error {

		var req CreatePasteRequest
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if req.Content == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		content, err := base64.StdEncoding.DecodeString(req.Content)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Content must be base64 encoded")
		}

		id, err := sid.Generate()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create paste")

		}

		paste := &util.Paste{
			ID:        id,
			Content:   content,
			CreatedAt: time.Now(),
		}

		if err := db.Create(paste); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create paste")
		}

		response := PasteResponse{
			ID:        paste.ID,
			Link:      baseURL + "/api/paste/" + paste.ID,
			CreatedAt: paste.CreatedAt,
		}

		return c.JSON(http.StatusCreated, response)
	})

	return nil
}
