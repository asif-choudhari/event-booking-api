package models

import "event-booking-api/utils"

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) EncodePassword() error {
	var err error
	u.Password, err = utils.Encode(u.Password)
	return err
}

func (u *User) ComparePassword(encodedPassword string) bool {
	return utils.Compare(u.Password, encodedPassword)
}

func (u *User) GenerateToken() (string, error) {
	return utils.GenerateToken(u.Email, u.Id)
}
