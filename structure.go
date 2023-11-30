package webhook

type Reply struct {
	Message string `bson:"messsage"`
}

type Logindata struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
