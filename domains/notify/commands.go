package notify

type SendSpecialEmailCmd struct {
	Email          string `json:"email"`
	Template       string `json:"template"`
	Subject        string `json:"subject"`
	Content        string `json:"content"`
	TemplateParams any    `json:"templateParams"`
	Translate      bool   `json:"translate"`
	Locale         string `json:"locale"`
}

type NotifySendSpecialSmsCmd struct {
	Phone     string `json:"phone"`
	Text      string `json:"text"`
	Locale    string `json:"locale"`
	Translate bool   `json:"translate"`
}
