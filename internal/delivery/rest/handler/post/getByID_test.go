package post_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/andiksetyawan/boilerplate_svc/internal/model/post/entity"
	"github.com/andiksetyawan/boilerplate_svc/pkg/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

func (s *HandlerSuite) TestHandler_GetByID() {
	s.Run("success", func() {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		id, _ := uuid.NewV7()
		c := s.echo.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		post := entity.Post{
			ID:   id,
			Desc: "foo",
		}

		s.usecaseMock.On("GetByID", mock.Anything, mock.Anything).Return(post, nil).Once()
		err := s.handler.GetByID(c)

		s.NoError(err)
		s.Equal(http.StatusOK, rec.Code)

		expectResponse := response.Response{
			Status:     "OK",
			StatusCode: 200,
			Message:    "fetch post successfully",
			Data:       post,
		}
		b, _ := json.Marshal(expectResponse)
		s.JSONEq(string(b), rec.Body.String())
	})

	s.Run("error_bind", func() {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := s.echo.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("fake-uuid")

		s.logMock.On("Error", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
		s.handler.GetByID(c)

		s.Equal(http.StatusBadRequest, rec.Code)
		//TODO equal body response
	})

	s.Run("error_validate", func() {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := s.echo.NewContext(req, rec)

		s.logMock.On("Error", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
		s.handler.GetByID(c)

		s.Equal(http.StatusBadRequest, rec.Code)
		//TODO equal body response
	})

	s.Run("error", func() {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		id, _ := uuid.NewV7()
		c := s.echo.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		s.usecaseMock.On("GetByID", mock.Anything, mock.Anything).
			Return(entity.Post{}, errors.New("something error")).Once()

		err := s.handler.GetByID(c)

		s.NoError(err)
		s.Equal(http.StatusInternalServerError, rec.Code)

		expectResponse := response.ErrorResponse{
			Response: response.Response{
				Status:     "error",
				StatusCode: http.StatusInternalServerError,
				Message:    http.StatusText(http.StatusInternalServerError),
			},
			Errors: []string{"something error"},
		}
		b, _ := json.Marshal(expectResponse)
		s.JSONEq(string(b), rec.Body.String())
	})
}
