package service

type DaoInterface interface {
	GetUserFromId(string) (string, error)
	CreateUser(string) (string, error)
}

type AccountService struct{
	Dao DaoInterface
}

type AccountServiceInterface interface {
	GetPongFromPing() (string)
	GetInfoFromId(string) (string, error)
	PostUser(string) (string, string, error)
}

func (_ *AccountService) GetPongFromPing() string {
	return "pong"
}

func (svc *AccountService) GetInfoFromId(id string) (string, error) {
	info, err := svc.Dao.GetUserFromId(id)
	return info, err
}

func (svc *AccountService) PostUser(email string) (string, string, error) {
	var err error
	var id string
	id, err = svc.Dao.CreateUser(email)
	if err != nil {
		return "", "", err
	}
	return id, email, err
}
