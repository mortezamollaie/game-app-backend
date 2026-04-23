package uservalidator

import (
	"fmt"
	"game-app/dto"
	"game-app/pkg/errmsg"
	"game-app/pkg/richerror"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}

func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,
		// TODO - add 3 to config
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{10,}$")).Error("password is not valid.")),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^09[0-9]{9}$")).Error("phone number is not valid."), validation.By(v.checkPhoneNumberUniqueness)),
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
