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
	_ = Validate.RegisterValidation("customArrayURL", validateURLArray)
	_ = Validate.RegisterValidation("picture", validatePictureURL)
	_ = Validate.RegisterValidation("pictures", validatePicturesURL)
	_ = Validate.RegisterValidation("mongoID", validateMongoID)
	_ = Validate.RegisterValidation("latitude", validateLatitude)
	_ = Validate.RegisterValidation("longitude", validateLongitude)
	_ = Validate.RegisterValidation("phone", validatePhoneNumber)
	_ = Validate.RegisterValidation("countryISOCode", validateCountryISOCode)
	_ = Validate.RegisterValidation("mongoIDs", validateMongoIDs)

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

	_ = Validate.RegisterTranslation("customArrayURL", trans, func(ut ut.Translator) error {
		return ut.Add("urlarray", "The field {0} contains an invalid URL", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("urlarray", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("picture", trans, func(ut ut.Translator) error {
		return ut.Add("picture", "The field {0} is not a valid picture URL", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("picture", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("pictures", trans, func(ut ut.Translator) error {
		return ut.Add("pictures", "The field {0} contains an invalid picture URL", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("pictures", fe.Field())
		return t
	})

	// Add translations for other validation tags here
	_ = Validate.RegisterTranslation("mongoID", trans, func(ut ut.Translator) error {
		return ut.Add("mongoID", "The field {0} is not a valid Mongo ID", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mongoID", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("mongoIDs", trans, func(ut ut.Translator) error {
		return ut.Add("mongoIDs", "The field {0} contains an invalid Mongo ID", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mongoIDs", fe.Field())
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

	_ = Validate.RegisterTranslation("latitude", trans, func(ut ut.Translator) error {
		return ut.Add("latitude", "The field {0} must be a valid latitude", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("latitude", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("longitude", trans, func(ut ut.Translator) error {
		return ut.Add("longitude", "The field {0} must be a valid longitude", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("longitude", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "The field {0} must be a valid phone number", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("countryISOCode", trans, func(ut ut.Translator) error {
		return ut.Add("countryISOCode", "The field {0} must be a valid country ISO code", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("countryISOCode", fe.Field())
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

	_ = Validate.RegisterTranslation("customArrayURL", trans, func(ut ut.Translator) error {
		return ut.Add("urlarray", "El campo {0} contiene una URL inválida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("urlarray", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("picture", trans, func(ut ut.Translator) error {
		return ut.Add("picture", "El campo {0} no es una URL de imagen válida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("picture", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("pictures", trans, func(ut ut.Translator) error {
		return ut.Add("pictures", "El campo {0} contiene una URL de imagen inválida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("pictures", fe.Field())
		return t
	})

	// Add translations for other validation tags here
	_ = Validate.RegisterTranslation("mongoID", trans, func(ut ut.Translator) error {
		return ut.Add("mongoID", "El campo {0} no es un ID de Mongo válido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mongoID", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("mongoIDs", trans, func(ut ut.Translator) error {
		return ut.Add("mongoIDs", "El campo {0} contiene un ID de Mongo inválido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mongoIDs", fe.Field())
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

	_ = Validate.RegisterTranslation("latitude", trans, func(ut ut.Translator) error {
		return ut.Add("latitude", "El campo {0} debe ser una latitud válida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("latitude", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("longitude", trans, func(ut ut.Translator) error {
		return ut.Add("longitude", "El campo {0} debe ser una longitud válida", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("longitude", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "El campo {0} debe ser un número de teléfono válido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("countryISOCode", trans, func(ut ut.Translator) error {
		return ut.Add("countryISOCode", "El campo {0} debe ser un código ISO de país válido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("countryISOCode", fe.Field())
		return t
	})
}
