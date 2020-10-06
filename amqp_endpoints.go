package movie_store

import (
	"encoding/json"
	"errors"
	"github.com/djumanoff/amqp"
)

type AMQPEndpointFactory struct {
	movieService MovieService
}

func NewAMQPEndpointFactory(movieService MovieService) *AMQPEndpointFactory {
	return &AMQPEndpointFactory{movieService: movieService}
}

type ErrorSt struct {
	Text string `json:"text"`
}

func (fac *AMQPEndpointFactory) GetMovieByIdAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetMovieByIdCommand{}

		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(&ErrorSt{err.Error()})
		}

		if cmd.Id == 0 {
			return AMQPError(&ErrorSt{errors.New("no movie id").Error()})
		}

		resp, err := cmd.Exec(fac.movieService)
		if err != nil {
			return AMQPError(&ErrorSt{err.Error()})
		}
		return OK(resp)
	}
}

func (fac *AMQPEndpointFactory) CreateMovieAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateMovieCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(err)
		}
		resp, err := cmd.Exec(fac.movieService)
		if err != nil {
			return AMQPError(err)
		}
		return OK(resp)
	}
}

func (fac *AMQPEndpointFactory) ListMoviesAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message{
		cmd := &ListMoviesCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(err)
		}
		resp, err := cmd.Exec(fac.movieService)
		if err != nil {
			return AMQPError(err)
		}
		return OK(resp)
	}
}

func (fac *AMQPEndpointFactory) DeleteMovieAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message{
		cmd := &DeleteMovieCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(err)
		}
		if cmd.Id == 0 {
			return AMQPError(&ErrorSt{errors.New("no movie id").Error()})
		}
		resp, err := cmd.Exec(fac.movieService)
		if err != nil {
			return AMQPError(err)
		}
		return OK(resp)
	}
}

func (fac *AMQPEndpointFactory) UpdateProductAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message{
		cmd := &UpdateMovieCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(err)
		}
		if cmd.Id == 0 {
			return AMQPError(&ErrorSt{errors.New("no movie id").Error()})
		}
		resp, err := cmd.Exec(fac.movieService)
		if err != nil {
			return AMQPError(err)
		}
		return OK(resp)
	}
}

func (fac *AMQPEndpointFactory) GetMovieByNameAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message{
		cmd := &GetMovieByNameCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(err)
		}
		if cmd.Name == "" {
			return AMQPError(&ErrorSt{errors.New("no movie name").Error()})
		}
		resp, err := cmd.Exec(fac.movieService)
		if err != nil {
			return AMQPError(err)
		}
		return OK(resp)
	}
}

func OK(d interface{}) *amqp.Message {
	data, _ := json.Marshal(d)
	return &amqp.Message{Body: data}
}

func AMQPError(e interface{}) *amqp.Message {
	errObj, _ := json.Marshal(e)
	return &amqp.Message{Body: errObj}
}