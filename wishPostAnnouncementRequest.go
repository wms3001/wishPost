package wishPost

type WishPostAnnouncementRequest struct {
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
}
