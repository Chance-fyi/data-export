package boot

import (
	validator2 "data-export/pkg/validator"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func initValidator() {
	validate := binding.Validator.Engine().(*validator.Validate)
	validator2.InitTrans(validate)
}
