package uservalidator

import (
	"fmt"
	"game-app/param"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRequest"

	if err := validation.ValidateStruct(&req,
		// TODO - add 3 to config
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 50)),
		validation.Field(&req.Password,
			validation.Required,
			validation.Match(regexp.MustCompile(passwordRegex)).
				Error(errmsg.ErrorMsgPasswordIsNotValid)),
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).
				Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
			validation.By(v.checkPhoneNumberUniqueness)),
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

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)

	if isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			// under layer handle other data in err
			return err
		}

		return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotUnique)
	}

	return nil
}
