package booking

//BookingModel
type BookingModel struct {
	ID         int    `json:"id" eq:"id"`
	NotetypeID int    `json:"typeid" eq:"typeid"`
	Notetype   string `json:"type" eq:"type"`
	CategoryID int    `json:"categoryid" eq:"categoryid"`
	Category   string `json:"category" eq:"category"`
	Amount     int    `json:"amount" eq:"amount"`
	Note       string `json:"note" eq:"note"`
	CreatedAt  string `json:"created_at" eq:"created_at"`
	UpdatedAt  string `json:"updated_at" eq:"updated_at"`
}
