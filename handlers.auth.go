package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	valid "github.com/asaskevich/govalidator"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AuthHandler struct {
	db              *gorm.DB
	responseHandler ResponseHandler
}

func InitAuthHandler(app *App) *AuthHandler {
	h := &AuthHandler{
		app.Db(),
		app.ResponseHandler(),
	}

	app.Engine().POST("/token", h.Token)

	return h
}

func (h *AuthHandler) Token(c *gin.Context) {
	data := struct {
		Tel      string `form:"tel" valid:"required~手机不能为空,int~手机必须是数字,stringlength(11|11)~手机必须为11位"`
		Password string `form:"password" valid:"required~密码不能为空,stringlength(6|60)~密码至少6位"`
		Country  string `form:"country" valid:"required~国家区号不能为空,int~国家区号必须是数字"`
	}{}

	if err := c.ShouldBind(&data); err != nil {
		// todo not malformed json. fix error
		h.responseHandler.MalformedJSON(c)
		return
	}

	if _, err := valid.ValidateStruct(data); err != nil {
		h.responseHandler.ValidationErrors(c, err)
		return
	}

	user := new(User)
	password := []byte(data.Password)
	hasher := md5.New()
	hasher.Write(password)
	if err := h.db.Where("tel = ? and pwd = ?", data.Country+"-"+data.Tel, hex.EncodeToString(hasher.Sum(nil))).Find(user).Error; err != nil {
		h.responseHandler.Error(c, AuthenticationError, http.StatusBadRequest, "Tel or Password is incorrect")
		return
	}

	token, err := GenerateToken(user.ID, user.Tel)
	if err != nil {
		h.responseHandler.GenerateTokenError(c)
		return
	}
	h.responseHandler.JSONToken(c, http.StatusOK, user, token)
}
