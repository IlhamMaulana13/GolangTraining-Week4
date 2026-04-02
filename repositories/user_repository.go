package repositories
import (
"github.com/IlhamMaulana13/GolangTraining-Week4/config"
"github.com/IlhamMaulana13/GolangTraining-Week4/models"
)
type UserRepository struct{}
func NewUserRepository() *UserRepository {
return &UserRepository{}
}