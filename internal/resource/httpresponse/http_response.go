package httpresponse

import (
	"github.com/andiksetyawan/boilerplate_svc/pkg/response"
	"github.com/andiksetyawan/log"
)

func NewHttpResponse(log log.Logger) (httpResponse *response.HttpResponse) {
	httpResponse, _ = response.New(response.WithErrLogger(log))
	return
}
