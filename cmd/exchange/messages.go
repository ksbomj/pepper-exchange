package main

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"time"
)

// {"id": 123, "direction": "bid", "exchange":"BTCUSD", "number":123,"price":1000 }
//Order is any bid or ask on the exchange
type Order struct {
	ID                uuid.UUID `json:"id"`                // The id of the order
	Direction         Operation `json:"direction"`         // Whether this order is buying (bid) or selling (ask)
	Exchange          Exchange  `json:"exchange"`          // The exchange either BTC/USD, BTC/LTC, BTC/Doge, BTC/XMR(Monero)
	Number            int       `json:"number"`            // The number of coins
	NumberOutstanding int       `json:"numberOutstanding"` //Outstanding coins
	Price             int       `json:"price"`             //price is always in Satoshis
	Timestamp         int       `json:"timestamp"`         // timestamp in nanoseconds
	UserId            string
}

func NewOrder(req OrderRequest) Order {
	uid := uuid.NewV4()
	return Order{
		ID:                uid,
		Direction:         OperationFromStr(req.Direction),
		Exchange:          ExchangeFromStr(req.Exchange),
		Number:            req.Number,
		NumberOutstanding: req.Number,
		Price:             req.Price,
		Timestamp:         time.Now().Nanosecond(),
		UserId:            req.UserID}
}

func NewCancelOrder(req CancelRequest) Order {

	return Order{
		ID:                req.OrderID,
		Direction:         CANCEL,
		Exchange:          ExchangeFromStr(req.Exchange),
		Number:            0,
		NumberOutstanding: 0,
		Price:             0,
		Timestamp:         time.Now().Nanosecond(),
		UserId:            req.UserID}
}

func NewStatusOrder(req StatusRequest) Order {
	uid := uuid.NewV4()
	return Order{
		ID:                uid,
		Direction:         STATUS,
		Exchange:          ExchangeFromStr(req.Exchange),
		Number:            0,
		NumberOutstanding: 0,
		Price:             0,
		Timestamp:         time.Now().Nanosecond(),
		UserId:            req.UserID}
}

//Fill is a match between a bid and ask for x satoshis and y number of coins
type Fill struct {
	ID           uuid.UUID `json:"id"`
	Exchange     Exchange  `json:"exchange"`
	Number       int       `json:"number"`
	Price        int       `json:"price"`
	Timestamp    int       `json:"timestamp"`
	Participants []Order   `json:"participants"`
	Closed       []Order   `json:"closed"`
}

func NewFill(exc Exchange, num int, price int, part []Order, closed []Order) Fill {
	uid := uuid.NewV4()

	return Fill{
		ID:           uid,
		Exchange:     exc,
		Number:       num,
		Price:        price,
		Timestamp:    time.Now().Nanosecond(),
		Participants: part,
		Closed:       closed}
}

func (f Fill) Json() string {
	b, err := json.Marshal(f)
	if err != nil {
		return "Cannot convert to JSON"
	}
	return string(b)
}
