package view

type Home struct {
	Title  string
	Active string
}

func NewHome() Home {
	result := Home{
		Active: "home",
		Title:  "Appeals",
	}
	return result
}
