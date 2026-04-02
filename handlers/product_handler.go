package handlers
import (
"net/http"
"strconv"
"github.com/gin-gonic/gin"
"github.com/IlhamMaulana13/GolangTraining-Week4/models"
"github.com/IlhamMaulana13/GolangTraining-Week4/services"
)
type ProductHandler struct {
productService *services.ProductService
}
func NewProductHandler() *ProductHandler {
return &ProductHandler{productService: services.NewProductService()}
}
// GetAll - GET /products?page=1&limit=10&category=makanan
func (h *ProductHandler) GetAll(c *gin.Context) {
page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
category := c.Query("category")
products, total, err := h.productService.GetAll(page, limit, category)
if err != nil {
c.JSON(http.StatusInternalServerError, gin.H{
"success": false, "message": "Gagal mengambil data produk",
})
return
}