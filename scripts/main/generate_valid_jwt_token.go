package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"maxprofit/config"
)

func main() {
	// negative means before X hours
	// positive means after  X hours
	expirationInHours, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	configs, err := config.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expirationInHours))),
	})

	if signedToken, err := token.SignedString([]byte(configs.SecretKey)); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(signedToken)
	}
}
