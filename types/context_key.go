package types

type contextKey string

// ContentLanguageKey is a context key used to store the content language.
// It is used to retrieve the content language from the context.
const ContentLanguageKey contextKey = "contentLanguage"

// CountryCodeKey is a context key used to store the country code.
// It is used to retrieve the country code from the context.
const CountryCodeKey contextKey = "countryCode"
