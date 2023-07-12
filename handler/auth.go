package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"ntika/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type handler struct {
	userService auth.Service
}

func NewHandler(userService auth.Service) *handler {
	return &handler{userService}
}

func (h *handler) CreateUser(c *gin.Context) {
	var user auth.UserInput

	err := c.ShouldBind(&user)
	if err != nil {

		messages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
			messages = append(messages, errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": messages,
		})
		return

	}

	_, err = h.userService.Create(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal Registrasi",
			"err": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"msgs": "Registrasi Berhasil",
		})
	}

}

func (h *handler) SignIn(c *gin.Context) {

	var signin auth.SignIn

	err := c.ShouldBind(&signin)
	if err != nil {

		messages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
			messages = append(messages, errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": messages,
		})
		return

	}

	res, err := h.userService.SignIn(signin)
	msg := fmt.Sprintf("Berhasil Login Sebagai %s", res.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Password atau Email Salah",
		})
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": res.Email,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"msg":   msg,
			"token": tokenString,
		})
	}

}

func (h *handler) GetAllUsers(c *gin.Context) {

	users, err := h.userService.FindAll()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"You":  user,
		"list": users,
	})
}

// FIX THIS PLS
func (h *handler) UserProfile(c *gin.Context) {

	// user, err := h.userService.FindAll()

	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": err,
	// 	})
	// 	return
	// }
	user_email := Ambil(c)
	user, err := h.userService.FindByEmail(user_email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	u := convertToResponse(user)
	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})

	// user, _ := c.Get("user")
	// c.JSON(http.StatusOK, gin.H{
	// 	"You": user,
	// })
}

func (h *handler) UpdateAddress(c *gin.Context) {
	user_email := Ambil(c)
	fmt.Println("NIH:")
	fmt.Println(user_email)

	var address auth.AddressInput

	err := c.ShouldBind(&address)
	if err != nil {

		messages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
			messages = append(messages, errormsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": messages,
		})
		return

	}

	_, err = h.userService.UpdateAddress(address, user_email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal Update Alamat",
			"err": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"msgs": "Update Alamat Berhasil",
		})
	}
}

func Ambil(c *gin.Context) string {
	// Get the Cookie
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		fmt.Println("ERROR GAES")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Authorization Header Not Found"})
		return "Authorization Header No"
	}

	// Decode/validate

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find the user
		str := fmt.Sprintf("%v", claims["sub"])
		return str

		// fmt.Println(claims["sub"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return "x"
}

func convertToResponse(b auth.User) auth.UserResponse {

	user := auth.UserResponse{
		UserName: b.Name,
		Email:    b.Email,
		Phone:    b.Phone,
		Address:  b.Address,
		// Title:       b.Title,
		// Price:       b.Price,
		// Description: b.Description,
		// Rating:      b.Rating,
		// Discount:    b.Discount,
	}
	return user

}
