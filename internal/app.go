package internal

import (
	"e-backend/internal/models"
	"e-backend/internal/modules"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/go-playground/locales/uk"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	uk_translations "github.com/go-playground/validator/v10/translations/uk"
)

// BuildTime is time when executable was built
var BuildTime string = "unknown"
var Version string = "0.0.0"

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

type CustomValidationError struct {
	Message  string
	Messages map[string]string
}

func (e *CustomValidationError) Error() string {
	return e.Message
}

func NewCustomValidationError(message string, messages map[string]string) *CustomValidationError {
	return &CustomValidationError{message, messages}
}

func NewCustomValidationFieldError(message string, field string) *CustomValidationError {
	messages := map[string]string{}
	messages[field] = message
	return &CustomValidationError{message, messages}
}

type HTTPErrorResponse struct {
	Message  string
	Messages map[string]string `json:",omitempty"`
}

type HTTPApp struct {
	Core models.Core
}

func NewHTTPApp(config models.Config) HTTPApp {
	return HTTPApp{Core: models.Core{
		Config: config,
	}}
}

func (a *HTTPApp) Run() {
	a.Core.Version = Version
	a.Core.BuildTime = BuildTime

	// Connect to the database
	db, err := gorm.Open(postgres.Open(a.Core.Config.DB.DSN), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	a.Core.DB = db

	// Prepare HTTP-server
	a.Core.Echo = echo.New()
	a.Core.Echo.HideBanner = true
	a.Core.Echo.HTTPErrorHandler = a.httpErrorHandler

	uk := uk.New()
	uni := ut.New(uk, uk)
	trans, _ := uni.GetTranslator("uk")
	validate := validator.New(validator.WithRequiredStructEnabled())
	uk_translations.RegisterDefaultTranslations(validate, trans)

	a.Core.Trans = &trans

	a.Core.Echo.Validator = &CustomValidator{
		validator: validate,
	}

	// a.Core.Echo.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	a.Core.Echo.Use(middleware.Logger())
	a.Core.Echo.Use(middleware.Recover())

	a.Core.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders:    []string{"Server", "Content-Type", "Content-Disposition"},
		AllowCredentials: true,
	}))

	// Run modules
	for _, m := range modules.EnabledModules {
		fmt.Printf("Run module %s\n", m.Name())
		err := m.Run(&a.Core)
		if err != nil {
			log.Println(err)
		}
	}

	// Run HTTP-server
	a.Core.Echo.Logger.Fatal(a.Core.Echo.Start(fmt.Sprintf(":%d", a.Core.Config.HTTP.Port)))
}

func (a *HTTPApp) httpErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var (
		code = http.StatusInternalServerError
		msg  HTTPErrorResponse
	)

	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		msg = HTTPErrorResponse{Message: fmt.Sprintf("%v", e.Message)}
	case validator.ValidationErrors:
		code = http.StatusBadRequest
		messages := make(map[string]string)
		for _, err := range e {
			var message string
			if a.Core.Trans != nil {
				message = err.Translate(*a.Core.Trans)
			} else {
				message = err.Error()
			}
			messages[err.Field()] = message
		}
		msg = HTTPErrorResponse{Message: "Помилка валідації форми", Messages: messages}
	case *CustomValidationError:
		code = http.StatusBadRequest
		msg = HTTPErrorResponse{Message: e.Error(), Messages: e.Messages}
	default:
		msg = HTTPErrorResponse{Message: err.Error()}
	}

	// Send response
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(code)
	} else {
		err = c.JSON(code, msg)
	}
	if err != nil {
		a.Core.Echo.Logger.Error(err)
	}
}
