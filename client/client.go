package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MetMuseumResponseObject struct {
	ObjectId              int               `json:"objectID"`
	IsHighlight           bool              `json:"isHighlight"`
	AccessionNumber       string            `json:"accessionNumber"`
	AccessionYear         string            `json:"accessionYear"`
	IsPublicDomain        bool              `json:"isPublicDomain"`
	PrimaryImage          string            `json:"primaryImage"`
	PrimaryImageSmall     string            `json:"primaryImageSmall"`
	AdditionalImages      []AdditionalImage `json:"additionalImages"`
	Constituents          []Constituent     `json:"constituents"`
	Department            string            `json:"department"`
	ObjectName            string            `json:"objectName"`
	Title                 string            `json:"title"`
	Culture               string            `json:"culture"`
	Period                string            `json:"period"`
	Dynasty               string            `json:"dynasty"`
	Reign                 string            `json:"reign"`
	Portfolio             string            `json:"portfolio"`
	ArtistRole            string            `json:"artistRole"`
	ArtistPrefix          string            `json:"artistPrefix"`
	ArtistDisplayName     string            `json:"artistDisplayName"`
	ArtistDisplayBio      string            `json:"artistDisplayBio"`
	ArtistSuffix          string            `json:"artistSuffix"`
	ArtistAlphaSort       string            `json:"artistAlphaSort"`
	ArtistNationality     string            `json:"artistNationality"`
	ArtistBeginDate       string            `json:"artistBeginDate"`
	ArtistEndDate         string            `json:"artistEndDate"`
	ArtistGender          string            `json:"artistGender"`
	ArtistWikidata_URL    string            `json:"artistWikidata_URL"`
	ArtistULAN_URL        string            `json:"artistULAN_URL"`
	ObjectDate            string            `json:"objectDate"`
	ObjectBeginDate       int               `json:"objectBeginDate"`
	ObjectEndDate         int               `json:"objectEndDate"`
	Medium                string            `json:"medium"`
	Dimensions            string            `json:"dimensions"`
	DimensionsParsed      float64           `json:"dimensionsParsed"`
	Measurements          []Measurement     `json:"measurements"`
	CreditLine            string            `json:"creditLine"`
	GeographyType         string            `json:"geographyType"`
	City                  string            `json:"city"`
	State                 string            `json:"state"`
	County                int               `json:"county"`
	Country               int               `json:"country"`
	Region                string            `json:"region"`
	Subregion             string            `json:"subregion"`
	Locale                string            `json:"locale"`
	Locus                 int               `json:"locus"`
	Excavation            int               `json:"excavation"`
	River                 string            `json:"river"`
	Classification        int               `json:"classification"`
	RightsAndReproduction string            `json:"rightsAndReproduction"`
	LinkResource          string            `json:"linkResource"`
	MetadataDate          string            `json:"metadataDate"`
	Repository            int               `json:"repository"`
	ObjectURL             int               `json:"objectURL"`
	Tags                  struct {
		Term         string `json:"term"`
		AAT_URL      string `json:"AAT_URL"`
		Wikidata_URL string `json:"Wikidata_URL"`
	} `json:"tags"`
	ObjectWikidata_URL string `json:"objectWikidata_URL"`
	IsTimelineWork     bool   `json:"isTimelineWork"`
	GalleryNumber      string `json:"GalleryNumber"`
}

type AdditionalImage struct {
	Url string
}

type Constituent struct {
	ConstituentId       int    `json:"constituentID"`
	Role                string `json:"role"`
	Name                string `json:"name"`
	ConstituentULAN_URL string `json:"constituentULAN_URL"`
}

type Measurement struct {
	ElementName         string `json:"elementName"`
	ElementDescription  string `json:"elementDescription"`
	ElementMeasurements struct {
		Height float64 `json:"Height"`
		Length float64 `json:"Length"`
		Width  float64 `json:"Width"`
	} `json:"elementMeasurements"`
}

type MetMuseumResponse struct {
	MetMuseumResponseObjects []MetMuseumResponseObject
}

type ObjectIdResponse struct {
	Total     int   `json:"total"`
	ObjectIds []int `json:"objectIDs"`
}

func GetMetMuseumObjects(objectId int, c *http.Client) (obj *MetMuseumResponse, err error) {

	requestUrl := fmt.Sprintf("https://collectionapi.metmuseum.org/public/collection/v1/objects/%d", objectId)

	req, err := http.NewRequest("GET", requestUrl, nil)

	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	resp, err := c.Do(req)

	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(obj)

	if err != nil {
		log.Println("Error parsing response body:", err)
		return nil, err
	}

	return obj, nil
}

func GetObjectIds(medium string, c *http.Client) (*ObjectIdResponse, error) {
	requestUrl := fmt.Sprintf("https://collectionapi.metmuseum.org/public/collection/v1/search?hasImages=true&medium=%s", medium)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	var obj *ObjectIdResponse
	err = json.NewDecoder(resp.Body).Decode(obj)

	if err != nil {
		log.Println("Error parsing response body:", err)
		return nil, err
	}

	return obj, nil
}
