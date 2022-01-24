package generate

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var GenerateJwt = &cobra.Command{
	Use:   "jwt",
	Short: "Generate JSON Web Token (JWT)",
	Long:  `generate various amount and/or validity of JSON web token (JWT)`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < int(Total); i++ {
			token, err := generateJwt()
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
			fmt.Println(token)
		}
	},
}

func generateJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	var mySigningKey = []byte("")

	claim := token.Claims.(jwt.MapClaims)

	if IsValid {
		claim["exp"] = time.Now().Add(time.Minute * 30).Unix() // token expire in 30 minutes
	} else {
		claim["exp"] = time.Now().Add(-(time.Minute * 30)).Unix() // token expire 30 minutes ago
	}
	claim["iat"] = time.Now().Unix()
	claim["jti"] = uuid.NewV4()
	claim["iss"] = "ting-cli-tools"

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
