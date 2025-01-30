package models

type Id struct {
	Id string
}

type IsExists struct {
	TableName  string
	ClomunName string
	ExpValue   string
}

type IsExistsResp struct {
	IsExists bool
	Status   string
}

type Claim struct {
	UserId   string
	UserRole string
}
