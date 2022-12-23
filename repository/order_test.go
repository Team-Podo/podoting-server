package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type OrderTestSuite struct {
	suite.Suite
	BuyerUID        string
	orderRepository OrderRepository
	mock            sqlmock.Sqlmock
}

func (suite *OrderTestSuite) SetupTest() {
	db, err := gorm.Open(
		sqlite.Open(
			utils.RootPath()+"sqlite/test.db",
		),
		&gorm.Config{},
	)
	suite.NoError(err)

	suite.BuyerUID = uuid.New().String()

	suite.orderRepository = OrderRepository{DB: db.Debug().Begin()}
}

func (suite *OrderTestSuite) MakeOrder() Order {
	var details []OrderDetail
	details = append(details, OrderDetail{
		SeatBookingID:  uint(1),
		OriginalPrice:  uint(1),
		OrderDetailKey: utils.GenerateOrderDetailKey(),
	})

	var order Order
	order.PerformanceID = uint(1)
	order.ScheduleUUID = uuid.New().String()
	order.OrderKey = utils.GenerateOrderKey()
	order.BuyerUID = suite.BuyerUID
	order.Paid = true
	order.Details = details

	err := suite.orderRepository.Save(&order)
	suite.NoError(err)

	return order
}

func (suite *OrderTestSuite) TestCancelOrder() {
	order := suite.MakeOrder()

	err := suite.orderRepository.CancelOrder(&order)

	suite.Equal(true, suite.orderRepository.FindByID(order.ID).Canceled)
	suite.Equal(true, suite.orderRepository.FindByID(order.ID).Details[0].Canceled)

	suite.NoError(err)
}

func (suite *OrderTestSuite) TestCancelOrder_다른_주문은_취소되면_안됨() {
	order := suite.MakeOrder()
	order2 := suite.MakeOrder()

	err := suite.orderRepository.CancelOrder(&order)

	suite.Equal(true, suite.orderRepository.FindByID(order.ID).Canceled)
	suite.Equal(true, suite.orderRepository.FindByID(order.ID).Details[0].Canceled)

	suite.Equal(false, suite.orderRepository.FindByID(order2.ID).Canceled)
	suite.Equal(false, suite.orderRepository.FindByID(order2.ID).Details[0].Canceled)

	suite.NoError(err)
}

func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}
