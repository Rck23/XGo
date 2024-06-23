package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"XGo/bd"
	"XGo/jwt"
	"XGo/models"
	"XGo/models/usuarios"

	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.ResApi {

	var usuario usuarios.Usuario
	var respuesta models.ResApi

	respuesta.Status = 400

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &usuario)

	if err != nil {
		respuesta.Message = "Credenciales incorrectas " + err.Error()
		return respuesta
	}

	if len(usuario.Email) == 0 {
		respuesta.Message = "El correo electrónico es requerido"
		return respuesta
	}

	usuarioData, existe := bd.IntentoLogin(usuario.Email, usuario.Password)
	if !existe {
		respuesta.Message = "Credenciales incorrectas "
		return respuesta
	}

	jwtKey, err := jwt.GenerarToken(ctx, usuarioData)
	if err != nil {
		respuesta.Message = "Ocurrio un error al intentar generar el token > " + err.Error()
		return respuesta
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		respuesta.Message = "Ocurrio un error al intentar formatear el token > " + err2.Error()
		return respuesta
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}

	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	respuesta.Status = 200
	respuesta.Message = string(token)
	respuesta.CustomResp = res

	return respuesta

}
