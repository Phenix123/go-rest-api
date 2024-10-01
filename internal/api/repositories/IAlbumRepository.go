package repositories

import "testApi/internal/api/models"

// check Implementation
var _ IAlbumRepository = &AlbumRepository{}

type IAlbumRepository interface {
	GetById(int64) (models.Album, error)
	GetAll() ([]models.Album, error)
	Save(*models.CreateAlbum) (int64, error)
}
