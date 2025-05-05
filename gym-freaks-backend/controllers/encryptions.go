package controllers

import (
	"fmt"
	"gym-freaks-backend/middleware"
	"gym-freaks-backend/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type MyClaims struct {
	UserID    int         `json:"userid"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Phone     int         `json:"phone"`
	Dob       models.Date `json:"dob"`
	Role      models.Role `json:"role"`
	CreatedAt time.Time   `json:"createdAt"`
	jwt.RegisteredClaims
}

var _ = godotenv.Load()

func HashPassword(s string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(hashedPassword, password string) bool {
	// Compare the hashed password with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return false // Passwords do not match
	}
	return true // Passwords match
}

func CreateJWT(user models.User) (string, error) {

	// Set claims for the JWT token
	claims := MyClaims{
		UserID:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Dob:       user.Dob,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Replace "your-secret-key" with your actual secret key
	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*MyClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("secret key is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	}, jwt.WithLeeway(5*time.Second), jwt.WithValidMethods([]string{"HS256"})) // Allows slight clock skew

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	return claims, nil
}

func GetTokenPayloadFromRequest(r *http.Request) (*MyClaims, error) {
	token, err := middleware.GetTokenFromRequest(r)
	if err != nil {
		return nil, err
	}
	claims, err := VerifyJWT(token)
	return claims, err

}

func CheckTokenExpired(tokenString string, secretKey string) (bool, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is what you expect (typically HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key for validation
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, fmt.Errorf("error parsing token: %v", err)
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if expiration is in the past
		expiration := claims["exp"].(float64)
		if time.Unix(int64(expiration), 0).Before(time.Now()) {
			return true, nil // Token is expired
		}
		return false, nil // Token is not expired
	} else {
		return false, fmt.Errorf("invalid token")
	}
}

func GetUserIDFromToken(tokenString string) (userid int, err error) {
	claims, err := VerifyJWT(tokenString)
	if err != nil {
		return 0, err
	}
	userid = claims.UserID
	return userid, nil
}

func GetUserRoleFromToken(tokenString string) (role models.Role, err error) {
	claims, err := VerifyJWT(tokenString)
	if err != nil {
		return "", err
	}
	role = claims.Role
	return role, nil
}
