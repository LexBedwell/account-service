package service

type DAOInterface interface {
	GetUserFromId(string) (string, error)
	CreateUser(string) (string, error)
}

type AccountService struct{
	DAO DAOInterface
}

func (_ *AccountService) GetPongFromPing() string {
	return "pong"
}

func (svc *AccountService) GetInfoFromId(id string) (string, error) {
	info, err := svc.DAO.GetUserFromId(id)
	return info, err
}

func (svc *AccountService) PostUser(email string) (string, string, error) {
	var err error
	var id string
	id, err = svc.DAO.CreateUser(email)
	if err != nil {
		return "", "", err
	}
	return id, email, err
}
