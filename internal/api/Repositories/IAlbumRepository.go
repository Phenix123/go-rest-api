package Repositories

import "testApi/internal/api/Models"

// check Implementation
var _ IAlbumRepository = &AlbumRepository{}

type IAlbumRepository interface {
	GetById(int64) (Models.Album, error)
	GetAll() ([]Models.Album, error)
	Save(*Models.CreateAlbum) (int64, error)
}
