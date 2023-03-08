package cerror

import "github.com/go-playground/validator/v10"

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required" + fe.Param()
	case "min":
		return "Check min number" + fe.Param()
	case "startswith":
		return "Check startWith " + fe.Param()
	}

	return "Unknown error"
}
