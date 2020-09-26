package movie_store


type ListMoviesCommand struct {
}

func (cmd *ListMoviesCommand) Exec(service MovieService) (interface{}, error) {
	return service.ListMovies()
}

type CreateMovieCommand struct {
	Name     string `json:"name"`
	Description    string  `json:"description"`
	Score 	float64 `json:"score"`
}

func (cmd *CreateMovieCommand) Exec(service MovieService) (interface{}, error) {
	return service.CreateMovie(cmd)
}

type GetMovieByIdCommand struct {
	Id int64 `json:"id"`
}

func (cmd *GetMovieByIdCommand) Exec(service MovieService) (interface{}, error) {
	return service.GetMovieById(cmd)
}

type UpdateMovieCommand struct {
	Id       int64   `json:"id"`
	Name     *string `json:"name"`
	Description    *string  `json:"description"`
	Score *float64 `json:"score"`
}

func (cmd *UpdateMovieCommand) Exec(service MovieService) (interface{}, error) {
	return service.UpdateMovie(cmd)
}

type DeleteMovieCommand struct {
	Id int64 `json:"id"`
}

func (cmd *DeleteMovieCommand) Exec(service MovieService) (interface{}, error) {
	return nil,service.DeleteMovie(cmd)
}