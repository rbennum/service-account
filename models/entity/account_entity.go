package entity

type AccountEntity struct {
	AccountNum string
	NIK        string
	Balance    int
}

func NewAccountEntity(accountNum string, nik string, balance int) AccountEntity {
	return AccountEntity{
		AccountNum: accountNum,
		NIK:        nik,
		Balance:    balance,
	}
}
