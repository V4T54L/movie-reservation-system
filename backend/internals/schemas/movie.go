package schemas

import "time"

type AddMovie struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	PosterURL   string `json:"poster_uri" db:"poster_image"`
	Genre       string `json:"genre" db:"genre"`
}

type UpdateMovie struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	PosterURL   string `json:"poster_uri" db:"poster_image"`
	Genre       string `json:"genre" db:"genre"`
}

type MovieDetail struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	PosterURL   string     `json:"poster_url" db:"poster_image"`
	Genre       string     `json:"genre" db:"genre"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
