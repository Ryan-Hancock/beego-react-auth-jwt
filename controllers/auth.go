package controllers

import (
	"authJWT/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

//AuthController ...
type AuthController struct {
	beego.Controller
}

//URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Validate", c.Validate)
	c.Mapping("Refresh", c.Refresh)
}

type AuthResponse struct {
	Token string `json:"token"`
	Error error  `json:"error"`
}

type LoginRequest struct {
	Username string `json:"useranme"`
	Password string `json:"password"`
}

// Login ...
// @router /auth/login [post]
// @Param   username		body    string true "user name"
// @Param   password		body    string true "password"
func (c *AuthController) Login() {
	var res LoginRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &res)

	id, err := models.GetUserStorage().CheckPassword(res.Username, res.Password)
	if err != nil {
		log.Println(err)
		c.Abort("401")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":      id,
		"refreshTime": time.Now().Add(time.Hour * 8).Unix(),
		"exp":         time.Now().Add(time.Minute * 1).Unix(), //Short time for testing
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(`secret`))
	resp := AuthResponse{
		Token: tokenString,
		Error: err,
	}
	c.Data["json"] = &resp
	c.ServeJSON()
}

// Validate ...
// @router /auth/validate [post]
func (c *AuthController) Validate() {
	var req struct {
		TokenString string `json:"token"`
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	token, err := jwt.Parse(req.TokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(`secret`), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && !token.Valid {
		c.Abort("401")
	} else {
		log.Println(err)
		c.Abort("500")
	}

	c.ServeJSON()
	c.StopRun()
}

// Refresh ...
// @router /auth/refresh [post]
func (c *AuthController) Refresh() {
	var req struct {
		TokenString string `json:"token"`
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	token, err := jwt.Parse(req.TokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(`secret`), nil
	})

	if err != nil {
		c.Abort("500")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Now().Unix() > int64(claims["refreshTime"].(float64)) {
			c.Abort("401")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID":      claims["userID"],
			"refreshTime": time.Now().Add(time.Hour * 8).Unix(),
			"exp":         time.Now().Add(time.Minute * 15), //Short time for testing
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(`secret`))
		resp := AuthResponse{
			Token: tokenString,
			Error: err,
		}
		c.Data["json"] = &resp
		c.ServeJSON()
	}

	c.Abort("401")
}
