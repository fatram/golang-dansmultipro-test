package model

type Position struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}

type PositionFilter struct {
	PageSize    *int   `query:"limit"`
	PageNumber  *int   `query:"page"`
	Description string `query:"description"`
	Location    string `query:"location"`
	FullTime    *bool  `query:"full_time"`
}
