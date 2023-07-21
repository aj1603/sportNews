package schemas

type Customer struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	PASSWORD  string `json:"password"`
	EMAIL     string `json:"email"`
	PHONE_NUM int    `json:"phone_num"`
}
