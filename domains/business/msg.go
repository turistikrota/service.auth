package business

type messages struct {
	BusinessNotFound string
	BusinessFailed   string
}

var I18nMessages = messages{
	BusinessNotFound: "business_not_found",
	BusinessFailed:   "business_failed",
}
