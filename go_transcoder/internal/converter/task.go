package converter

type VideoConverter struct {
}

type VideoTask struct {
	VideoID int    `json:"video_id"`
	Path    string `json:"path"`
}
