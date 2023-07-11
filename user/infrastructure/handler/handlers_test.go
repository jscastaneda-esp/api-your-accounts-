package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"
	"time"
	"your-accounts-api/shared/domain/validation"
	"your-accounts-api/user/application"
	"your-accounts-api/user/application/mocks"
	"your-accounts-api/user/domain"
	"your-accounts-api/user/infrastructure/model"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite
	uid   string
	email string
	token string
	app   *fiber.App
	mock  *mocks.IUserApp
}

func (suite *TestSuite) SetupSuite() {
	suite.uid = "01234567890123456789012345678901"
	suite.email = "example@exaple.com"
	suite.token = "<token>"
}

func (suite *TestSuite) SetupTest() {
	suite.mock = mocks.NewIUserApp(suite.T())
	ctrl := &controller{
		app: suite.mock,
	}

	suite.app = fiber.New()
	suite.app.Post("/", ctrl.create)
	suite.app.Post("/auth", ctrl.auth)
}

func (suite *TestSuite) TestCreate201() {
	require := require.New(suite.T())
	requestBody := &model.CreateRequest{
		UID:   suite.uid,
		Email: suite.email,
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	result := &domain.User{
		ID:        1,
		UID:       requestBody.UID,
		Email:     requestBody.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	suite.mock.On("SignUp", mock.Anything, mock.Anything).Return(result, nil)
	expectedBody, err := json.Marshal(model.CreateResponse{
		ID:        result.ID,
		UID:       result.UID,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
	require.NoError(err)

	request := httptest.NewRequest(fiber.MethodPost, "/", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusCreated, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedBody, resp)
}

func (suite *TestSuite) TestCreate400() {
	require := require.New(suite.T())

	request := httptest.NewRequest(fiber.MethodPost, "/", nil)
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func (suite *TestSuite) TestCreate409() {
	require := require.New(suite.T())
	requestBody := &model.CreateRequest{
		UID:   suite.uid,
		Email: suite.email,
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	suite.mock.On("SignUp", mock.Anything, mock.Anything).Return(nil, application.ErrUserAlreadyExists)
	expectedErr := []byte(application.ErrUserAlreadyExists.Error())

	request := httptest.NewRequest(fiber.MethodPost, "/", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusConflict, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedErr, resp)
}

func (suite *TestSuite) TestCreate422() {
	require := require.New(suite.T())
	requestBody := &model.CreateRequest{
		UID:   "invalid",
		Email: suite.email,
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	validationErrors := []*validation.ErrorResponse{}
	validationErrors = append(validationErrors, &validation.ErrorResponse{
		Field:      "uid",
		Constraint: "len=32",
	})
	expectedBody, err := json.Marshal(validationErrors)
	require.NoError(err)

	request := httptest.NewRequest(fiber.MethodPost, "/", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusUnprocessableEntity, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedBody, resp)
}

func (suite *TestSuite) TestCreate500() {
	require := require.New(suite.T())
	requestBody := &model.CreateRequest{
		UID:   suite.uid,
		Email: suite.email,
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	suite.mock.On("SignUp", mock.Anything, mock.Anything).Return(nil, gorm.ErrInvalidField)
	expectedErr := []byte("Error sign up user")

	request := httptest.NewRequest(fiber.MethodPost, "/", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusInternalServerError, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedErr, resp)
}

func (suite *TestSuite) TestAuth200() {
	require := require.New(suite.T())
	requestBody := &model.AuthRequest{
		CreateRequest: model.CreateRequest{
			UID:   suite.uid,
			Email: suite.email,
		},
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	suite.mock.On("Auth", mock.Anything, suite.uid, suite.email).Return(suite.token, nil)
	expectedBody, err := json.Marshal(model.AuthResponse{
		Token: suite.token,
	})
	require.NoError(err)

	request := httptest.NewRequest(fiber.MethodPost, "/auth", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusOK, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedBody, resp)
}

func (suite *TestSuite) TestAuth400() {
	require := require.New(suite.T())

	request := httptest.NewRequest(fiber.MethodPost, "/", nil)
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func (suite *TestSuite) TestAuth401() {
	require := require.New(suite.T())
	requestBody := &model.AuthRequest{
		CreateRequest: model.CreateRequest{
			UID:   suite.uid,
			Email: suite.email,
		},
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	suite.mock.On("Auth", mock.Anything, suite.uid, suite.email).Return("", gorm.ErrRecordNotFound)
	expectedErr := []byte("Invalid credentials")

	request := httptest.NewRequest(fiber.MethodPost, "/auth", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusUnauthorized, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedErr, resp)
}

func (suite *TestSuite) TestAuth422() {
	require := require.New(suite.T())
	requestBody := &model.AuthRequest{
		CreateRequest: model.CreateRequest{
			UID:   "invalid",
			Email: suite.email,
		},
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	validationErrors := []*validation.ErrorResponse{}
	validationErrors = append(validationErrors, &validation.ErrorResponse{
		Field:      "uid",
		Constraint: "len=32",
	})
	expectedBody, err := json.Marshal(validationErrors)
	require.NoError(err)

	request := httptest.NewRequest(fiber.MethodPost, "/auth", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusUnprocessableEntity, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedBody, resp)
}

func (suite *TestSuite) TestAuth500() {
	require := require.New(suite.T())
	requestBody := &model.AuthRequest{
		CreateRequest: model.CreateRequest{
			UID:   suite.uid,
			Email: suite.email,
		},
	}
	body, err := json.Marshal(requestBody)
	require.NoError(err)
	suite.mock.On("Auth", mock.Anything, suite.uid, suite.email).Return("", gorm.ErrInvalidField)
	expectedErr := []byte("Error authenticate user")

	request := httptest.NewRequest(fiber.MethodPost, "/auth", bytes.NewReader(body))
	request.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	response, err := suite.app.Test(request)

	require.NoError(err)
	require.NotNil(response)
	require.Equal(fiber.StatusInternalServerError, response.StatusCode)
	resp, err := io.ReadAll(response.Body)
	require.NoError(err)
	require.Equal(expectedErr, resp)
}

func (suite *TestSuite) TestNewRoute() {
	require := require.New(suite.T())
	app := fiber.New()

	NewRoute(app)

	routes := app.GetRoutes()
	require.Len(routes, 2)

	route1 := routes[0]
	require.Equal(fiber.MethodPost, route1.Method)
	require.Equal("/user/", route1.Path)
	require.Len(route1.Handlers, 1)

	route2 := routes[1]
	require.Equal(fiber.MethodPost, route2.Method)
	require.Equal("/user/auth", route2.Path)
	require.Len(route2.Handlers, 1)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
