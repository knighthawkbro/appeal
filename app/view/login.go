package view

type Login struct {
	Title string
	//Active   string
	Email    string
	Warning  string
	Password string
}

func NewLogin() Login {
	result := Login{
		//Active: "home",
		Title: "Appeal - Login",
	}
	return result
}
