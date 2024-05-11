package post_test

import (
	"context"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func (p *postRepositorySuite) TestPostRepository_GetByID() {
	id, _ := uuid.NewV7()
	rows := sqlmock.NewRows([]string{"id", "desc"}).AddRow(id, "foo")

	p.sqlMock.ExpectQuery("SELECT * FROM post WHERE id=$1").WithArgs(id).WillReturnRows(rows)

	post, err := p.repo.GetByID(context.TODO(), p.db, id)

	p.NoError(err)
	p.Equal(id, post.ID)
	p.Equal("foo", post.Desc)
}
