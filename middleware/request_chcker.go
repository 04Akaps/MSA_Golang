package middleware

import (
	"errors"

	"GO_MSA/cerror"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// v10버전을 안쓰면 CHeckBinding은 동작 x => 왜냐면 구버전이기 떄문에
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func CheckBodyBinding(req interface{}, ctx *gin.Context) []ErrorMsg {
	// interface의 default값은 nil 이기 떄문에
	// 사실 어차피 error가 있다는 뜻이기 떄문에 의미는 없는 코드가 된다.

	if req == nil {
		return []ErrorMsg{{
			Field:   "req is nil",
			Message: "req is nil",
		}}
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// bind 체크를 위한 코드
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), cerror.GetErrorMsg(fe)}
			}
			return out
		}
	}

	return nil
}
