package mqservice

// SendEmail data structure
type SendEmail struct {
	Email      string `json:"email"`
	CC         string `json:"cc"`
	BCC        string `json:"bcc"`
	Content    string `json:"content"`
	Subject    string `json:"subject"`
	Attachment string `json:"attachment"`
}

type RestAPIResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

type PushNotification struct {
	UserId       string       `json:"userId"`
	UserType     string       `json:"userType"`
	Notification Notification `json:"notification"`
	Data         interface{}  `json:"data"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NotificationRestAPIResult struct {
	Meta   MetaRestAPIResult `json:"meta"`
	Data   interface{}       `json:"data"`
	Errors string            `json:"errors"`
}

type MetaRestAPIResult struct {
	Message string `json:"message"`
}
