package database

type Artist struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Record struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Year      *int     `json:"year"`
	Genre     *string  `json:"genre"`
	Support   *string  `json:"support"`
	NbSupport *int     `json:"nb_supports"`
	Label     *string  `json:"label"`
	Tracks    []*Track `json:"tracks"`
}

type Track struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Length *int   `json:"length"`
}
