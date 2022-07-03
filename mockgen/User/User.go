package User

type User struct {
	User UserInterface
}

func (u *User) Use() {
	u.User.AddUser(1, "Kenan UZ")

}
