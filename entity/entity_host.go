package entity

type Host struct {
	Data    []Hosts `json:"data"`
	Message string  `json:"message"`
	Status  int64   `json:"status"`
}

type Hosts struct {
	ID            int64   `json:"id"`
	CompanyID     int64   `json:"company_id"`
	HostName      string  `json:"host_name"`
	FrontendURL   string  `json:"frontend_url"`
	BackendURL    string  `json:"backend_url"`
	ViewFolder    string  `json:"view_folder"`
	Status        int64   `json:"status"`
	DefaultRoleID int64   `json:"default_role_id"`
	DefaultAppID  int64   `json:"default_app_id"`
	Company       Company `json:"company"`
}

type Company struct {
	ID                int64  `json:"id"`
	CompanyName       string `json:"company_name"`
	Logo              string `json:"logo"`
	Color             string `json:"color"`
	CompanyAddress    string `json:"company_address"`
	CompanyPhone      string `json:"company_phone"`
	CompanyEmail      string `json:"company_email"`
	CompanyWebsite    string `json:"company_website"`
	PicName           string `json:"pic_name"`
	PicPhone          string `json:"pic_phone"`
	AppID             int64  `json:"app_id"`
	ContractStartDate string `json:"contract_start_date"`
	ContractEndDate   string `json:"contract_end_date"`
	ContractValue     int64  `json:"contract_value"`
	Impression        int64  `json:"impression"`
	Token             int64  `json:"token"`
	Cpm               int64  `json:"cpm"`
	IncludeDapil      bool   `json:"include_dapil"`
	IncludeDisclaimer bool   `json:"include_disclaimer"`
	UserQuota         int64  `json:"user_quota"`
	PicEmail          string `json:"pic_email"`
	PicUserID         int64  `json:"pic_user_id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
