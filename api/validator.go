package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/trancongthanh1992/samplebank/util"
)

// hook reflection get field, the field of validator function.
var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)

	}
	return false
}
