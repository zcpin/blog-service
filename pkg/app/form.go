package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type VaildError struct {
	Key     string
	Message string
}

type VaildErrors []*VaildError

func (v *VaildError) Error() string {
	return v.Message
}

func (v VaildErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v VaildErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, VaildErrors) {
	var errs VaildErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return true, nil
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &VaildError{
				Key:     key,
				Message: value,
			})
		}
		return true, errs
	}
	return false, nil
}
