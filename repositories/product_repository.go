package repositories
import (
"github.com/IlhamMaulana13/GolangTraining-Week4/config"
"github.com/IlhamMaulana13/GolangTraining-Week4/models"
)
type ProductRepository struct{}
func NewProductRepository() *ProductRepository {
return &ProductRepository{}
}