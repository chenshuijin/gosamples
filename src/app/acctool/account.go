package main

import (
	"time"

	"github.com/go-ray/logging"
)

func createAllTables() {
	a := &EthAccount{}
	a.CreateTable()
	createTable(a)
	s := &Secp256Key{}
	createTable(s)
}

func createTable(i DBModel) {
	if err := i.CreateTable(); err != nil {
		logging.Error("create table failed:", err)
	}
}

type EthAccount Account
type BtcAccount Account

type DBModel interface {
	CreateTable() error
}

type Account struct {
	Address   string
	Public    []byte
	Balance   string
	CreatedAt time.Time  `json:"createat"`
	UpdatedAt time.Time  `json:"updateat"`
	DeletedAt *time.Time `json:"deletedat"`
}

func (a *EthAccount) CreateTable() error {
	return DefaultDB().CreateTable(a).Error
}

func (a *EthAccount) Store() error {
	return DefaultDB().Create(a).Error
}

func (a *EthAccount) Update() error {
	return DefaultDB().Model(a).Updates(a).Error
}

type Secp256Key struct {
	Public    []byte
	Private   []byte
	CreatedAt time.Time  `json:"createat"`
	UpdatedAt time.Time  `json:"updateat"`
	DeletedAt *time.Time `json:"deletedat"`
}

func (s *Secp256Key) CreateTable() error {
	return DefaultDB().CreateTable(s).Error
}

func (s *Secp256Key) Store() error {
	return DefaultDB().Create(s).Error
}

func (s *Secp256Key) Update() error {
	return DefaultDB().Model(s).Updates(s).Error
}
