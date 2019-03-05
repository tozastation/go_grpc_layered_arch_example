package db

// User is ...
type User struct {
	BaseModel
	Name        string
	CityName    string
	Password    []byte
	AccessToken string
}
