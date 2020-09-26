package movie_store


type MovieService interface {
	ListMovies() ([]Movie, error)
	CreateMovie(cmd *CreateMovieCommand) (*Movie, error)
	GetMovieById(cmd *GetMovieByIdCommand) (*Movie, error)
	UpdateMovie(cmd *UpdateMovieCommand) (*Movie, error)
	DeleteMovie(cmd *DeleteMovieCommand) error
}

type movieService struct {
	movieStore MovieStore
}

func NewMovieService(movieStore MovieStore) MovieService {
	return &movieService{movieStore: movieStore}
}

func (ps *movieService) ListMovies() ([]Movie, error) {
	movies, err := ps.movieStore.List()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (ps *movieService) CreateMovie(cmd *CreateMovieCommand) (*Movie, error) {
	movie := &Movie{
		Name:     cmd.Name,
		Description:    cmd.Description,
		Score: cmd.Score,
	}
	newMovie, err := ps.movieStore.Create(movie)
	if err != nil {
		return nil, err
	}
	return newMovie, nil
}

func (ps *movieService) GetMovieById(cmd *GetMovieByIdCommand) (*Movie, error) {
	movie, err := ps.movieStore.GetById(cmd.Id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (ps *movieService) UpdateMovie(cmd *UpdateMovieCommand) (*Movie, error) {
	updateMovie := &MovieUpdate{}
	updateMovie.Id = cmd.Id
	if cmd.Description != nil {
		updateMovie.Description = cmd.Description
	}
	if cmd.Name != nil {
		updateMovie.Name = cmd.Name
	}
	if cmd.Score != nil {
		updateMovie.Score = cmd.Score
	}
	cmdGetMovieById := &GetMovieByIdCommand{cmd.Id}
	_, err := ps.GetMovieById(cmdGetMovieById)
	if err != nil {
		return nil, err
	}
	updatedMovie, err := ps.movieStore.Update(updateMovie)
	if err != nil {
		return nil, err
	}
	return updatedMovie, nil
}

func (ps *movieService) DeleteMovie(cmd *DeleteMovieCommand) error {
	cmdGetMovieById := &GetMovieByIdCommand{cmd.Id}
	_, err := ps.GetMovieById(cmdGetMovieById)
	if err != nil {
		return err
	}
	err = ps.movieStore.Delete(cmd.Id)
	if err != nil {
		return err
	}
	return nil
}