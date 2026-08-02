package matchingvalidator

import (
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddToWaitingListRequest(req param.AddToWaitingListRequest) (map[string]string, error) {
	const op = "uservalidator.AddToWaitingList"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category,
			validation.Required,
			validation.By(v.isCategoryValid),
		),
	); err != nil {
		errV, ok := err.(validation.Errors)
		fieldErrors := map[string]string{}

		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}

	return nil, nil
}

func (v Validator) isCategoryValid(value interface{}) error {
	category := value.(entity.Category)

	if category.IsValid() {
		return fmt.Errorf(errmsg.ErrorMsgCategoryIsNotValid)
	}

	return nil
}
