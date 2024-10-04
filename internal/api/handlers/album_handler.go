package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testApi/internal/api/models"
)

// GetAlbums godoc
//
//	@Summary	Get all albums
//	@Tags		albums
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}	models.Album
//	@Router		/albums/ [get]
func (h *Handler) GetAlbums(c *gin.Context) {
	albums, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums godoc
//
//	@Summary	Create new album
//	@Tags		albums
//	@Accept		json
//	@Produce	json
//	@Param		request	body		models.CreateAlbum	true	"query params"
//	@Success	200		{object}	models.CreateAlbum
//	@Router		/albums/ [post]
func (h *Handler) PostAlbums(c *gin.Context) {
	var newAlbum models.CreateAlbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	success, err := h.Repo.Save(&newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if success == 0 {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID godoc
//
//	@Summary		Get album by id
//	@Description	get string by ID
//	@Tags			albums
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Album ID"
//	@Success		200	{object}	models.Album
//	@Router			/albums/{id} [get]
func (h *Handler) GetAlbumByID(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)

	a, err := h.Repo.GetById(int64(id))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, map[string]string{"errors": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, a)
	}
}
