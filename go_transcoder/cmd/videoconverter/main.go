package main

import "github.com/rodolfoHOk/fullcycle.imersao19/go_transcoder/internal/converter"

func main() {
	vc := converter.NewVideoConverter()
	vc.Handle([]byte(`{"video_id": 1, "path": "mediatest/media/uploads/1"}`))
}
