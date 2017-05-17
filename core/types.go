package core

type Experience struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	From        User   `json:"from"`
	Name        string `json:"name"`
	With        []User `json:"with"`
}

type Page struct {
	ID                 string                `json:"id"`
	About              string                `json:"about"`
	AdminPageNotes     []AdminPageNote       `json:"admin_notes"`
	AgeRange           AgeRange              `json:"age_range"`
	Birthday           string                `json:"birthday"`
	MeasurementRequest bool                  `json:"can_review_measurement_request"`
	Context            UserContext           `json:"context"`
	Cover              CoverPhoto            `json:"cover"`
	Currency           Currency              `json:"currency"`
	Devices            []Device              `json:"devices"`
	Education          []educationExperience `json:"education"`
	Email              string                `json:"email"`
	EmployeeNumber     string                `json:"employee_number"`
}

type User struct {
}

type Currency struct {
	Offset          uint32  `json:"currency_offset"`
	Exchange        float32 `json:"usd_exchange"`
	ExchangeInverse float32 `json:"usd_exchange_inverse"`
	UserCurrency    string  `json:"user_currency"`
}

type Device struct {
	Hardware string `json:"hardware"`
	OS       string `json:"os"`
}

type AgeRange struct {
	Min uint32 `json:"min"`
	Max uint32 `json:"max"`
}

type CoverPhoto struct {
	ID      string  `json:"id"`
	OffsetX float32 `json:"offset_x"`
	OffsetY float32 `json:"offset_y"`
	Source  string  `json:"source"`
}

type UserContext struct {
	ID string `json:"id"`
}

type AdminPageNote struct {
	Body string `json:"body"`
	From Page   `json:"from"`
	ID   string `json:"string"`
	User User   `json:"user"`
}
