package config

type MongoAuth struct {
	Host       string `env:"MONGO_AUTH_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_AUTH_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_AUTH_USERNAME" envDefault:""`
	Password   string `env:"MONGO_AUTH_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_AUTH_DATABASE" envDefault:"empty"`
	Collection string `env:"MONGO_AUTH_COLLECTION" envDefault:"empties"`
	Query      string `env:"MONGO_AUTH_QUERY" envDefault:""`
}

type Turnstile struct {
	Secret       string `env:"CF_TURNSTILE_SECRET_KEY"`
	MobileSecret string `env:"CF_TURNSTILE_MOBILE_SECRET_KEY"`
	Skip         bool   `env:"TURNSTILE_SKIP_AUTH"`
}

type Rpc struct {
	AccountHost     string `env:"RPC_ACCOUNT_HOST" envDefault:"localhost:3001"`
	AccountUsesSsl  bool   `env:"RPC_ACCOUNT_USES_SSL" envDefault:"localhost:3001"`
	BusinessHost    string `env:"RPC_BUSINESS_HOST" envDefault:"localhost:3002"`
	BusinessUsesSsl bool   `env:"RPC_BUSINESS_USES_SSL" envDefault:"localhost:3002"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type MongoAccount struct {
	Collection string `env:"MONGO_ACCOUNT_COLLECTION" envDefault:"empty"`
}

type MongoBusiness struct {
	Collection string `env:"MONGO_BUSINESS_COLLECTION" envDefault:"empty"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type CsrfRedis struct {
	Host string `env:"REDIS_CSRF_HOST"`
	Port int    `env:"REDIS_CSRF_PORT"`
	Pw   string `env:"REDIS_CSRF_PASSWORD"`
	Db   int    `env:"REDIS_CSRF_DB"`
}

type CsrfBaseEnv struct {
	SameSite   string `env:"CSRF_SAME_SITE"`
	HttpOnly   bool   `env:"CSRF_HTTP_ONLY"`
	Secure     bool   `env:"CSRF_SECURE"`
	Domain     string `env:"CSRF_DOMAIN"`
	Expiration int    `env:"CSRF_EXPIRATION"`
}

type Http struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"auth"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type Topics struct {
	Auth   AuthTopics
	Notify NotifyTopics
	Verify VerifyTopics
	Admin  AdminTopics
}

type Urls struct {
	Check2FA   string `env:"URL_CHECK_2FA" envDefault:"http://localhost:3000/auth/2fa/check"`
	VerifyMail string `env:"URL_VERIFY_MAIL" envDefault:"http://localhost:3000/auth/verify"`
}

type AuthTopics struct {
	// publishers
	Base          string `env:"STREAMING_TOPIC_AUTH_BASE"`
	Registered    string `env:"STREAMING_TOPIC_AUTH_REGISTERED"`
	LoggedIn      string `env:"STREAMING_TOPIC_AUTH_LOGGED_IN"`
	LoginFailed   string `env:"STREAMING_TOPIC_AUTH_LOGIN_FAILED"`
	TokenExtended string `env:"STREAMING_TOPIC_AUTH_TOKEN_EXTENDED"`
	UserVerified  string `env:"STREAMING_TOPIC_AUTH_USER_VERIFIED"`

	// listeners
	LoginVerified string `env:"STREAMING_TOPIC_AUTH_LOGIN_VERIFIED"`
	UserUpdated   string `env:"STREAMING_TOPIC_AUTH_USER_UPDATED"`
}

type NotifyTopics struct {
	SendEmailToActor string `env:"STREAMING_TOPIC_NOTIFY_SEND_EMAIL_TO_ACTOR"`
	SendSmsToActor   string `env:"STREAMING_TOPIC_NOTIFY_SEND_SMS_TO_ACTOR"`
	SendSpecialEmail string `env:"STREAMING_TOPIC_NOTIFY_SEND_SPECIAL_EMAIL"`
	SendSpecialSms   string `env:"STREAMING_TOPIC_NOTIFY_SEND_SPECIAL_SMS"`
	SendNotification string `env:"STREAMING_TOPIC_NOTIFY_SEND_NOTIFICATION"`
	SendPush         string `env:"STREAMING_TOPIC_NOTIFY_SEND_PUSH"`
}

type AdminTopics struct {
	PermissionsAdded   string `env:"STREAMING_TOPIC_ADMIN_PERMISSIONS_ADDED"`
	PermissionsRemoved string `env:"STREAMING_TOPIC_ADMIN_PERMISSIONS_REMOVED"`
}

type VerifyTopics struct {
	Start2FA string `env:"STREAMING_TOPIC_START_2FA"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		Auth     MongoAuth
		Account  MongoAccount
		Business MongoBusiness
	}
	Rpc         Rpc
	Redis       Redis
	Http        Http
	HttpHeaders HttpHeaders
	I18n        I18n
	Topics      Topics
	Nats        Nats
	Session     Session
	CacheRedis  CacheRedis
	TokenSrv    TokenSrv
	Urls        Urls
	Turnstile   Turnstile
	CSRF        struct {
		BaseEnv CsrfBaseEnv
		Redis   CsrfRedis
	}
	RSA RSA
}
