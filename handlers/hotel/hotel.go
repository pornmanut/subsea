package handlers

import (
	"net/http"
	"subsea/data"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
)

type Hotels struct {
	db *data.HotelsDB
}

func NewHotels(l hclog.Logger, db *data.HotelsDB) *Hotels {
	return &Hotels{db: db}
}

func (h *Hotels) ListAll(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! List All")
}
