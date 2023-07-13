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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "GAGAL GES, mungkin belum ada",
		})
		return
	}
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

func (h *handlerTag) DeleteCart(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	_, err = h.itemService.DeleteCart(id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	// msg := fmt.Sprintf("buku %s dengan id %d berhasil didelete", item.Title, buku.ID)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Berhasil Delete",
	})
}

func (h *handlerTag) Order(c *gin.Context) {
	// var image item.Images2Input
	var orderi item.OrderInput
	err := c.ShouldBind(&orderi)
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
	user_email := Ambil(c)
	user, err := h.userService.FindByEmail(user_email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	order, err := h.itemService.Order(user.ID, orderi)
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
	for _, i := range cart {
		_, err := h.itemService.CreateOrderItem(i, order.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
			return
		}
	}

	for _, i := range cart {
		_, err := h.itemService.DeleteCart(i.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "berhasil",
	})
}

func (h *handlerTag) UserHistory(c *gin.Context) {
	user_email := Ambil(c)
	user, err := h.userService.FindByEmail(user_email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	orders, err := h.itemService.UserHistory(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}

	responses := []item.HistoryResponse{}
	responses_oi := []item.OrderItemResponse{}
	for _, i := range orders {
		orderitems, _ := h.itemService.GetOrderItem(i.ID)

		for _, j := range orderitems {
			thumb, _ := h.itemService.Thumbnail(j.Product_ID)
			item, _ := h.itemService.ItemDetail(j.Product_ID)
			res1 := convertToResponseOrderItem(thumb, j, item)
			responses_oi = append(responses_oi, res1)
		}

		res := convertToResponseHistory(i, responses_oi)
		responses = append(responses, res)
		responses_oi = []item.OrderItemResponse{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responses,
	})
}

func convertToResponseOrderItem(img item.Images2, oi item.OrderItem, i item.Item) item.OrderItemResponse {
	res := item.OrderItemResponse{
		ID:       oi.ID,
		Name:     i.Name,
		Quantity: oi.Quantity,
		Size:     oi.Size,
		Image:    img.Based,
	}
	return res

}

func convertToResponseHistory(o item.Orders, oi []item.OrderItemResponse) item.HistoryResponse {
	start := o.StartDate.Format("2006-01-02")
	end := o.EndDate.Format("2006-01-02")

	res := item.HistoryResponse{
		ID:              o.ID,
		Shipping_Method: o.Shipping_Method,
		Total:           o.Total_Price,
		Start:           start,
		End:             end,
		Status:          o.Status,
		Items:           oi,
	}
	return res

}

func (h *handlerTag) AdminOrder(c *gin.Context) {
	orders, err := h.itemService.AdminOrder()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

func (h *handlerTag) AdminACC(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal ACC1",
			"err": err,
		})
	}

	orderitems, _ := h.itemService.GetOrderItem(id)

	for _, j := range orderitems {
		// item, _ := h.itemService.ItemDetail(j.Product_ID)
		size := 0
		if j.Size == "s" {
			size = 1
		} else if j.Size == "m" {
			size = 2
		} else if j.Size == "l" {
			size = 3
		} else if j.Size == "xl" {
			size = 4
		}
		PSS, _ := h.itemService.GetItemStock(j.Product_ID)

		for _, i := range PSS {
			if i.Size_ID == size {
				fmt.Println(i.Stock)
				fmt.Println(j.Quantity)
				if i.Stock-j.Quantity < 0 {
					fmt.Println("MIN GAN")
					c.JSON(http.StatusBadRequest, gin.H{
						"msg": "Gagal ACC, stok item(s) jadi minus",
					})
					return
				} else {
					setok := item.StockInput{
						Product_ID: i.Product_ID,
						Size_ID:    i.Size_ID,
						Stock:      j.Quantity,
					}
					h.itemService.MinStock(setok)

				}
			}
		}
		// res1 := convertToResponseOrderItem(thumb, j, item)
		// responses_oi = append(responses_oi, res1)
	}
	_, err = h.itemService.AdminACC(id)

	fmt.Println("BRASIL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal ACC3",
			"err": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"msgs": "ACC Berhasil",
		})
	}
}

func (h *handlerTag) AdminCancel(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal Cancel",
			"err": err,
		})
	}
	_, err = h.itemService.AdminCancel(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal Cancel",
			"err": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"msgs": "Cancel Berhasil",
		})
	}
}

func (h *handlerTag) AdminFin(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal Fin",
			"err": err,
		})
	}
	orderitems, _ := h.itemService.GetOrderItem(id)

	for _, j := range orderitems {
		// item, _ := h.itemService.ItemDetail(j.Product_ID)
		size := 0
		if j.Size == "s" {
			size = 1
		} else if j.Size == "m" {
			size = 2
		} else if j.Size == "l" {
			size = 3
		} else if j.Size == "xl" {
			size = 4
		}
		PSS, _ := h.itemService.GetItemStock(j.Product_ID)

		for _, i := range PSS {
			if i.Size_ID == size {
				setok := item.StockInput{
					Product_ID: i.Product_ID,
					Size_ID:    i.Size_ID,
					Stock:      j.Quantity,
				}
				h.itemService.UpdateStock(setok)
			}
		}
	}

	_, err = h.itemService.AdminFin(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Gagal Fin",
			"err": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"msgs": "Fin Berhasil",
		})
	}
}

func (h *handlerTag) BestSeller(c *gin.Context) {
	// filter := c.Query("filter")
	// sort := c.Query("sort")
	items, err := h.itemService.FindAll("", "")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	// responses := []item.CatResponse{}
	responses := []item.BestRes{}
	for _, i := range items {
		best, _ := h.itemService.CountSell(i.ID)
		res := convertToResponseBest(i.ID, best)
		// res := best
		responses = append(responses, res)
	}

	terbaik := item.BestRes{}
	for _, i := range responses {
		// best, _ := h.itemService.CountSell(i.ID)
		// res := convertToResponseBest(i.ID, best)
		// res := best
		if i.Sale > terbaik.Sale {
			terbaik = i
		}
		// responses = append(responses, res)
	}

	// newarr, err := h.itemService.NewArr()
	bez, err := h.itemService.ItemDetail(terbaik.ID)
	thumb, _ := h.itemService.Thumbnail(bez.ID)
	best_final := convertToResponseCatalog(bez, thumb)
	responsesBest := []item.CatResponse{}
	responsesBest = append(responsesBest, best_final)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			// "data": responses,
			// "atas": terbaik,
			// "best": bez,
			"best": responsesBest,
		})
	}
}
func convertToResponseBest(id int, sales int64) item.BestRes {
	res := item.BestRes{
		ID:   id,
		Sale: sales,
	}
	return res

}
func (h *handlerTag) NewArr(c *gin.Context) {
	// filter := c.Query("filter")
	// sort := c.Query("sort")

	// best, err := h.itemService.BestSeller()
	newarr, err := h.itemService.NewArr()
	new := []item.Item{}

	for i := 0; i < 2; i++ {
		new = append(new, newarr[i])
	}

	responses := []item.CatResponse{}
	for _, i := range new {
		thumb, _ := h.itemService.Thumbnail(i.ID)
		res := convertToResponseCatalog(i, thumb)
		responses = append(responses, res)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"data": responses,
		})
	}
}

func (h *handlerTag) Recommended(c *gin.Context) {
	items, err := h.itemService.FindAll("", "")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	}
	// responses := []item.CatResponse{}
	responses := []item.BestRes{}
	for _, i := range items {
		best, _ := h.itemService.CountSell(i.ID)
		res := convertToResponseBest(i.ID, best)
		// res := best
		responses = append(responses, res)
	}

	terbaik := item.BestRes{}
	for _, i := range responses {
		// best, _ := h.itemService.CountSell(i.ID)
		// res := convertToResponseBest(i.ID, best)
		// res := best
		if i.Sale > terbaik.Sale {
			terbaik = i
		}
		// responses = append(responses, res)
	}

	// newarr, err := h.itemService.NewArr()
	bez, err := h.itemService.ItemDetail(terbaik.ID)
	// filter := c.Query("filter")
	// sort := c.Query("sort")

	// best, err := h.itemService.BestSeller()
	newarr, err := h.itemService.NewArr()
	new := []item.Item{}

	for i := 0; i < 2; i++ {
		new = append(new, newarr[i])
	}

	final := []item.Item{}
	final = append(final, bez)
	final = append(final, new[0])
	final = append(final, new[1])

	responsesfinal := []item.CatResponse{}
	for _, i := range final {
		thumb, _ := h.itemService.Thumbnail(i.ID)
		resfinal := convertToResponseCatalog(i, thumb)
		responsesfinal = append(responsesfinal, resfinal)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"data": responsesfinal,
		})
	}
}
