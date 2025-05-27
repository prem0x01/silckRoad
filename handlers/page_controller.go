package handlers

import (
	"net/http"

	"github.com/prem0x01/silckRoad/templates"
)

type pageController struct{}

// constructor for pageController
func PageController() pageController {
	return pageController{}
}

func (c *pageController) Index(w http.ResponseWriter, r *http.Request) {
	templates.PagesIndex().Render(r.Context(), w)
}
