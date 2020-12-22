package dto

//UserUpdateDTO is used when client put update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password.omitempty" form:"password.omitempty" validate:"min:6"`
	//password kalo dia kosong ya kasih aja omitempty
}

//UserCreateDTO is used when register a user
type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password.omitempty" form:"password.omitempty" validate:"min:6" binding:"required"`
	//gausah pake id soalnya auto increment
	//password kalo dia kosong ya kasih aja omitempty
}
