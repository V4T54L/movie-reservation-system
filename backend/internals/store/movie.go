package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/V4T54L/movie-reservation-system/internals/schemas"
)

func (s *postgresStore) AddMovie(ctx context.Context, movie schemas.AddMovie) error {
	err := s.connect(ctx)
	if err != nil {
		return err
	}
	defer s.close()

	if _, err := s.db.NamedExecContext(
		ctx,
		`INSERT INTO movies 
		(title,description,poster_image,genre)
		VALUES
		(:title,:description,:poster_image,:genre)`,
		movie); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return fmt.Errorf("totally unexpected error : %s", err)
		}
		return err
	}
	return nil
}

func (s *postgresStore) GetMovieByID(ctx context.Context, movieID int) (*schemas.MovieDetail, error) {
	err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.close()

	movie := schemas.MovieDetail{}

	if err := s.db.SelectContext(
		ctx,
		&movie,
		`SELECT id,title,description,poster_image,genre, created_at, updated_at
		FROM users where id=$1;`,
		movieID); err != nil {
		return nil, err
	}
	return &movie, nil
}

func (s *postgresStore) GetMovieDetails(ctx context.Context) ([]schemas.MovieDetail, error) {
	err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.close()

	movies := []schemas.MovieDetail{}

	if err := s.db.SelectContext(
		ctx,
		&movies,
		`SELECT 
			id, title, description, poster_image,
			genre, created_at, updated_at 
		FROM movies;`); err != nil {
		return nil, err
	}
	return movies, nil
}
