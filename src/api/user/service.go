package users

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	res "sport_news/src/api/user/schemas"
)

func register(ctx *gin.Context) {
	var customer res.Register

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &customer)

	register_(&customer)

	ctx.JSON(201, "Successfully created")
}

func login(ctx *gin.Context) {
	var customer res.Update

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &customer)

	token, errors := login_(&customer)
	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.JSON(200, token)

}

func update(ctx *gin.Context) {
	var customer res.Update

	data := ctx.MustGet("data")
	byte_data, _ := json.Marshal(data)
	json.Unmarshal(byte_data, &customer)

	update_(&customer)

	ctx.JSON(200, "Successfully updated")
}

// func opt(ctx *gin.Context) {
// 	key, err := totp.Generate(totp.GenerateOpts{
// 		Issuer:      "Jora",
// 		AccountName: "jora@gmail.com",
// 	})
// 	if err != nil {
// 		fmt.Println("Error generating TOTP key:", err)
// 		return
// 	}
// 	code, err := totp.GenerateCode(key.Secret(), time.Now())
// 	if err != nil {
// 		fmt.Println("Error generating TOTP code:", err)
// 		return
// 	}

// 	fmt.Println("Generated TOTP code:", code)
// }
