package handlers

import (
	"groupie-tracker/nyeltay/algaliyev/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.notFound(w)
		return
	}

	if r.Method != http.MethodGet {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := h.apiRequest.Api()
	if err != nil {
		h.ServerError(w, err)
		return
	}

	filename := "index.html"

	h.render(w, filename, models.Artists, http.StatusOK)
}

func (h *Handler) Url(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/artists" {
		h.notFound(w)
		return
	}

	if r.Method != http.MethodGet {
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id := query.Get("id")
	_, err = strconv.Atoi(id)
	if strings.HasPrefix(id, "0") || err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	err = h.apiRequest.ApiArtist(id)
	if err != nil {
		h.notFound(w)
		return
	}

	filename := "artist.html"

	h.render(w, filename, models.ArtistOne, http.StatusOK)
}

func (h *Handler) notFound(w http.ResponseWriter) {
	h.clientError(w, http.StatusNotFound)
}
