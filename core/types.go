package core

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

type Engagement struct {
	Count                     uint32 `json:"count"`
	CountString               string `json:"count_string"`
	CountStringWithLike       string `json:"count_string_with_like"`
	CountStringWithoutLike    string `json:"count_string_without_like"`
	SocialSentence            string `json:"social_sentence"`
	SocialSentenceWithLike    string `json:"social_sentence_with_like"`
	SocialSentenceWithoutLike string `json:"social_sentence_without_like"`
}

type Location struct {
	City        string  `json:"city"`
	CityID      uint32  `json:"city_id"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Latitude    float32 `json:"latitude"`
	LocatedIn   string  `json:"located_in"`
	Longtitude  float32 `json:"longtitude"`
	Name        string  `json:"name"`
	Region      string  `json:"region"`
	RegionID    uint32  `json:"region_id"`
	State       string  `json:"state"`
	Street      string  `json:"street"`
	Zip         string  `json:"zip"`
}

type PageParking struct {
	Lot    uint32 `json:"lot"`
	Street uint32 `json:"street"`
	Valet  uint32 `json:"valet"`
}

type PagePaymentOptions struct {
	Amex       uint32 `json:"amex"`
	CashOnly   uint32 `json:"cash_only"`
	Discover   uint32 `json:"discover"`
	Mastercard uint32 `json:"mastercard"`
	Visa       uint32 `json:"visa"`
}

type Targeting struct {
	UserOS             []string `json:"user_os"`
	UserDevice         []string `json:"user_device"`
	ExcludedUserDevice []string `json:"excluded_user_device"`
	WirelessCarier     []string `json:"wireless_carier"`
}

type Page struct {
	ID                           string             `json:"id"`
	About                        string             `json:"about"`
	AccessToken                  string             `json:"access_token"`
	AdCampagin                   AdCampagin         `json:"ad_campagin"`
	Affiliation                  string             `json:"affiliation"`
	AppID                        string             `json:"app_id"`
	AppLinks                     AppLinks           `json:"app_links"`
	ArtistsWeLike                string             `json:"artists_we_like"`
	Attire                       string             `json:"attire"`
	Awards                       string             `json:"awards"`
	BandInterests                string             `json:"band_interests"`
	BandMembers                  string             `json:"band_members"`
	BestPage                     Page               `json:"best_page"`
	BIO                          string             `json:"bio"`
	Birthday                     string             `json:"birthday"`
	BookingAgent                 string             `json:"booking_agent"`
	Built                        string             `json:"built"`
	Business                     interface{}        `json:"business"`
	CanCheckIn                   bool               `json:"can_check_in"`
	CanPost                      bool               `json:"can_post"`
	Category                     string             `json:"category"`
	CategoryList                 []PageCategory     `json:"category_list"`
	Checkins                     uint32             `json:"checkins"`
	CompanyOverview              string             `json:"company_overview"`
	ContactAddress               MailingAddress     `json:"contact_address"`
	Context                      OpenGraphContext   `json:"context"`
	CountryPageLikes             uint32             `json:"country_page_likes"`
	Cover                        CoverPhoto         `json:"cover"`
	CulinaryTeam                 string             `json:"culinary_team"`
	CurrentLocation              string             `json:"current_location"`
	Description                  string             `json:"description"`
	DescriptionHTML              string             `json:"description_html"`
	DirectedBy                   string             `json:"directed_by"`
	DisplaySubtext               string             `json:"display_subtext"`
	DisplayedMessageResponseTime string             `json:"displayed_message_response_time"`
	Emails                       []string           `json:"emails"`
	Engagement                   Engagement         `json:"engagement"`
	FanCount                     uint32             `json:"fan_count"`
	FeaturedVideo                Video              `json:"featured_video"`
	Features                     string             `json:"features"`
	FoodStyles                   []string           `json:"food_styles"`
	Founded                      string             `json:"founded"`
	GeneralInfo                  string             `json:"general_info"`
	GeneralManager               string             `json:"general_manager"`
	Genre                        string             `json:"genre"`
	GlobalBrandPageName          string             `json:"global_brand_page_name"`
	GlobalBrandRootID            string             `json:"global_brand_root_id"`
	HasAddedApp                  bool               `json:"has_added_app"`
	Hometown                     string             `json:"hometown"`
	Hours                        map[string]string  `json:"hours"`
	Impressum                    string             `json:"impressum"`
	Influences                   string             `json:"influences"`
	InstantArticlesReviewStatus  string             `json:"instant_articles_review_status"`
	IsAlwaysOpen                 bool               `json:"is_always_open"`
	IsCommunityPage              bool               `json:"is_community_page"`
	IsPermanentlyClosed          bool               `json:"is_permanently_closed"`
	IsPublished                  bool               `json:"is_published"`
	IsUnclaimed                  bool               `json:"is_unclaimed"`
	IsVerified                   bool               `json:"is_verified"`
	IsWebhooksSubscribed         bool               `json:"is_webhooks_subscribed"`
	LeadgenFormPreviewDetails    interface{}        `json:"leadgen_form_preview_details"`
	LeadgenTosAcceptanceTime     string             `json:"leadgen_tos_acceptance_time"`
	LeadgenTosAccepted           bool               `json:"leadgen_tos_accepted"`
	LeadgenTosAcceptingUser      User               `json:"leadgen_tos_accepting_user"`
	Link                         string             `json:"link"`
	Location                     Location           `json:"location"`
	Members                      string             `json:"members"`
	MerchantID                   string             `json:"merchant_id"`
	MerchantReviewStatus         string             `json:"merchant_review_status"`
	Mission                      string             `json:"mission"`
	MPG                          string             `json:"mpg"`
	Name                         string             `json:"name"`
	NameWithLocationDescriptor   string             `json:"name_with_location_descriptor"`
	Network                      string             `json:"network"`
	NewLikeCount                 uint32             `json:"new_like_count"`
	OfferEligable                bool               `json:"offer_eligable"`
	OverallStarRating            float32            `json:"overall_star_rating"`
	ParentPage                   Page               `json:"parent_page"`
	Parking                      PageParking        `json:"parking"`
	PaymentOptions               PagePaymentOptions `json:"payment_options"`
	PersonalInfo                 string             `json:"personal_info"`
	PersonalInterests            string             `json:"personal_interests"`
	PharmaSafetyInfo             string             `json:"pharma_safety_info"`
	Phone                        string             `json:"phone"`
	PlaceType                    string             `json:"place_type"`
	PlotOutline                  string             `json:"plot_outline"`
	PreferredAudience            Targeting          `json:"preferred_audience"`
	PressContact                 string             `json:"press_contact"`
	PriceRange                   string             `json:"price_range"`
	ProducedBy                   string             `json:"produced_by"`
	Products                     string             `json:"products"`
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
