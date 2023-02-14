package shared

type DashboardTemplate struct {
	Name string `json:"name" form:"name=name" multipartForm:"name=name"`
	URL  string `json:"url" form:"name=url" multipartForm:"name=url"`
}
