package service

type Admin struct {
	Order *Order `inject:"Config.NewOrder()"`
}

func NewAdmin() *Admin {
	return &Admin{}
}
