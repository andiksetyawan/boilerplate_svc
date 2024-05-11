package post_test

import (
	"testing"

	"github.com/andiksetyawan/boilerplate_svc/internal/delivery/rest/handler/post"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/andiksetyawan/boilerplate_svc/pkg/httpserver"
	"github.com/andiksetyawan/boilerplate_svc/pkg/response"
	logmock "github.com/andiksetyawan/log/mocks"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
	*require.Assertions

	usecaseMock *usecasemock.Post
	logMock     *logmock.Logger
	echo        *echo.Echo
	resource    resource.Resource

	handler post.Handler
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.usecaseMock = new(usecasemock.Post)
	s.logMock = new(logmock.Logger)
	s.echo = echo.New()
	s.echo.Validator = httpserver.NewValidator(validator.New())

	httpRes, _ := response.New()
	s.resource = resource.Resource{Log: s.logMock, HttpRes: httpRes}
	s.handler = post.New(s.usecaseMock, s.resource)
}

func (s *HandlerSuite) TearDownTest() {}

func (s *HandlerSuite) TestNew() {
	s.Run("when instantiate", func() {
		s.NotNil(post.New(s.usecaseMock, s.resource))
	})
}
