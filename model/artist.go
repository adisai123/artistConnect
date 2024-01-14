package model

type Artists struct {
	ID             uint8  `gorm:"primaryKey"`
	Name           string `json:"name"`
	Subject        string `json:"subject"`
	Address        string `json:"address"`
	Profile        string `json:"profile"`
	Charges        uint16 `json:"charges"`
	Note           string `json:"note"`
	YoutubeChannel string `json:"youtubeChannel"`
	LatestLink     string `json:"latestLink"`
}
