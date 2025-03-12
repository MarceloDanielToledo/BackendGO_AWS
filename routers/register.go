package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"context"
	"encoding/json"
	"fmt"
)

func Register(ctx context.Context) models.ResponseAPI {
	var t models.User
	var r models.ResponseAPI

	r.StatusCode = 400

	fmt.Println("Register")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Email required"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Email) < 6 {
		r.Message = "Password required min 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, exist, _ := bd.CheckAlreadyExist(t.Email)

	if exist {
		r.Message = "Email already registered"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.AddRegister(t)

	if err != nil {
		r.Message = "A error ocurred in register method"
		fmt.Println(r.Message)
		return r
	}
	if !status {
		r.Message = "A error ocurred in register method"
		fmt.Println(r.Message)
		return r

	}

	r.StatusCode = 200
	r.Message = "Register OK"
	fmt.Println(r.Message)
	return r

}
