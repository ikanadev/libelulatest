package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Valor struct {
	ValorID  int     `json:"valorId"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

type ReqData struct {
	IDPrograma      string  `json:"idPrograma"`
	NumeroDocumento string  `json:"numeroDocumento"`
	FechaNacimiento string  `json:"fechaNacimiento"`
	Complemento     *string `json:"complemento"`
	Nombres         string  `json:"nombres"`
	ApellidoMaterno string  `json:"apellidoMaterno"`
	ApellidoPaterno string  `json:"apellidoPaterno"`
	NroTicket       int     `json:"nroTicket"`
	TipoDocumento   string  `json:"tipoDocumento"`
	Sucursal        int     `json:"sucursal"`
	MetodoPago      int     `json:"metodoPago"`
	Correo          string  `json:"correo"`
	URLRetorno      string  `json:"urlRetorno"`
	Valores         []Valor `json:"valores"`
}

type Err struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type Resp struct {
	Estado    bool    `json:"estado"`
	Codigo    int     `json:"codigo"`
	Respuesta string  `json:"respuesta"`
	Mensaje   string  `json:"mensaje"`
	Errors    *string `json:"errors"`
}

func customErr(c *fiber.Ctx, code int, error, message string) error {
	return c.Status(code).JSON(Err{
		Error:            error,
		ErrorDescription: message,
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("status: ok")
	})

	app.Post("/api/matgen", func(c *fiber.Ctx) error {
		authHeaders := c.GetReqHeaders()
		val, ok := authHeaders["Authorization"]
		if !ok {
			return customErr(c, 400, "token_not_found", "token was not provided in the headers")
		}
		if val != "Bearer jwttoken" {
			return customErr(c, 401, "invalid_token", "token in not valid")
		}
		var reqData ReqData
		err := c.BodyParser(&reqData)
		if err != nil {
			return customErr(c, 400, "bad_req_body", "no valid body structure")
		}
		return c.JSON(Resp{
			Estado:    true,
			Codigo:    144,
			Respuesta: "https://www.todotix.com/pagostt?id=3063a5f2-e6c2-4496-8913-ee56b8fbb417",
			Mensaje:   "Deuda registrada con exito, para completar el pago debe redireccionar al cliente a la pasarela de pago",
			Errors:    nil,
		})
	})

	// err := app.Listen("0.0.0.0:8000")
	err := app.Listen(fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		panic(err)
	}
}
