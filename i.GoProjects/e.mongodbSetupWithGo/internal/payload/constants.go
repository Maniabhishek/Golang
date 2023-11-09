package payload

type PostData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Writer      string `json:"writer"`
}
