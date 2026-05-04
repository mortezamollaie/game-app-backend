package uservalidator

import (
	"fmt"
	"game-app/dto"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateLoginRequest(req dto.RegisterRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).
				Error("phone number is not valid."),
			validation.By(v.doesPhoneNumberUniqueness),
		),
		validation.Field(&req.Password, validation.Required),
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

func (v Validator) doesPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)

	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return fmt.Errorf(errmsg.ErrorMsgNotFound)
	}

	return nil
}
