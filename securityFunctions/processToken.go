package securityfunctions

import (
	"errors"

	//"log"

	//"log"
	"strings"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/projects/klynx/models"
	dbquery "github.com/projects/klynx/usersManagment/dataBases/dbQuery"
)

var UserKeyField string
var Iduser string
var Status string
var Active string

func ProcessToken(tkn string, keyOfToken string) (*models.Claim, bool, string, error) {

	keyword := []byte(keyOfToken)

	claims := &models.Claim{}

	splitToken := strings.Split(tkn, "Bearer")

	if len(splitToken) != 2 {

		return claims, false, string(""), errors.New("token format is invalid")

	}

	tkn = strings.TrimSpace(splitToken[1])

	tkn2, err := jwt2.ParseWithClaims(tkn, claims, func(token *jwt2.Token) (interface{}, error) {

		return keyword, nil
	})

	if err == nil {

		userlocal, exists, _, _ := dbquery.GetUserByID(claims.ID.Hex())

		if exists {

			UserKeyField = claims.UserKeyField
			Iduser = claims.ID.Hex()

			Status = userlocal.Status
			Active = userlocal.Active

		}

		return claims, exists, Iduser, nil
	}

	if !tkn2.Valid {

		return claims, false, string(""), errors.New("the token is invalid")

	}

	return claims, false, string(""), err

}
