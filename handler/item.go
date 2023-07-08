package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

	responses := []item.CatResponse{}
	for _, i := range items {
		thumb, _ := h.itemService.Thumbnail(i.ID)
		res := convertToResponseCatalog(i, thumb)
		responses = append(responses, res)
	}

	// bukus := []song.SongResponse{}
	// for _, b := range songs {
	// 	buku := convertToResponse(b)
	// 	bukus = append(bukus, buku)
	// }
	c.JSON(http.StatusOK, gin.H{
		"data":   responses,
		"Filter": filter,
		"Sort":   sort,
	})
}

func (h *handlerTag) Thumbnail(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	pic, err := h.itemService.Thumbnail(id)

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
		"data": pic.Based,
	})
}

func (h *handlerTag) ItemDetail(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	item, err := h.itemService.ItemDetail(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	images, err := h.itemService.FindImages(id)
	res := convertToResponseDetail(item, images)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
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

// func (h *handlerTag) GetAllItemStock(c *gin.Context) {

// 	items, err := h.itemService.GetAllItemStock()

// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": err,
// 		})
// 		return
// 	}
// 	// bukus := []song.SongResponse{}
// 	// for _, b := range songs {
// 	// 	buku := convertToResponse(b)
// 	// 	bukus = append(bukus, buku)
// 	// }
// 	c.JSON(http.StatusOK, gin.H{
// 		"data":   items,
// 	})
// }

func (h *handlerTag) GetItemStock(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	items, err := h.itemService.GetItemStock(id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": items,
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

func convertToResponseCatalog(i item.Item, t item.Images2) item.CatResponse {

	res := item.CatResponse{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		Thumbnail:   t.Based,
	}
	return res

}

func convertToResponseDetail(i item.Item, images []item.Images2) item.DetailResponse {

	res := item.DetailResponse{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		Images:      images,
	}
	return res

}

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

func (h *handlerTag) GetCart(c *gin.Context) {
	user_email := Ambil(c)
	user, err := h.userService.FindByEmail(user_email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	cart, err := h.itemService.GetCart(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	responses := []item.CartItemResponse{}
	for _, i := range cart {
		thumb, _ := h.itemService.Thumbnail(i.Product_ID)
		item, _ := h.itemService.ItemDetail(i.Product_ID)
		res := convertToResponseCart(item, thumb, i)
		responses = append(responses, res)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responses,
	})
}

func convertToResponseCart(i item.Item, image item.Images2, c item.CartItem) item.CartItemResponse {

	res := item.CartItemResponse{
		ID:    c.ID,
		Name:  i.Name,
		Price: i.Price,
		Qty:   c.Quantity,
		Size:  c.Size,
		Image: image.Based,
	}
	return res

}
