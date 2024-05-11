package post_test

import (
	"context"

	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

func (p *postUsecaseSuite) TestPostUsecase_GetByID() {
	p.Run("success", func() {
		id, _ := uuid.NewV7()
		expectedRes := entity.Post{
			ID:   id,
			Desc: "foo",
		}

		p.dbMock.On("GetMaster").Return(&sqlx.DB{})
		p.repoMock.On("GetByID", mock.Anything, mock.Anything, id).Return(expectedRes, nil).Once()
		post, err := p.usecase.GetByID(context.TODO(), id)

		p.NoError(err)
		p.Equal(expectedRes.ID, post.ID)
		p.Equal(expectedRes.Desc, post.Desc)
	})
}
