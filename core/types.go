package core

type Page struct {
}

type User struct {
	ID                        string                `json:"id"`
	About                     string                `json:"about"`
	AdminPageNotes            []AdminPageNote       `json:"admin_notes"`
	AgeRange                  AgeRange              `json:"age_range"`
	Birthday                  string                `json:"birthday"`
	MeasurementRequest        bool                  `json:"can_review_measurement_request"`
	Context                   UserContext           `json:"context"`
	Cover                     CoverPhoto            `json:"cover"`
	Currency                  Currency              `json:"currency"`
	Devices                   []Device              `json:"devices"`
	Education                 []educationExperience `json:"education"`
	Email                     string                `json:"email"`
	EmployeeNumber            string                `json:"employee_number"`
	FavoriteAthletes          []Experience          `json:"favorite_athletes"`
	FavoriteTeams             []Experience          `json:"favorite_teams"`
	FirstName                 string                `json:"first_name"`
	Gender                    string                `json:"gender"`
	Hometown                  Page                  `json:"hometown"`
	InspirationalPeople       []Experience          `json:"inspirational_people"`
	InstallType               string                `json:"install_type"`
	Installed                 bool                  `json:"installed"`
	InterestedIn              []string              `json:"interested_in"`
	IsSharedLogin             bool                  `json:"is_shared_login"`
	IsVerified                bool                  `json:"is_verified"`
	Labels                    []PageLabel           `json:"labels"`
	Languages                 []Experience          `json:"languages"`
	LastName                  string                `json:"last_name"`
	Link                      string                `json:"link"`
	Locale                    string                `json:"locale"`
	Location                  Page                  `json:"location"`
	MeetingFor                []string              `json:"meeting_for"`
	MiddleName                string                `json:"middle_name"`
	Name                      string                `json:"name"`
	NameFormat                string                `json:"name_format"`
	PaymentPricepoints        PaymentPricepoints    `json:"payment_pricepoints"`
	Political                 string                `json:"political"`
	PublicKey                 string                `json:"public_key"`
	Quotes                    string                `json:"quotes"`
	RelationshipStatus        string                `json:"relationship_status"`
	Religion                  string                `json:"religion"`
	SecuritySettings          SecuritySettings      `json:"security_settings"`
	SharedLoginUpdateRequired string                `json:"shared_login_update_required_by"`
	ShortName                 string                `json:"short_name"`
	SignificantOther          User                  `json:"significant_other"`
	Sports                    []Experience          `json:"sports"`
	TestGroup                 uint32                `json:"test_group"`
	ThirdPartyID              string                `json:"third_party_id"`
	Timezone                  float32               `json:"timezone"`
	TokenForBusiness          string                `json:"token_for_business"`
	UpdateTime                string                `json:"update_time"`
	Verified                  bool                  `json:"verified"`
	VideoUploadLimits         VideoUploadLimits     `json:"video_upload_limits"`
	ViewerCanSendGift         bool                  `json:"viewer_can_send_gift"`
	Website                   string                `json:"website"`
	Work                      []WorkExperience      `json:"work"`
}

type educationExperience struct {
	ID            string `json:"id"`
	Classes       []Experience
	Concentration []Page
	Degree        Page
	School        Page
	With          []User
	Year          Page
	Field         string `json:"field"`
}

type ProjectExperience struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	EndDate     string `json:"end_date"`
	From        User   `json:"from"`
	Name        string `json:"name"`
	StartDate   string `json:"start_date"`
	With        []User `json:"with"`
}

type WorkExperience struct {
	ID          int64               `json:"id"`
	Desctiption string              `json:"desctiption"`
	Employer    Page                `json:"employer"`
	EndDate     string              `json:"end_date"`
	From        User                `json:"from"`
	Location    Page                `json:"location"`
	Position    Page                `json:"position"`
	Projects    []ProjectExperience `json:"projects"`
	StartDate   string              `json:"start_date"`
	With        []User              `json:"with"`
}

type PageLabel struct {
	CreationTime string `json:"creation_time"`
	//TODO A profile can be a: User Page Group Event Application
	CreatorID interface{} `json:"creator_id"`
	From      Page        `json:"from"`
	ID        string      `json:"id"`
	Name      string      `json:"name"`
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

type PaymentPricepoint struct {
	Credits       float32 `json:"credits"`
	LocalCurrency string  `json:"local_currency"`
	UserPrice     string  `json:"user_price"`
}

type PaymentPricepoints struct {
	Mobile []PaymentPricepoint `json:"mobile"`
}

type SecureBrowsing struct {
	Enabled bool `json:"enabled"`
}

type SecuritySettings struct {
	SecureBrowsing SecureBrowsing `json:"secure_browsing"`
}

type VideoUploadLimits struct {
	Length uint32 `json:"length"`
	Size   int    `json:"size"`
}

type Experience struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	From        User   `json:"from"`
	Name        string `json:"name"`
	With        []User `json:"with"`
}
