package structures

// Container struct
type Container struct {
	ID     string   `json:"id"`
	Image  string   `json:"image"`
	Names  []string `json:"names"`
	Status string   `json:"status"`
}
