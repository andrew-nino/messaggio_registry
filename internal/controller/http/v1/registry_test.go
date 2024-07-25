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

func TestHandler_addClient(t *testing.T) {

	type mockBehavior func(r *handler_mocks.MockRegistry, client models.Client)

	logger, _ := test.NewNullLogger()

	type fields struct {
		approval Approval
	}

	tests := []struct {
		name                 string
		inputBody            string
		inputClient          models.Client
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
		fields
	}{
		{
			name:      "Ok",
			inputBody: `{ "surname":"Termit", "name":"Andrew", "patronymic":"Nino", "email":"andrew.nino@example.com"}`,
			inputClient: models.Client{
				Surname:    "Termit",
				Name:       "Andrew",
				Patronymic: "Nino",
				Email:      "andrew.nino@example.com",
			},
			mockBehavior: func(r *handler_mocks.MockRegistry, client models.Client) {
				r.EXPECT().RegisterClient(client).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"message":"Client added successfully"}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"name": 126}`,
			inputClient:          models.Client{},
			mockBehavior:         func(r *handler_mocks.MockRegistry, Client models.Client) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{ "surname":"Termit", "name":"Andrew", "patronymic":"Nino", "email":"andrew.nino@example.com"}`,
			inputClient: models.Client{
				Surname:    "Termit",
				Name:       "Andrew",
				Patronymic: "Nino",
				Email:      "andrew.nino@example.com",
			},
			mockBehavior: func(r *handler_mocks.MockRegistry, client models.Client) {
				r.EXPECT().RegisterClient(client).Return(0, errors.New("internal server error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"internal server error"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			handler_mock := handler_mocks.NewMockRegistry(c)
			test.mockBehavior(handler_mock, test.inputClient)

			h := &Handler{
				log:      logger,
				services: handler_mock,
				approval: test.fields.approval,
			}

			r := gin.New()
			r.POST("/add", h.addClient)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/add",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
