package database

import (
	"database/sql"
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) SetupTest() {
	_, err := suite.Db.Exec("DELETE FROM orders")
	if err != nil {
		panic(err)
	}
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestFindAllOrders() {
	orderA, err := entity.NewOrder("1", 10.0, 1.0)
	suite.NoError(err)
	suite.NoError(orderA.CalculateFinalPrice())

	orderB, err := entity.NewOrder("2", 15.0, 2.0)
	suite.NoError(err)
	suite.NoError(orderB.CalculateFinalPrice())

	orderC, err := entity.NewOrder("3", 20.0, 3.0)
	suite.NoError(err)
	suite.NoError(orderC.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)

	for _, o := range [3]entity.Order{
		*orderA,
		*orderB,
		*orderC,
	} {
		err = repo.Save(&o)
		suite.NoError(err)
	}

	ordersResult, err := repo.FindAll()
	suite.NoError(err)

	suite.Equal(orderA.ID, ordersResult[0].ID)
	suite.Equal(orderA.Price, ordersResult[0].Price)
	suite.Equal(orderA.Tax, ordersResult[0].Tax)
	suite.Equal(orderA.FinalPrice, ordersResult[0].FinalPrice)

	suite.Equal(orderB.ID, ordersResult[1].ID)
	suite.Equal(orderB.Price, ordersResult[1].Price)
	suite.Equal(orderB.Tax, ordersResult[1].Tax)
	suite.Equal(orderB.FinalPrice, ordersResult[1].FinalPrice)

	suite.Equal(orderC.ID, ordersResult[2].ID)
	suite.Equal(orderC.Price, ordersResult[2].Price)
	suite.Equal(orderC.Tax, ordersResult[2].Tax)
	suite.Equal(orderC.FinalPrice, ordersResult[2].FinalPrice)
}
