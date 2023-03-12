package main

type user struct {
	id       int
	username int
	password int
}
// Struktur user tidak mengikuti aturan konvensi dalam penamaan field
// Sebaiknya field ditulis dengan huruf besar pada awal kata
// Ex: Id, Username, Password

type userservice struct {
	t []user
}
// Nama field 't' tidak menjelaskan field tersebut digunakan untuk apa, sebaiknya gunakan nama yang lebih deskriptif
// Ex: UserList []user

func (u userservice) getallusers() []user {
	return u.t
}
// Nama method Sebaiknya gunakan format camelCase
// Ex: UserService / userServices, GetAllUsers / getAllUsers

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
// Nama parameter 'id' tidak menjelaskan informasi apa yang harus dimasukkan
// Sebaiknya gunakan nama parameter yang lebih menjelaskan informasi apa yang harus dimasukkan
// Ex: GetUserById(userID int)

// Jika user dengan id yang diminta tidak ditemukan, method getuserbyid mengembalikan sebuah user kosong (user{}). 
// Sebaiknya mengembalikan error atau nilai-nilai default yang lebih baik untuk memberitahu programmer bahwa user yang diminta tidak ditemukan.