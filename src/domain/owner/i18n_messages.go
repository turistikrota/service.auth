package owner

type messages struct {
	OwnerNotFound string
	OwnerFailed   string
}

var I18nMessages = messages{
	OwnerNotFound: "owner_not_found",
	OwnerFailed:   "owner_failed",
}
