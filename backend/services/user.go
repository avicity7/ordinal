package services

import (
	"context"
	"fmt"
	"os"
	"server/config"
	"server/structs"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(s structs.Signup) error {
	query := `
		INSERT INTO ordinal_user (email, name, password, role_id) 
		VALUES (@Name, LOWER(@Email), @Password, @RoleID)
	`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.Password), 10)
	if err != nil {
		return err
	}

	args := pgx.NamedArgs{
		"Name":     s.Name,
		"Email":    s.Email,
		"Password": hashedPassword,
		"RoleID":   s.RoleID,
	}

	_, err = config.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func Login(login structs.Login) (structs.ReturnedUser, error) {
	var user structs.User
	query := `
		SELECT name, email, password, ou.role_id, our.name FROM ordinal_user ou
		JOIN ordinal_user_role our ON our.role_id = ou.role_id
		WHERE email = LOWER(@Email)
	`
	args := pgx.NamedArgs{
		"Email": login.Email,
	}

	err := config.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Name, &user.Email, &user.Password, &user.RoleID, &user.RoleName)
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return structs.ReturnedUser{}, err
	}

	returnedUser := structs.ReturnedUser{
		Name:     user.Name,
		Email:    user.Email,
		RoleID:   user.RoleID,
		RoleName: user.RoleName,
	}

	return returnedUser, nil
}

func CreateAccessToken(user structs.ReturnedUser) []byte {
	secret := os.Getenv("SECRET")

	access_token, err := jwt.NewBuilder().
		Issuer("server").
		Claim("Name", user.Name).
		Claim("Email", user.Email).
		Claim("Role", user.RoleID).
		Expiration(time.Now().Add(8760 * time.Hour)).
		Build()
	if err != nil {
		fmt.Println("Access Token generation failed.")
	}

	signed_access_token, err := jwt.Sign(access_token, jwt.WithKey(jwa.HS256, []byte(secret)))
	if err != nil {
		fmt.Println(err)
	}

	return signed_access_token
}
