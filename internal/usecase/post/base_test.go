package post_test

import (
	"testing"

	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/andiksetyawan/boilerplate_svc/internal/usecase"
	"github.com/andiksetyawan/boilerplate_svc/internal/usecase/post"
	repomock "github.com/andiksetyawan/boilerplate_svc/mocks/repository"
	sqlxmock "github.com/andiksetyawan/database/sqlx/mocks"
	logmock "github.com/andiksetyawan/log/mocks"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type postUsecaseSuite struct {
	suite.Suite
	*require.Assertions

	repoMock *repomock.Post
	logMock  *logmock.Logger
	dbMock   *sqlxmock.DB

	resource resource.Resource
	usecase  usecase.Post
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(postUsecaseSuite))
}

func (p *postUsecaseSuite) SetupTest() {
	p.Assertions = require.New(p.T())
	p.repoMock = new(repomock.Post)
	p.logMock = new(logmock.Logger)
	p.dbMock = new(sqlxmock.DB)

	p.resource = resource.Resource{Log: p.logMock, DB: p.dbMock}
	p.usecase = post.NewUsecase(p.repoMock, p.resource)
}

func (p *postUsecaseSuite) TearDownTest() {}

func (p *postUsecaseSuite) TestNewUsecase() {
	p.Run("when instantiate", func() {
		p.NotNil(post.NewUsecase(p.repoMock, p.resource))
	})
}
