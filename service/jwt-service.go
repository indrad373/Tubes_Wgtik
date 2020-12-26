package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//kita pake interface ya karena kan pake clean architecture
//JWTService is a contract of what jwtservice can do
type JWTService interface {
	//nah disini kita isi kontrak nya
	//kita misal mau generate token apa
	//disini kita generate user_id karena dari user_id kita bisa tahu siapa yang lagi login dll kan
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

//bikin custom klaim

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	//jwt.StandardClaims ini adalah standar jwt untuk expired time gitu, biasa ada issuer nya gitu
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService method is create a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "Indra",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	// secretKey := os.Getenv("JWT_SECRET")
	secretKey := "topsecret"
	if secretKey != "" {
		//kasih hardcode dulu kalo misal secretkey nya ga ketemu
		secretKey = "topsecret"
	}

	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			//ini tokennya akan keluar setelah berapa waktu
			//kalo di aplikasi2 kaya payment2 gitu dana ovo dll, itu kalo misal dalam berapa waktu
			//kita ga pake aplikasi nya itu akan logout sndiri dan topkennya ga valid lagi, itu namanya jwt expiration

			//ini kita coba expire nya 15 menit ya
			//ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),

			//ini yang 1 tahun
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	//sign ini kedalam beberapa algoritma
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//convert ke array byte
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	//t_ adalah sebuah pointer ke jwt
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			//signin method nya ga dikenali
			return nil, fmt.Errorf("Unexpected signin method %v", t_.Header["alg"])
		}

		//kalo ada serangan buat ngasih algo yang salah kita bisa validasi disini
		return []byte(j.secretKey), nil
	})
}
