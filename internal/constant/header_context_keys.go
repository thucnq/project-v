package constant

// the keys in the fiber context
const (
	SpanKey = "span"
)

const (
	XGapoRole        = "x-gapo-role"
	XGapoUserId      = "x-gapo-user-id"
	XGapoApiKey      = "x-gapo-api-key"
	XGapoLang        = "x-gapo-lang"
	XGapoWorkspaceId = "x-gapo-workspace-id"

	HeaderAuthorization  = "authorization"
	HeaderContentType    = "Content-Type"
	MIMEApplicationJson  = "application/json"
	HeaderTimezoneOffset = "x-timezone-offset"
)

const DefaultLang = "vi"

var SupportLang = map[string]struct{}{"vi": {}, "en": {}}
