package wishPost

type WishPostAnnouncement struct {
	Url          string   `json:"url"`
	Publish_time string   `json:"publish_time"`
	Title        string   `json:"title"`
	Text         []string `json:"text"`
}

type WishPostPaginationNext struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type WishPostPagination struct {
	Total    int                    `json:"total"`
	Next     WishPostPaginationNext `json:"next"`
	Previous WishPostPaginationNext `json:"previous"`
}

type WishPostAnnouncementData struct {
	Announcements []WishPostAnnouncement `json:"announcements"`
	Pagination    WishPostPagination     `json:"pagination"`
}

type WishPostAnnouncementResponse struct {
	BaseResponse
	Data       WishPostAnnouncementData `json:"data"`
	Extra_data interface{}              `json:"extra_data"`
	Start_date string                   `json:"start_date"`
	End_date   string                   `json:"end_date"`
}
