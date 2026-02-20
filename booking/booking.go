package booking

type Response struct {
	Extensions Extensions `json:"extensions"`
	Data       Data       `json:"data"`
}

type Extensions struct {
	LatencyInsights map[string]LatencyInsight `json:"latency_insights"`
}

type LatencyInsight struct {
	StartTime  int64 `json:"start_time"`
	EndTime    int64 `json:"end_time"`
	DurationMs int   `json:"duration_ms"`
}

type Data struct {
	RoomDetail   RoomDetail         `json:"roomDetail"`
	ReviewScores ReviewScoresResult `json:"reviewScores"`
}

type RoomDetail struct {
	HasRoomsWithDifferentPrivacyLevels bool                           `json:"hasRoomsWithDifferentPrivacyLevels"`
	TypeName                           string                         `json:"__typename"`
	AreaMeasurementUnit                string                         `json:"areaMeasurementUnit"`
	Property                           Property                       `json:"property"`
	HighlightsForAllRooms              []HighlightsForAllRooms        `json:"highlightsForAllRooms"`
	CategorizedFacilitiesForAllRooms   []CategorizedFacilitiesForRoom `json:"categorizedFacilitiesForAllRooms"`
}

type Property struct {
	AccommodationType        AccommodationType `json:"accommodationType"`
	HasDesignatedSmokingArea bool              `json:"hasDesignatedSmokingArea"`
	RoomsDetails             []RoomDetails     `json:"roomsDetails"`
	HighFloorStartsAt        int               `json:"highFloorStartsAt"`
	Name                     string            `json:"name"`
	ID                       int               `json:"id"`
	TypeName                 string            `json:"__typename"`
}

type AccommodationType struct {
	TypeName string `json:"__typename"`
	ID       int    `json:"id"`
	Type     string `json:"type"`
}

type RoomDetails struct {
	RoomFloor                    []interface{}              `json:"roomFloor"`
	BedConfigurations            []BedConfiguration         `json:"bedConfigurations"`
	BathroomCount                int                        `json:"bathroomCount"`
	CribsAvailableForFree        bool                       `json:"cribsAvailableForFree"`
	Occupancy                    RoomOccupancy              `json:"occupancy"`
	ApartmentRooms               []ApartmentRoom            `json:"apartmentRooms"`
	IsSmoking                    bool                       `json:"isSmoking"`
	RoomSizeM2                   float64                    `json:"roomSizeM2"`
	RoomTypeID                   int                        `json:"roomTypeId"`
	RoomPhotos                   []RoomPhoto                `json:"roomPhotos"`
	BathroomFacilityAttributes   BathroomFacilityAttributes `json:"bathroomFacilityAttributes"`
	Translations                 RoomTranslation            `json:"translations"`
	IsBiggerThanAverageRoomInUfi bool                       `json:"isBiggerThanAverageRoomInUfi"`
	TypeName                     string                     `json:"__typename"`
	ID                           int                        `json:"id"`
	PrivacyLevel                 int                        `json:"privacyLevel"`
}

type BedConfiguration struct {
	Beds            []Bed  `json:"beds"`
	TypeName        string `json:"__typename"`
	ConfigurationID int    `json:"configurationId"`
}

type Bed struct {
	TypeName    string             `json:"__typename"`
	Translation BedTypeTranslation `json:"translation"`
	BedType     string             `json:"bedType"`
	Count       int                `json:"count"`
}

type BedTypeTranslation struct {
	TypeName string `json:"__typename"`
	Name     string `json:"name"`
}

type RoomOccupancy struct {
	MaxGuests   int    `json:"maxGuests"`
	TypeName    string `json:"__typename"`
	MaxChildren int    `json:"maxChildren"`
	MaxPersons  int    `json:"maxPersons"`
}

type ApartmentRoom struct {
	Beds            []Bed  `json:"beds"`
	EnsuiteBathroom int    `json:"ensuiteBathroom"`
	Count           int    `json:"count"`
	MaxPersons      int    `json:"maxPersons"`
	ID              int    `json:"id"`
	TypeName        string `json:"__typename"`
	RoomType        string `json:"roomType"`
}

type RoomPhoto struct {
	ThumbnailURI string `json:"thumbnailUri"`
	TypeName     string `json:"__typename"`
	ID           int    `json:"id"`
	PhotoURI     string `json:"photoUri"`
}

type BathroomFacilityAttributes struct {
	IsExternalBathroom bool   `json:"isExternalBathroom"`
	TypeName           string `json:"__typename"`
	IsEnsuiteBathroom  bool   `json:"isEnsuiteBathroom"`
}

type RoomTranslation struct {
	Description string `json:"description"`
	TypeName    string `json:"__typename"`
	Name        string `json:"name"`
}

// Highlights

type HighlightsForAllRooms struct {
	TypeName       string          `json:"__typename"`
	RoomID         int             `json:"roomId"`
	RoomHighlights []RoomHighlight `json:"roomHighlights"`
}

// RoomHighlight is a union type; use interface fields to cover all variants.
type RoomHighlight struct {
	TypeName            string   `json:"__typename"`
	PrivacyLevel        string   `json:"privacyLevel,omitempty"`
	AreaValue           *float64 `json:"areaValue,omitempty"`
	ID                  *int     `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	TranslationOverride *string  `json:"translationOverride,omitempty"`
}

// Categorized Facilities

type CategorizedFacilitiesForRoom struct {
	RoomID                int                      `json:"roomId"`
	CategorizedFacilities []FacilitiesWithCategory `json:"categorizedFacilities"`
	TypeName              string                   `json:"__typename"`
}

type FacilitiesWithCategory struct {
	Facilities []SimpleFacility `json:"facilities"`
	TypeName   string           `json:"__typename"`
	Category   string           `json:"category"`
}

type SimpleFacility struct {
	Name     string `json:"name"`
	TypeName string `json:"__typename"`
	ID       int    `json:"id"`
}

// Review Scores

type ReviewScoresResult struct {
	TypeName     string        `json:"__typename"`
	ReviewScores []ReviewScore `json:"reviewScores"`
}

type ReviewScore struct {
	TypeName     string  `json:"__typename"`
	Name         string  `json:"name"`
	Translation  *string `json:"translation"`
	Count        int     `json:"count"`
	Value        float64 `json:"value"`
	CustomerType string  `json:"customerType"`
}
