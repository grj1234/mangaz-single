package meta

type Data struct {
	Location Location          `json:"Location"`
	Orders   []Order           `json:"Orders"`
	Book     Book              `json:"Book"`
	Authors  map[string]Author `json:"Authors"`
}
type Author struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
	Name     string `json:"name"`
}
type Book struct {
	Baid          int      `json:"baid"`
	SeriesID      int      `json:"series_id"`
	Status        string   `json:"status"`
	Reading       int      `json:"reading"`
	UploadedBy    *string  `json:"uploaded_by"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	SiteTarget    int      `json:"site_target"`
	R18           int      `json:"r18"`
	Volume        string   `json:"volume"`
	URL           string   `json:"url"`
	CoverImage    string   `json:"cover_image"`
	Publication   string   `json:"publication"`
	PdfSale       int      `json:"pdf_sale"`
	VwShare       int      `json:"vw_share"`
	Category      string   `json:"category"`
	Rating        int      `json:"rating"`
	Premium       int      `json:"premium"`
	Pages         int      `json:"pages"`
	TrialPages    *int     `json:"trial_pages"`
	Authors       []string `json:"Authors"`
	PageLayout    string   `json:"page_layout"`
	PageDirection int      `json:"page_direction"`
	PageMaxWidth  int      `json:"page_max_width"`
	PageMaxHeight int      `json:"page_max_height"`
	ImageCount    int      `json:"image_count"`
}
type Location struct {
	Base         string `json:"base"`
	Scramble_dir string `json:"scramble_dir"`
}
type Order struct {
	No       int    `json:"no"`
	Name     string `json:"name"`
	Side     string `json:"side"`
	URL      string
	Scramble *Scramble `json:"scramble,omitempty"`
	ImgRaw   []byte    `json:"-"`
}
type Scramble struct {
	W     int    `json:"w"`
	H     int    `json:"h"`
	Crops []Crop `json:"crops"`
}
type Crop struct {
	X  int `json:"x"`
	Y  int `json:"y"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
	W  int `json:"w"`
	H  int `json:"h"`
}
