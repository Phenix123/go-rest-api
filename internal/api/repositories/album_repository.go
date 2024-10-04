package repositories

import (
	"database/sql"
	"log"
	"testApi/internal/api/models"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) IAlbumRepository {
	return &AlbumRepository{
		db: db,
	}
}
func (repo *AlbumRepository) GetById(id int64) (models.Album, error) {
	var album models.Album
	row := repo.db.QueryRow("SELECT * FROM albums WHERE id = ?", id)
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	log.Println(*row, album)
	if err != nil {
		return album, err
	}
	return album, nil
}

func (repo *AlbumRepository) GetAll() ([]models.Album, error) {
	var albums []models.Album
	rows, err := repo.db.Query("SELECT * FROM albums")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var album models.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (repo *AlbumRepository) Save(album *models.CreateAlbum) (int64, error) {
	res, err := repo.db.Exec("INSERT INTO albums (title, artist, price) values (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()

	return id, nil
}
