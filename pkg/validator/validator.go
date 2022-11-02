package validator

import (
	"data-export/pkg/console"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

func InitTrans(validate *validator.Validate) {
	zhT := zh.New()
	trans, _ = ut.New(zhT, zhT).GetTranslator("zh")
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	console.ExitIf(err)
	validate.SetTagName("validate")
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("name")
		if name != "" {
			return name
		}
		return field.Name
	})
}

// ProcessErr handling error messages
func ProcessErr(s interface{}, err error) map[string]string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	var (
		msgMap = make(map[string]string)
		ref    = reflect.TypeOf(s)
	)

	for _, e := range errs {
		key, msg := "", ""
		field, ok := ref.FieldByName(e.StructField())
		if ok {
			key = field.Tag.Get("json")
			message := processMessage(field.Tag.Get("message"))
			msg = message[e.Tag()]
		}
		if key == "" {
			key = e.StructField()
		}
		if msg == "" {
			msg = e.Translate(trans)
		}

		msgMap[key] = msg
	}
	return msgMap
}

// process message tag into a map
// example map[tag:message]
func processMessage(s string) map[string]string {
	if len(s) == 0 {
		return nil
	}
	var message = make(map[string]string)
	msgArr := strings.Split(s, ",")
	for _, msg := range msgArr {
		if !strings.Contains(msg, ":") {
			continue
		}
		m := strings.Split(msg, ":")
		message[m[0]] = m[1]
	}
	return message
}
