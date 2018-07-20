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
	CreatedAt time.Time  `gorm:"-"`
	UpdatedAt time.Time  `gorm:"-"`
	DeletedAt *time.Time `gorm:"-"`
}

func (a *Account) GetById(from, to int) []Account {
	accs := []Account{}
	DefaultDB().Raw("SELECT * FROM eth_accounts_with_keyjsonfile where id >= ? and id < ?", from, to).Scan(&accs)
	return accs
}

func (a *Account) Update() error {
	return DefaultDB().Exec("UPDATE eth_accounts_with_keyjsonfile SET balance = ? where address = ?", a.Balance, a.Address).Error
}

func (Account) TableName() string {
	return "eth_accounts_with_keyjsonfile"
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

func (a *EthAccount) GetById(id int) []EthAccount {
	s := time.Now()
	accs := []EthAccount{}
	DefaultDB().Raw("SELECT * FROM eth_accounts where id = ?", id).Scan(&accs)
	logging.Debug("Getbyid cost time:", time.Since(s))
	return accs
}

func (a *EthAccount) UpdateOne() error {
	s := time.Now()
	err := DefaultDB().Exec("UPDATE eth_accounts SET balance = ? where address = ?", a.Balance, a.Address).Error
	logging.Debug("updateone cost time:", time.Since(s))
	return err
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

type Key2File struct {
	Address []byte `asn1:"tag:0"`
	Public  []byte `asn1:"tag:1"`
	Private []byte `asn1:"tag:2"`
}
