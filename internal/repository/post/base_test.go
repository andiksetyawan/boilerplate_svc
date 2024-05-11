package post_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andiksetyawan/boilerplate_svc/internal/repository"
	"github.com/andiksetyawan/boilerplate_svc/internal/repository/post"
	"github.com/andiksetyawan/boilerplate_svc/internal/resource"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type postRepositorySuite struct {
	suite.Suite
	*require.Assertions

	sqlMock sqlmock.Sqlmock
	db      *sqlx.DB

	resource resource.Resource
	repo     repository.Post
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(postRepositorySuite))
}

func (p *postRepositorySuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		p.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	p.sqlMock = mock
	p.db = sqlx.NewDb(db, "postgres")
	p.Assertions = require.New(p.T())
	p.repo = post.NewRepository(p.resource)
}

func (p *postRepositorySuite) TearDownTest() {
	p.db.Close()
}

func (p *postRepositorySuite) TestNewUsecase() {
	p.Run("when instantiate", func() {
		p.NotNil(post.NewRepository(p.resource))
	})
}
