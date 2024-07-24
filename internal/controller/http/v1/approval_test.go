package v1

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	handler_mocks "github.com/andrew-nino/messaggio/internal/controller/http/v1/mocks"
	"github.com/andrew-nino/messaggio/internal/domain/models"

	"github.com/bmizerany/assert"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
)

func TestHandler_updateStatus(t *testing.T) {

	type mockBehavior func(r *handler_mocks.MockApproval, client models.Answer)

	logger, _ := test.NewNullLogger()

	type fields struct {
		services Registry
	}

	tests := []struct {
		name                 string
		inputBody            string
		inputClient          models.Answer
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
		fields
	}{
		{
			name:      "Ok",
			inputBody: `{"id":126, "approve": 1}`,
			inputClient: models.Answer{
				ID:      126,
				Approve: 1,
			},
			mockBehavior: func(r *handler_mocks.MockApproval, answer models.Answer) {
				r.EXPECT().Approve(answer).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"success"}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"id": 126`,
			inputClient:          models.Answer{},
			mockBehavior:         func(r *handler_mocks.MockApproval, answer models.Answer) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"id":126, "approve": 1}`,
			inputClient: models.Answer{
				ID:      126,
				Approve: 1,
			},
			mockBehavior: func(r *handler_mocks.MockApproval, answer models.Answer) {
				r.EXPECT().Approve(answer).Return(errors.New("failed to approve client"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"failed to approve client"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			handler_mock := handler_mocks.NewMockApproval(c)
			test.mockBehavior(handler_mock, test.inputClient)

			h := &Handler{
				log:      logger,
				services: test.fields.services,
				approval: handler_mock,
			}

			r := gin.New()
			r.POST("/approval", h.updateStatus)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/approval",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
