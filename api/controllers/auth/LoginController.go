package auth

import (
	"wb2-master/api/utils"
	"wb2-master/api/types"
	"wb2-master/api/utils/password"
	"wb2-master/api/auth"
	"wb2-master/api/models"
	"wb2-master/api/databases"
	"errors"
	
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// FindUser searches the user's table with the condition given
func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return databases.DB.Where(&models.User{}).Take(dest, conds...)
}
// FindUserByEmail searches the user's table with the email given
func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email = ?", email)
}
// Login
func Login(ctx *fiber.Ctx) error {
	loginInput := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, loginInput); err != nil {
		return err
	}
	user := &types.User{}

	err := FindUserByEmail(user, loginInput.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(user.Password, loginInput.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	token, expire := auth.Generate(&auth.TokenPayload{
		ID: user.ID,
	})
	
	return ctx.JSON(&types.AuthResponse{
		User: user,
		Auth: &types.AccessResponse{
			Token: token,
			Expire: expire,
		},
	})
}
