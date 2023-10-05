package models

type Track struct {
	ID     uint   `gorm:"primaryKey;auto_increment" db:"id"`
	Title  string `gorm:"type:varchar(256);not_null" db:"title"`
	Artist string `gorm:"type:varchar(256);not_null" db:"artist"`
	Album  string `gorm:"type:varchar(256);not_null" db:"album"`
	Genre  string `gorm:"type:varchar(256);not_null" db:"genre"`
	URL    string `gorm:"type:varchar(512)" db:"url"`
}

type TrackMetadata struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
}

type IdRequest struct {
	Id string `json:"id"`
}

type TrackResponse struct {
	TrackId string `json:"trackId"`
	URL     string `json:"url"`
}

type TrackInfoResponse struct {
	TrackId string `json:"trackId"`
	Title   string `json:"title"`
	Artist  string `json:"artist"`
	Album   string `json:"album"`
	Genre   string `json:"genre"`
	URL     string `json:"url"`
}

type EditTrackRequest struct {
	TrackId  string        `json:"trackId"`
	Metadata TrackMetadata `json:"metadata"`
}
