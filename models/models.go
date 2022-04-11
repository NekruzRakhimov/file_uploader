package models

import "time"

type User struct {
	ID           int       `json:"id"`
	FullName     string    `json:"full_name" gorm:"column:name"`
	Name         string    `json:"-"`
	Surname      string    `json:"-"`
	LastName     string    `json:"-"`
	Login        string    `json:"login"`
	Email        string    `json:"email"`
	Password     string    `json:"password,omitempty"`
	Status       string    `json:"status"`
	Organization string    `json:"organization"`
	Roles        []RoleDTO `json:"roles,omitempty"`
}

type Role struct {
	ID          int        `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Rights      []RightDTO `json:"rights,omitempty"`
}

type RoleDTO struct {
	ID          int        `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Rights      []RightDTO `json:"omitempty"`
	IsAttached  bool       `json:"is_attached"`
}

type Right struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Section     string `json:"section"`
	Description string `json:"description"`
}

type Cars struct {
	Brand string `json:"brand"`
}

type RightDTO struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Section     string `json:"section"`
	Description string `json:"description"`
	IsAttached  bool   `json:"is_attached"`
}

type SearchContract struct {
	Id             int       `json:"id"`
	Status         string    `json:"status"`
	Beneficiary    string    `json:"beneficiary"`
	ContractNumber string    `json:"contract_number"`
	ContractType   string    `json:"contract_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Author         string    `json:"author"`
	Amount         string    `json:"amount"`
	EndDate        string    `json:"end_date"`
	Comment        string    `json:"comment"`
}

type ClientBin struct {
	Bin string `json:"bin"`
}

type ContractNumber struct {
	ContractNumber string `json:"contract_number"`
}

type Country struct {
	CountryArr []struct {
		CountryName string `json:"country_name"`
		CountryCode string `json:"country_code"`
	} `json:"country_arr"`
}

type PriceTypeCreate struct {
	PricetypeName     string `json:"pricetype_name"`
	PricetypeCurrency string `json:"pricetype_currency,omitempty"`
	ClientBin         string `json:"client_bin"`
}

type BinPriceType struct {
	ClientBin string `json:"client_bin"`
}

type ContractCode struct {
	ExtContractCode string `json:"ext_contract_code"`
}
