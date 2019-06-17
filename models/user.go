package Models

type User struct {
	ID       string `bson:"_id" json:"_id,omitempty"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`
	Key      string `bson:"key" json:"key"`
	Secret   string `bson:"secret" json:"secret"`
}
