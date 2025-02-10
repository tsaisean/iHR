package employee

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"iHR/handler/authenticate"
	mocks "iHR/repositories/mocks"
	"iHR/repositories/model"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe("CreateEmployeeHandler", func() {
	var (
		router       *gin.Engine
		mockEmpRepo  *mocks.EmployeeRepository
		mockAccRepo  *mocks.AccountRepository
		mockAuthRepo *mocks.AuthRepository
		recorder     *httptest.ResponseRecorder
		token        string
	)

	// Shared setup for all tests
	BeforeEach(func() {
		mockEmpRepo = new(mocks.EmployeeRepository)
		empHandler := NewEmployeeHandler(mockEmpRepo)

		mockAccRepo = new(mocks.AccountRepository)
		mockAuthRepo = new(mocks.AuthRepository)
		testSecret := "testsecret"
		authHandler := authenticate.NewAuthenticateHandler(testSecret, mockAccRepo, mockAuthRepo)

		gin.SetMode(gin.TestMode)
		router = gin.Default()
		router.Use(authHandler.AuthMiddleware)
		router.POST("/employee", empHandler.CreateEmployee)

		recorder = httptest.NewRecorder()

		token, _ = authenticate.GenerateToken(testSecret, 1, "testuser", time.Now().Add(10*time.Minute), time.Now())
		token = "Bearer " + token
	})

	Context("When the request is valid", func() {
		It("should create an employee and return 201 status", func() {
			// Arrange
			inputEmployee := &model.Employee{FirstName: "John", LastName: "Doe"}
			mockEmpRepo.On("CreateEmployee", mock.Anything, inputEmployee).Return(&model.Employee{ID: 1, FirstName: "John", LastName: "Doe"}, nil)

			// Act
			executeRequest(router, token, inputEmployee, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusCreated))
			var responseBody map[string]interface{}
			err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
			gomega.Expect(err).To(gomega.BeNil())

			gomega.Expect(responseBody["id"]).To(gomega.Equal(float64(1)))
			gomega.Expect(responseBody["first_name"]).To(gomega.Equal("John"))
			gomega.Expect(responseBody["last_name"]).To(gomega.Equal("Doe"))

			mockEmpRepo.AssertExpectations(GinkgoT())
		})
	})

	Context("When the request is valid, but no token was passed", func() {
		It("should return 401 status", func() {
			// Arrange
			inputEmployee := &model.Employee{FirstName: "John", LastName: "Doe"}

			// Act
			executeRequest(router, "", inputEmployee, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusUnauthorized))
		})
	})

	Context("When the request payload is invalid", func() {
		It("should return 400 status for malformed JSON", func() {
			// Arrange
			invalidJSON := `{"FirstName": "John", "LastName":"}`

			// Act
			executeRequest(router, token, []byte(invalidJSON), recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusBadRequest))
		})

		It("should return 400 status for empty body", func() {
			// Act
			executeRequest(router, token, []byte{}, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusBadRequest))
		})

		It("should return 400 status for mistype field", func() {
			// Arrange
			invalidJSON := `{"first_name": "John", "last_name":12222}`

			// Act
			executeRequest(router, token, []byte(invalidJSON), recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusBadRequest))
		})
	})

	Context("When the repository returns an error", func() {
		It("should return 500 status", func() {
			// Arrange
			inputEmployee := &model.Employee{FirstName: "Jane", LastName: "Doe"}
			mockEmpRepo.On("CreateEmployee", mock.Anything, inputEmployee).Return(nil, errors.New("repository error"))

			// Act
			executeRequest(router, token, inputEmployee, recorder)

			// Assert
			gomega.Expect(recorder.Code).To(gomega.Equal(http.StatusInternalServerError))

			mockEmpRepo.AssertExpectations(GinkgoT())
		})
	})
})

func executeRequest(router *gin.Engine, token string, body interface{}, recorder *httptest.ResponseRecorder) {
	var requestBody []byte
	var err error

	switch v := body.(type) {
	case []byte:
		requestBody = v
	case *model.Employee:
		requestBody, err = json.Marshal(v)
		if err != nil {
			panic("Failed to marshal request body")
		}
	default:
		panic("Unsupported body type")
	}

	// Create and execute the HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/employee", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	router.ServeHTTP(recorder, req)
}
