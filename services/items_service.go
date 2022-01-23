package services

//example service & the interface
var (
	ItemService itemsInterface = &itemsService{}
)

type itemsInterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct {
}

func (i itemsService) GetItem() {
	panic("implement me")
}

func (i itemsService) SaveItem() {
	panic("implement me")
}

