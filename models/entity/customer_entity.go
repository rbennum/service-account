package entity

type CustomerEntity struct {
	Name     string
	NIK      string
	PhoneNum string
}

func NewCustomerEntity(name string, nik string, phoneNum string) CustomerEntity {
	return CustomerEntity{
		Name:     name,
		NIK:      nik,
		PhoneNum: phoneNum,
	}
}
