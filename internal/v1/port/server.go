package port

import (
	"context"
	"sync"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/rs/zerolog"

	"github.com/j1cs/api-user/internal/v1/app"
)

type Server struct {
	app    *app.Application
	logger *zerolog.Logger
	Lock   sync.Mutex
}

// Make sure we conform to StrictServerInterface
// https://github.com/deepmap/oapi-codegen/blob/d26d251362ce50640ca4f418eb4922cfb1757856/examples/petstore-expanded/strict/api/petstore.go#L19
var _ StrictServerInterface = (*Server)(nil)

func NewServer(app *app.Application, logger *zerolog.Logger) *Server {
	return &Server{app: app, logger: logger}
}

func (s *Server) GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering post user controller")
	userDomain, err := s.app.Service.User.GetAll(ctx)
	result := UserDomainsToUsers(userDomain)
	return GetUsers200JSONResponse(result), err
}

func (s *Server) PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	newUser := request.Body

	s.logger.Info().Str("requestId", middleware.GetReqID(ctx)).Msg("Entering post user controller")

	domainUser, err := newUser.ToDomain()
	if err != nil {
		responseError := Error{Code: 400, Message: err.Error()}
		return PostUsers400JSONResponse(responseError), nil
	}

	resultUser, err := s.app.Service.User.Create(ctx, domainUser)

	user := User{}
	user.FromDomain(resultUser)

	return PostUsers201JSONResponse(user), err
}

func (s *Server) DeleteUsersUuid(ctx context.Context, request DeleteUsersUuidRequestObject) (DeleteUsersUuidResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetUsersUuid(ctx context.Context, request GetUsersUuidRequestObject) (GetUsersUuidResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PutUsersUuid(ctx context.Context, request PutUsersUuidRequestObject) (PutUsersUuidResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
