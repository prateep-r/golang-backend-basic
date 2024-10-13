package validator

import (
	"context"
	validatorlib "github.com/go-playground/validator/v10"
)

var validator *validatorlib.Validate

func init() {
	validator = validatorlib.New()
}

func Validate(ctx context.Context, s interface{}) error {
	return validator.StructCtx(ctx, s)
}
