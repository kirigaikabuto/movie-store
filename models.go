package movie_store

type Movie struct {
	Id          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Score       float64 `json:"score,omitempty"`
}

type MovieUpdate struct {
	Id          int64    `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Score       *float64 `json:"score,omitempty"`
}

type MovieStore interface {
	List() ([]Movie, error)
	Create(movie *Movie) (*Movie, error)
	GetById(id int64) (*Movie, error)
	Update(movie *MovieUpdate) (*Movie, error)
	Delete(id int64) error
}
