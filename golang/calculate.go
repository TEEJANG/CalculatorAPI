package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/plus", PlusAB)
	r.GET("/minus", MinusAB)
	r.GET("/multiply", MultiplyAB)
	r.GET("/divine", DivineAB)
	return r
}

/*PlusAB is API for Plus A and B*/
func PlusAB(c *gin.Context) {
	A := c.Query("A")
	B := c.Query("B")
	X, Y, valid := ConvertStringtoFloat(A, B)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid Query String")
		return
	}
	var result float64 = X + Y
	resultString := fmt.Sprintf("%f", result)
	c.String(http.StatusOK, resultString)
}

/*MinusAB is API for Minus A and B*/
func MinusAB(c *gin.Context) {
	A := c.Query("A")
	B := c.Query("B")
	X, Y, valid := ConvertStringtoFloat(A, B)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid Query String")
		return
	}
	var result float64 = X - Y
	resultString := fmt.Sprintf("%f", result)
	c.String(http.StatusOK, resultString)
}

/*MultiplyAB is API for Multiply A and B*/
func MultiplyAB(c *gin.Context) {
	A := c.Query("A")
	B := c.Query("B")
	X, Y, valid := ConvertStringtoFloat(A, B)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid Query String")
		return
	}
	var result float64 = X * Y
	resultString := fmt.Sprintf("%f", result)
	c.String(http.StatusOK, resultString)
}

/*DivineAB is API for Divine A and B*/
func DivineAB(c *gin.Context) {
	A := c.Query("A")
	B := c.Query("B")
	X, Y, valid := ConvertStringtoFloat(A, B)
	if !valid {
		c.String(http.StatusBadRequest, "Invalid Query String")
		return
	}
	if Y == 0 {
		c.String(http.StatusBadRequest, "B Cannot be zero")
		return
	}
	var result float64 = X / Y
	resultString := fmt.Sprintf("%f", result)
	c.String(http.StatusOK, resultString)
}

/*ConvertStringtoFloat is a function to convert string to float*/
func ConvertStringtoFloat(A, B string) (float64, float64, bool) {
	var x, y float64
	if s, err := strconv.ParseFloat(A, 64); err == nil {
		x = s
	} else {
		return -1, -1, false
	}
	if s, err := strconv.ParseFloat(B, 64); err == nil {
		y = s
	} else {
		return -1, -1, false
	}
	return x, y, true
}
