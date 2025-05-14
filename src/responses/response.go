package response

type VacanciesResponse struct {
	Items        []Vacancy `json:"items"`
	Found        int       `json:"found"`
	Pages        int       `json:"pages"`
	Page         int       `json:"page"`
	PerPage      int       `json:"per_page"`
	AlternateUrl string    `json:"alternate_url"`
}

type Counters struct {
	Responses      int `json:"responses"`
	TotalResponses int `json:"total_responses"`
}

type Vacancy struct {
	Id      string `json:"id"`
	Premium bool   `json:"premium"`
	Name    string `json:"name"`
	//Department             string      `json:"department"` // ? type
	HasTest                bool        `json:"has_test"`
	ResponseLetterRequired bool        `json:"response_letter_required"`
	Area                   Area        `json:"area"`
	Salary                 Salary      `json:"salary"`
	SalaryRange            SalaryRange `json:"salary_range"`
	Type                   Type        `json:"type"`
	Address                Address     `json:"address"`
	ResponseUrl            string      `json:"response_url"`
	SortPointDistance      string      `json:"sort_point_distance"`
	PublishedAt            string      `json:"published_at"`
	CreatedAt              string      `json:"created_at"`
	Archived               bool        `json:"archived"`
	ApplyAlternateUrl      string      `json:"apply_alternate_url"`
	ShowLogoInSearch       bool        `json:"show_logo_in_search"`
	ShowContacts           bool        `json:"show_contacts"`
	//InsiderInterview       string      `json:"insider_interview"` // ? type
	Url          string `json:"url"`
	AlternateUrl string `json:"alternate_url"`
	//Relations []? `json:"relations"`
	Employer Employer `json:"employer"`
	Snippet  Snippet  `json:"snippet"`
	Contacts Contacts `json:"contacts"`
	Schedule Schedule `json:"schedule"`
	Counters Counters `json:"counters"`
	// `json:"working_days"`
	// `json:"working_time_intervals"`
	// `json:"working_time_modes"`
	AcceptTemporary bool `json:"accept_temporary"`
	// `json:"fly_in_fly_out_duration"`
	WorkFormat              []IdName `json:"work_format"`
	WorkingHours            []IdName `json:"working_hours"`
	WorkScheduleByDays      []IdName `json:"work_schedule_by_days"`
	NightShifts             bool     `json:"night_shifts"`
	ProfessionalRoles       []IdName `json:"professional_roles"`
	AcceptIncompleteResumes bool     `json:"accept_incomplete_resumes"`
	Experience              IdName   `json:"experience"`
	Employment              IdName   `json:"employment"`
	EmploymentForm          IdName   `json:"employment_form"`
	Internship              bool     `json:"internship"`
	// `json:"adv_response_url"`
	IsAdvVacancy bool `json:"is_adv_vacancy"`
	// `json:"adv_context"`
}

type Contacts struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	// `json:"phones"`
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
	// `json:"building"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	// `json:"description"`
	Raw string `json:"raw"`
	// `json:"metro"`
	// `json:"metro_stations"`
	Id string `json:"id"`
}

type IdName struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Area struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Salary struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	Currency string `json:"currency"`
	Gross    bool   `json:"gross"`
}

type SalaryRange struct {
	From      int    `json:"from"`
	To        int    `json:"to"`
	Currency  string `json:"currency"`
	Gross     bool   `json:"gross"`
	Mode      Mode   `json:"mode"`
	Frequency IdName `json:"frequency"`
}

type Mode struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Type struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Employer struct {
	Id                   string         `json:"id"`
	Name                 string         `json:"name"`
	Url                  string         `json:"url"`
	AlternateUrl         string         `json:"alternate_url"`
	LogoUrls             LogoUrls       `json:"logo_urls"`
	VacanciesUrl         string         `json:"vacancies_url"`
	AccreditedItEmployer bool           `json:"accredited_it_employer"`
	EmployerRating       EmployerRating `json:"employer_rating"`
	Trusted              bool           `json:"trusted"`
}

type LogoUrls struct {
	Original string
	// 90 string ?
	// 240 string ?
}

type EmployerRating struct {
	TotalRating  string `json:"total_rating"`
	ReviewsCount int    `json:"reviews_count"`
}

type Snippet struct {
	Requirement    string `json:"requirement"`
	Responsibility string `json:"responsibility"`
}

type Schedule struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AreaRoot struct {
	Id       string `json:"id"`
	ParentId string `json:"parent_id"`
	Name     string `json:"name"`
}