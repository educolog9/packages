package validations

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate
var Uni *ut.UniversalTranslator

func init() {
	// Initialize the validator and register the custom URL validation function
	Validate = validator.New()
	_ = Validate.RegisterValidation("customurl", validateURL)
	_ = Validate.RegisterValidation("mongoID", validateMongoID)

	// Initialize the translators
	en := en.New()
	es := es.New()

	Uni = ut.New(en, es)

	// Set the translator for English
	transEn, _ := Uni.GetTranslator("en")
	registerENTranslations(transEn)

	// Set the translator for Spanish
	transEs, _ := Uni.GetTranslator("es")
	registerESTranslations(transEs)
}

// registerENTranslations registers custom translations for validation tags in English.
// It takes a `trans` ut.Translator as a parameter and adds translations for the "customurl" and "mongoID" validation tags.
// The translations are added using the `Validate.RegisterTranslation` function, which takes a tag name, translator, and translation functions.
// The translation functions are responsible for adding the translated error message for each validation tag.
// The "customurl" translation function adds the error message "{0} is not a valid URL" for the "customurl" tag.
// The "mongoID" translation function adds the error message "{0} is not a valid Mongo ID" for the "mongoID" tag.
// This function does not return any value.
func registerENTranslations(trans ut.Translator) {
	_ = Validate.RegisterTranslation("customurl", trans, func(ut ut.Translator) error {
		return ut.Add("customurl", "The field {0} is not a valid URL", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("customurl", fe.Field())
		return t
	})

	// Add translations for other validation tags here
	_ = Validate.RegisterTranslation("mongoID", trans, func(ut ut.Translator) error {
		return ut.Add("mongoID", "The field {0} is not a valid Mongo ID", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mongoID", fe.Field())
		return t
	})

	// Add translations for the "required" tag
	_ = Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "The field {0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// Add translations for the "min" and "max" tags
	_ = Validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "The field {0} must be at least {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})

	_ = Validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "The field {0} must be no more than {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})

	_ = Validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "The field {0} must be a valid email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
}

// registerESTranslations registers custom translations for validation tags in the provided translator.
// It adds translations for the "customurl" and "mongoID" validation tags.
// The translations are used to provide localized error messages for validation failures.
// The translator parameter should implement the ut.Translator interface.
// The function returns an error if there was a problem registering the translations.
func registerESTranslations(trans ut.Translator) {
	_ = Validate.RegisterTranslation("customurl", trans, func(ut ut.Translator) error {
		return ut.Add("customurl", "El campo {0} no es una URL válida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("customurl", fe.Field())
		return t
	})

	// Add translations for other validation tags here
	_ = Validate.RegisterTranslation("mongoID", trans, func(ut ut.Translator) error {
		return ut.Add("mongoID", "El campo {0} no es un ID de Mongo válido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mongoID", fe.Field())
		return t
	})

	// Add translations for the "required" tag
	_ = Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "El campo {0} es requerido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// Add translations for the "min" and "max" tags
	_ = Validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "El campo {0} debe ser al menos {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})

	_ = Validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "El campo {0} no debe ser más de {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})

	_ = Validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} debe ser un correo electrónico válido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
}
