package movie_store

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

var Queries = []string{
	`CREATE TABLE IF NOT EXISTS movies (
		id serial primary key,
		name text,
		photo text,
		description text,
		genre text,
		year text,
		count_episode text,
		score float8
	);`,
}

type postgreStore struct {
	db *sql.DB
}

func NewPostgreStore(cfg Config) (MovieStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &postgreStore{db: db}, err
}

func (ps *postgreStore) List(count int64) ([]Movie, error) {
	var movies []Movie
	data, err := ps.db.Query("select * from movies limit $1", count)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		movie := Movie{}
		err = data.Scan(&movie.Id, &movie.Name, &movie.Photo, &movie.Description, &movie.Genre, &movie.Year, &movie.CountEpisode, &movie.Score)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (ps *postgreStore) Create(movie *Movie) (*Movie, error) {
	err := ps.db.QueryRow("insert into movies (name,photo,description,genre,year,count_episode,score) values ($1,$2,$3,$4,$5,$6,$7) RETURNING id", movie.Name, movie.Photo, movie.Description, movie.Genre, movie.Year, movie.CountEpisode, movie.Score).Scan(&movie.Id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (ps *postgreStore) GetById(id int64) (*Movie, error) {
	movie := &Movie{}
	err := ps.db.QueryRow("select * from movies where id= $1", id).Scan(&movie.Id, &movie.Name, &movie.Photo, &movie.Description, &movie.Genre, &movie.Year, &movie.CountEpisode, &movie.Score)
	if err != nil {
		return nil, err
	}
	if movie.Id == 0 {
		fmt.Println("go to error with movie")
		return nil, errors.New("no data by id")
	}
	return movie, nil
}

func (ps *postgreStore) Update(movie *MovieUpdate) (*Movie, error) {
	query := "update movies set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if movie.Name != nil {
		cnt++
		parts = append(parts, "name = $"+strconv.Itoa(cnt))
		values = append(values, movie.Name)
	}
	if movie.Description != nil {
		cnt++
		parts = append(parts, "description = $"+strconv.Itoa(cnt))
		values = append(values, movie.Description)
	}
	if movie.Score != nil {
		cnt++
		parts = append(parts, "score = $"+strconv.Itoa(cnt))
		values = append(values, movie.Score)
	}
	if len(parts) <= 0 {
		return nil, errors.New("nothing to update")
	}
	cnt++
	query = query + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, movie.Id)
	result, err := ps.db.Exec(query, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("movie not found")
	}

	return ps.GetById(movie.Id)
}

func (ps *postgreStore) Delete(id int64) error {
	_, err := ps.db.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
