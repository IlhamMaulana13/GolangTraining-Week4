package services
import (
"context"
"errors"
"os"
"strconv"
"time"
"github.com/golang-jwt/jwt/v5"
"github.com/IlhamMaulana13/GolangTraining-Week4/config"
"github.com/IlhamMaulana13/GolangTraining-Week4/models"
"github.com/IlhamMaulana13/GolangTraining-Week4/repositories"
"gorm.io/gorm"
)

type AuthService struct {
userRepo *repositories.UserRepository
}
func NewAuthService() *AuthService {
return &AuthService{userRepo: repositories.NewUserRepository()}
}
// VerifyFirebaseToken verifikasi token dari Firebase,
// pastikan email sudah verified, lalu return Backend JWT
func (s *AuthService) VerifyFirebaseToken(firebaseToken string) (string,
*models.User, error) {
// 1. Verifikasi Firebase ID Token ke server Google
token, err := config.FirebaseAuth.VerifyIDToken(context.Background(),
firebaseToken)
if err != nil {
return "", nil, errors.New("firebase token tidak valid ataukadaluarsa")
}