package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/onsana/order_service/data"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
	"github.com/onsana/order_service/handlers"
)

type orderStorage interface {
	CreateOrder(order *model.Order) error
	GetAllOrders() []model.Order
	GetOrderById(id uuid.UUID) (*model.Order, error)
	UpdateOrder(order *model.Order) error
	DeleteOrderById(id uuid.UUID) error
}

type addressStorage interface {
	CreateAddress(address *model.Address) error
}

type productStorage interface {
	CreateProducts(order *[]model.Product) error
}

type orderService struct {
	oSt orderStorage
	aS  addressService
	pS  productService
	uS  userService
}

type addressService struct {
	aSt addressStorage
}

type productService struct {
	pSt productStorage
	pG  handlers.ProductGateway
}

type userService struct {
	uG handlers.UserGateway
}

func NewOrderService(oSt orderStorage, aS addressService, pS productService, uS userService) *orderService {
	return &orderService{
		oSt: oSt,
		aS:  aS,
		pS:  pS,
		uS:  uS,
	}
}
func NewAddressService(aSt addressStorage) *addressService {
	return &addressService{aSt: aSt}
}

func NewProductService(pSt productStorage, pG handlers.ProductGateway) *productService {
	return &productService{pSt: pSt, pG: pG}
}

func NewUserService(uG handlers.UserGateway) *userService {
	return &userService{uG: uG}
}

func NewProductGatewayMock(idToProductDto map[uuid.UUID]dto.Product) *handlers.ProductGatewayMock {
	return &handlers.ProductGatewayMock{
		IdToProductDto: idToProductDto,
	}
}

func NewProductGatewayImpl() *handlers.ProductGatewayImpl {
	return &handlers.ProductGatewayImpl{}
}

func NewUserGatewayMock(idToProductDto map[uuid.UUID]dto.UserDto) *handlers.UserGatewayMock {
	return &handlers.UserGatewayMock{
		IdToUserDto: idToProductDto,
	}
}

func NewUserGatewayImpl() *handlers.UserGatewayImpl {
	return &handlers.UserGatewayImpl{}
}

func (a *addressService) CreateAddress(addressDto *dto.Address, order model.Order) (*dto.Address, error) {
	address := data.ConvertAddress(*addressDto, order)
	err := a.aSt.CreateAddress(address)
	if err != nil {
		return addressDto, err
	}
	return data.ConvertAddressToDto(*address), nil
}

func (p *productService) CreateProducts(productsDto *[]dto.Product, order model.Order) (*[]dto.Product, error) {
	products := data.ConvertProduct(*productsDto, order)
	err := p.pSt.CreateProducts(products)
	if err != nil {
		return nil, err
	}
	return data.ConvertProductToDto(*products), nil
}

func (p *productService) ValidateProducts(productsDto *[]dto.Product) (*[]dto.Product, error) {
	realProducts, absentIds := p.pG.GetExistingProducts(productsDto)
	if len(absentIds) > 0 {
		return nil, fmt.Errorf("Order cannot be created due to the absence of products: %v !", absentIds)
	}
	return realProducts, nil
}

func (o *orderService) CreateOrder(orderDto *dto.OrderDto) (uuid.UUID, error) {
	user, err := o.uS.uG.GetExistingUser(&orderDto.UserDto)
	if err != nil {
		return uuid.Nil, err
	}
	if user.Blocked {
		return uuid.Nil, fmt.Errorf("can not create order because uses %s is blocked", user.ID)
	}

	validatedProducts, err := o.pS.ValidateProducts(&orderDto.Products)
	if err != nil {
		return uuid.Nil, err
	}

	order := data.ConvertOrder(*orderDto)
	err = o.oSt.CreateOrder(order)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = o.aS.CreateAddress(&orderDto.Address, *order)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = o.pS.CreateProducts(validatedProducts, *order)
	if err != nil {
		return uuid.Nil, err
	}

	return order.ID, nil
}
func (o *orderService) GetAllOrders() []model.Order {
	return o.oSt.GetAllOrders()
}
func (o *orderService) GetOrderById(id uuid.UUID) (*model.Order, error) {
	return o.oSt.GetOrderById(id)
}

func (o *orderService) DeleteOrderById(id uuid.UUID) error {
	return o.oSt.DeleteOrderById(id)
}

func (o *orderService) UpdateOrder(orderDto *dto.OrderDto) (*dto.OrderDto, error) {
	user, err := o.uS.uG.GetExistingUser(&orderDto.UserDto)
	if err != nil {
		return nil, err
	}
	if user.Blocked {
		return nil, fmt.Errorf("can not update order because uses %s is blocked", user.ID)
	}

	orderModel, err := o.oSt.GetOrderById(orderDto.ID)
	if err != nil {
		return nil, fmt.Errorf("order with id = %v doesn't exsist", orderDto.ID)
	}

	err = validateOrderStatus(orderModel, user)
	if err != nil {
		return nil, err
	}
	order := data.ConvertOrder(*orderDto)
	err = o.oSt.UpdateOrder(order)
	if err != nil {
		return nil, err
	}
	return data.ConvertOrderToDto(order), nil
}

func validateOrderStatus(orderModel *model.Order, user *dto.UserDto) error {
	var isAdmin bool
	roles := user.Roles
	for _, role := range roles {
		if "admin" == role {
			isAdmin = true
		}
	}
	if orderModel != nil && orderModel.Status != "pending" && !isAdmin {
		return fmt.Errorf("order can be updated just in Pending status")
	}
	return nil
}
