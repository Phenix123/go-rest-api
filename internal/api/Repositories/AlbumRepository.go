package Repositories

import (
	"database/sql"
	"log"
	"testApi/internal/api/Models"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) IAlbumRepository {
	return &AlbumRepository{
		db: db,
	}
}
func (repo *AlbumRepository) GetById(id int64) (Models.Album, error) {
	var album Models.Album
	row := repo.db.QueryRow("SELECT * FROM albums WHERE id = ?", id)
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	log.Println(*row, album)
	if err != nil {
		return album, err
	}
	return album, nil
}

func (repo *AlbumRepository) GetAll() ([]Models.Album, error) {
	var albums []Models.Album
	rows, err := repo.db.Query("SELECT * FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album Models.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (repo *AlbumRepository) Save(album *Models.CreateAlbum) (int64, error) {
	res, err := repo.db.Exec("INSERT INTO albums (title, artist, price) values (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()

	return id, nil
}
