package dto

//DTO ini untuk validasi keluar masuknya data, jadi misal kita mau get atau post apa itu ntar divalidasi dulu
//kalo lolos bakal dilanjut kalo gagal bakal dibalikin datanya.
//Pada DTO ga ada gorm nya, semuanya json atau binding atau form atau validator
//BookUpdateDTO ini model untuk memvalidasi data update buku
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id.omitempty" form:"user_id.omitempty"`
	//omitempty itu sama dengan optional
}

//BookUpdateDTO is a model that client use when create a new book
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id.omitempty" form:"user_id.omitempty"`
}
