package Handlers

import "testApi/internal/api/Repositories"

type Handler struct {
	Repo Repositories.IAlbumRepository
}
