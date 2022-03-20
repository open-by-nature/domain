package fosterhome

type FosterHome struct {
	Id          uint   `json:"id" bson:"id" msgpack:"id" validate:"empty=false"`
	Name        string `json:"name" bson:"name" msgpack:"name" validate:"empty=false"`
	Address     string `json:"address" bson:"address" msgpack:"address"`
	City        string `json:"city" bson:"city" msgpack:"city"`
	PostalCode  string `json:"postalcode" bson:"postalcode" msgpack:"postalcode"`
	State       string `json:"state" bson:"state" msgpack:"state"`
	PhoneNumber string `json:"phonenumber" bson:"phonenumber" msgpack:"phonenumber"`
	EMail       string `json:"email" bson:"email" msgpack:"email"`
}
