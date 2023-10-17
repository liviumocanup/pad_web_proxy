package models

type Playlist struct {
	ID     uint            `gorm:"primaryKey;auto_increment" db:"id"`
	Name   string          `gorm:"type:varchar(256);not_null" db:"name"`
	UserID string          `gorm:"type:varchar(256);not_null" db:"user_id"`
	Tracks []PlaylistTrack `gorm:"foreignKey:PlaylistID"`
}

type PlaylistTrack struct {
	ID         uint   `gorm:"primaryKey;auto_increment"`
	TrackID    string `db:"track_id"`
	PlaylistID uint
	TrackMetadata
}

type CreatePlaylistRequest struct {
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type PlaylistIdRequest struct {
	PlaylistId string `json:"playlistId"`
}

type RemovePlaylistRequest struct {
	PlaylistId string `json:"playlistId"`
	UserId     string `json:"userId"`
}

type AddTracksToPlaylistRequest struct {
	PlaylistId string   `json:"playlistId"`
	TrackIds   []string `json:"trackIds"`
	UserId     string   `json:"userId"`
}

type RemoveTracksFromPlaylistRequest struct {
	PlaylistId string   `json:"playlistId"`
	TrackIds   []string `json:"trackIds"`
	UserId     string   `json:"userId"`
}

type PlaylistResponse struct {
	PlaylistId string          `json:"playlistId"`
	Name       string          `json:"name"`
	Tracks     []TrackMetadata `json:"tracks"`
}

type TrackMetadata struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	UserID string `json:"userId"`
	URL    string `json:"url"`
}

type PlayPlaylistResponse struct {
	PlaylistName string              `json:"playlistName"`
	Tracks       []TrackPlayMetadata `json:"tracks"`
}

type TrackPlayMetadata struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
