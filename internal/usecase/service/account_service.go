package service

type AccountService struct{
	DAO interface{
		GetUserFromId(string) (string, error)
		CreateUser(string) error
	}
}

func (_ *AccountService) GetPongFromPing() string {
	return "pong"
}

func (svc *AccountService) GetInfoFromId(id string) (string, error) {
	info, err := svc.DAO.GetUserFromId(id)
	return info, err
}

func (svc *AccountService) PostUser(email string) (string, error) {
	var err error
	err = svc.DAO.CreateUser(email)
	if err != nil {
		return "", err
	}
	return "created", err
}
