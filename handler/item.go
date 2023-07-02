package handler

import (
	"fmt"
	"net/http"

	"ntika/auth"
	"ntika/item"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/golang-jwt/jwt/v5"
)

type handlerTag struct {
	itemService item.Service
	userService auth.Service
}

func NewHandlerItems(itemService item.Service, songService auth.Service) *handlerTag {
	return &handlerTag{itemService, songService}
}

// // ================================================
// QUERY PARAMS OK
// TODO :
// IMPLEMENT FILTER & SORT
func (h *handlerTag) Catalog(c *gin.Context) {
	filter := c.Query("filter")
	sort := c.Query("sort")

	items, err := h.itemService.FindAll(filter, sort)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	// bukus := []song.SongResponse{}
	// for _, b := range songs {
	// 	buku := convertToResponse(b)
	// 	bukus = append(bukus, buku)
	// }
	c.JSON(http.StatusOK, gin.H{
		"data":   items,
		"Filter": filter,
		"Sort":   sort,
	})
}

func (h *handlerTag) Create(c *gin.Context) {
	// var image item.Images2Input
	var item item.ItemInput
	err := c.ShouldBind(&item)
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
	h.itemService.Create(item)
	// images := item.Images

	// h.itemService.Pap(images)

	c.JSON(http.StatusOK, gin.H{
		"msg": item,
	})
}

func (h *handlerTag) AddToCart(c *gin.Context) {
	// var image item.Images2Input
	println("HAND1")
	var cart item.CartInput
	err := c.ShouldBind(&cart)
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
	println("HAND2")
	user, _ := c.Get("user")
	find, err := h.userService.FindByEmail(user)

	println("HAND")
	h.itemService.AddToCart(cart, find)
	// images := item.Images

	// h.itemService.Pap(images)

	c.JSON(http.StatusOK, gin.H{
		"msg": "BerhasilAddToCart",
	})
}

func (h *handlerTag) UpdateStock(c *gin.Context) {
	// var image item.Images2Input
	var item_stock item.StockInput
	err := c.ShouldBind(&item_stock)
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
	newStock, err := h.itemService.UpdateStock(item_stock)
	fmt.Printf("error handler %s", err)
	fmt.Println("")
	// images := item.Images
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "GAGAL GES, mungkin belum ada",
		})
		return
	}
	// h.itemService.Pap(images)

	c.JSON(http.StatusOK, gin.H{
		"msg":   "berhasil",
		"stock": newStock,
	})
}

// func (h *handlerTag) Pap(c *gin.Context) {
// 	var img item.Images2Input
// 	err := c.ShouldBind(&img)
// 	if err != nil {

// 		messages := []string{}

// 		for _, e := range err.(validator.ValidationErrors) {
// 			errormsg := fmt.Sprintf("Error pada field %s, condition %s", e.Field(), e.ActualTag())
// 			messages = append(messages, errormsg)
// 		}
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": messages,
// 		})
// 		return

// 	}
// 	h.itemService.Pap(img)

// 	c.JSON(http.StatusOK, gin.H{
// 		"msg": img,
// 	})
// }

// // func convertToResponseTag(b song.Song) song.SongResponse {

// // 	buku := song.SongResponse{
// // 		ID:   b.ID,
// // 		YtID: b.YtID,
// // 		// Title:       b.Title,
// // 		// Price:       b.Price,
// // 		// Description: b.Description,
// // 		// Rating:      b.Rating,
// // 		// Discount:    b.Discount,
// // 	}
// // 	return buku

// // }

func (h *handlerTag) AddSize(c *gin.Context) {
	// var image item.Images2Input
	var item item.SizeInput
	err := c.ShouldBind(&item)
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
	h.itemService.AddSize(item)
	// images := item.Images

	// h.itemService.Pap(images)

	c.JSON(http.StatusOK, gin.H{
		"msg": "berhasil",
	})
}
