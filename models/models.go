package models

type Artist struct {
	Id                int      `json:"id"`
	Image             string   `json:"image"`
	Name              string   `json:"name"`
	Members           []string `json:"members"`
	CreationDate      int      `json:"creationDate"`
	FirstAlbum        string   `json:"firstAlbum"`
	RelationAPI       string   `json:"relations"`
	RelationPerArtist map[string][]string
}

type Error struct {
	ErrorText string
	ErrorCode int
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var (
	Artists     []Artist
	ArtistOne   Artist
	RelationOne Relation
)
