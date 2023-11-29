package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/alexedwards/flow"
	"github.com/rs/zerolog"

	"github.com/agilefoxHQ/go-function-template/service"
)

type (
	handler struct {
		service *service.Service
		logger  *zerolog.Logger
	}

	// ErrorResponse ...
	ErrorResponse struct {
		Error   error  `json:"error"`
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	// SuccessResponse ...
	SuccessResponse struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message,omitempty"`
	}
)

func NewHandler(s *service.Service, l *zerolog.Logger) http.Handler {
	h := &handler{service: s, logger: l}
	// the whole route generation could also be moved to individual routes.go
	mux := flow.New()
	mux.Use(
	//middleware.Logger,
	//render.SetContentType(render.ContentTypeJSON),
	//middleware.Compress(5),
	//middleware.Recoverer,
	)

	// version
	mux.HandleFunc("/live", h.Version, "GET")

	// v1Router
	const v1Path = "/api/v1"
	mux.Group(
		func(mux *flow.Mux) {
			// VerifyIDToken would be used as a middleware for all routes in this group
			// mux.Use(h.HandlePreFlight, h.VerifyIDToken)
			//mux.HandleFunc(fmt.Sprintf("%s/setup", v1Path), h.SetupAccount, "POST")
			// this accepts regular expression's
			// ex: mux.HandleFunc("/profile/:name/:age|^[0-9]{1,3}$", exampleHandlerFunc2, "GET")
			// mux.HandleFunc(fmt.Sprintf("%s/notify/:id", v1Path), h.HandleHook, "POST")
		},
	)

	return mux
}

// WriteJSON writes a HTTP response with JSON serialized payload.
func WriteJSON(w http.ResponseWriter, code int, content interface{}) {
	if code == http.StatusNoContent {
		// just in case I am stupid elsewhere
		w.WriteHeader(http.StatusNoContent)
		return
	}

	b, err := json.Marshal(content)
	if err != nil {
		code = http.StatusInternalServerError
		// this is the serialisation of ErrorResponse{}
		b = []byte(`{"message":"response serialisation failure."}`)
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)

	_, _ = w.Write(b)
	_, _ = w.Write([]byte{'\n'}) // Be nice to CLI.
}

func WriteJSONErr(w http.ResponseWriter, message string, code int) {
	payload := ErrorResponse{
		Message: message,
		Code:    code,
	}
	w.Header().Set("X-Content-Type-Options", "nosniff")
	WriteJSON(w, code, payload)
}

func printStruct(i interface{}) {

	reflectedStruct := reflect.ValueOf(i)
	// in case you want to list all values stored at the location of this pointer to a struct
	if reflectedStruct.Kind() == reflect.Ptr {
		reflectedStruct = reflectedStruct.Elem()
	}
	structType := reflectedStruct.Type()
	fmt.Println("---------")
	fmt.Printf("Printing struct: %v", structType.String())
	fmt.Println("---------")
	values := make(map[string]interface{}, reflectedStruct.NumField())
	for i := 0; i < reflectedStruct.NumField(); i++ {
		values[structType.Field(i).Name] = reflectedStruct.Field(i).Interface()
		fmt.Printf(
			"%s: %v\n",
			structType.Field(i).Name,
			reflectedStruct.Field(i).Interface(),
		)
	}
	fmt.Println("---------")
}
