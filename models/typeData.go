package Models

type TypeData struct {
	UserId  string `bson:"UserId" json:"-"`
	Success string `bson:"Success" json:"-"`
	Info    string `bson:"Info" json:"-"`
	Warning string `bson:"Warning" json:"-"`
	Error   string `bson:"Error" json:"-"`
}
