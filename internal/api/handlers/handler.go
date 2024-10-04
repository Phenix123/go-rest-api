package handlers

import "testApi/internal/api/repositories"

type Handler struct {
	Repo repositories.IAlbumRepository
}
