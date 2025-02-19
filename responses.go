package cqc

type ProviderItem struct {
	ProviderID string `json:"providerId"`
	Name       string `json:"providerName"`
}

type ProvidersResponse struct {
	FirstPageUri    string         `json:"firstPageUri"`
	LastPageUri     string         `json:"lastPageUri"`
	NextPageUri     string         `json:"nextPageUri"`
	Page            int            `json:"page"`
	PerPage         int            `json:"perPage"`
	PreviousPageUri *string        `json:"previousPageUri"`
	Providers       []ProviderItem `json:"providers"`
	Total           int            `json:"total"`
	TotalPages      int            `json:"totalPages"`
	Uri             string         `json:"uri"`
}

type Provider struct {
	ProviderID            string                 `json:"providerId"`
	LocationIDs           []string               `json:"locationIds"`
	OrganisationType      string                 `json:"organisationType"`
	OwnershipType         string                 `json:"ownershipType"`
	Type                  string                 `json:"type"`
	Name                  string                 `json:"name"`
	BrandID               string                 `json:"brandId"`
	BrandName             string                 `json:"brandName"`
	OdsCode               string                 `json:"odsCode"`
	RegistrationStatus    string                 `json:"registrationStatus"`
	RegistrationDate      string                 `json:"registrationDate"`
	CompaniesHouseNumber  string                 `json:"companiesHouseNumber"`
	CharityNumber         string                 `json:"charityNumber"`
	Website               string                 `json:"website"`
	PostalAddress         PostalAddress          `json:"-"`
	Region                string                 `json:"region"`
	AlsoKnownAs           string                 `json:"alsoKnownAs"`
	DeregistrationDate    string                 `json:"deregistrationDate"`
	Uprn                  string                 `json:"uprn"`
	OnspdLatitude         float64                `json:"onspdLatitude"`
	OnspdLongitude        float64                `json:"onspdLongitude"`
	OnspdIcbCode          string                 `json:"onspdIcbCode"`
	OnspdIcbName          string                 `json:"onspdIcbName"`
	MainPhoneNumber       string                 `json:"mainPhoneNumber"`
	InspectionDirectorate string                 `json:"inspectionDirectorate"`
	Constituency          string                 `json:"constituency"`
	LocalAuthority        string                 `json:"localAuthority"`
	LastInspection        Inspection             `json:"lastInspection"`
	LastReport            Report                 `json:"lastReport"`
	Contacts              []Contact              `json:"contacts"`
	Relationships         []ProviderRelationship `json:"relationships"`
	RegulatedActivities   []RegulatedActivity    `json:"regulatedActivities"`
	InspectionCategories  []InspectionCategory   `json:"inspectionCategories"`
	InspectionAreas       []InspectionArea       `json:"inspectionAreas"`
	CurrentRatings        Ratings                `json:"currentRatings"`
	HistoricRatings       []HistoricRating       `json:"historicRatings"`
	Reports               []Report               `json:"reports"`
	UnpublishedReports    []UnpublishedReport    `json:"unpublishedReports"`
}

// LocationItem represents a provider's location.
type LocationItem struct {
	LocationID string `json:"locationId"`
	Name       string `json:"locationName"`
	PostalCode string `json:"postalCode"`
}

// LocationsResponse represents a list of locations.
type LocationsResponse struct {
	FirstPageUri    string         `json:"firstPageUri"`
	LastPageUri     string         `json:"lastPageUri"`
	NextPageUri     string         `json:"nextPageUri"`
	Page            int            `json:"page"`
	PerPage         int            `json:"perPage"`
	PreviousPageUri *string        `json:"previousPageUri"`
	Locations       []LocationItem `json:"locations"`
	Total           int            `json:"total"`
	TotalPages      int            `json:"totalPages"`
	Uri             string         `json:"uri"`
}

// Location represents the detailed information of a provider's location.
type Location struct {
	LocationID                  string                   `json:"locationId"`
	ProviderID                  string                   `json:"providerId"`
	OrganisationType            string                   `json:"organisationType"`
	Type                        string                   `json:"type"`
	Name                        string                   `json:"name"`
	BrandID                     string                   `json:"brandId"`
	BrandName                   string                   `json:"brandName"`
	OnspdCcgCode                string                   `json:"onspdCcgCode"`
	OnspdCcgName                string                   `json:"onspdCcgName"`
	OdsCcgCode                  string                   `json:"odsCcgCode"`
	OdsCcgName                  string                   `json:"odsCcgName"`
	OnspdIcbCode                string                   `json:"onspdIcbCode"`
	OnspdIcbName                string                   `json:"onspdIcbName"`
	OdsCode                     string                   `json:"odsCode"`
	RegistrationStatus          string                   `json:"registrationStatus"`
	RegistrationDate            string                   `json:"registrationDate"`
	DeregistrationDate          string                   `json:"deregistrationDate"`
	Dormancy                    string                   `json:"dormancy"`
	DormancyStartDate           string                   `json:"dormancyStartDate"`
	DormancyEndDate             string                   `json:"dormancyEndDate"`
	AlsoKnownAs                 string                   `json:"alsoKnownAs"`
	OnspdLatitude               float64                  `json:"onspdLatitude"`
	OnspdLongitude              float64                  `json:"onspdLongitude"`
	CareHome                    string                   `json:"careHome"`
	InspectionDirectorate       string                   `json:"inspectionDirectorate"`
	Website                     string                   `json:"website"`
	PostalAddressLine1          string                   `json:"postalAddressLine1"`
	PostalAddressLine2          string                   `json:"postalAddressLine2"`
	PostalAddressTownCity       string                   `json:"postalAddressTownCity"`
	PostalAddressCounty         string                   `json:"postalAddressCounty"`
	Region                      string                   `json:"region"`
	PostalCode                  string                   `json:"postalCode"`
	Uprn                        string                   `json:"uprn"`
	MainPhoneNumber             string                   `json:"mainPhoneNumber"`
	RegisteredManagerAbsentDate string                   `json:"registeredManagerAbsentDate"`
	NumberOfBeds                int                      `json:"numberOfBeds"`
	Constituency                string                   `json:"constituency"`
	LocalAuthority              string                   `json:"localAuthority"`
	LastInspection              LastInspection           `json:"lastInspection"`
	LastReport                  LastReport               `json:"lastReport"`
	Relationships               []LocationRelationship   `json:"relationships"`
	LocationTypes               []LocationType           `json:"locationTypes"`
	RegulatedActivities         []RegulatedActivity      `json:"regulatedActivities"`
	GacServiceTypes             []GacServiceType         `json:"gacServiceTypes"`
	Specialisms                 []Specialism             `json:"specialisms"`
	InspectionCategories        []InspectionCategory     `json:"inspectionCategories"`
	InspectionAreas             []InspectionArea         `json:"inspectionAreas"`
	CurrentRatings              Ratings                  `json:"currentRatings"`
	HistoricRatings             []HistoricRating         `json:"historicRatings"`
	Reports                     []Report                 `json:"reports"`
	UnpublishedReports          []UnpublishedReport      `json:"unpublishedReports"`
	ProviderInspectionAreas     []ProviderInspectionArea `json:"providerInspectionAreas"`
}

// Address struct
type PostalAddress struct {
	Line1      string `json:"postalAddressLine1"`
	Line2      string `json:"postalAddressLine2"`
	Town       string `json:"postalAddressTownCity"`
	County     string `json:"postalAddressCounty"`
	PostalCode string `json:"postalCode"`
}

// Inspection struct
type Inspection struct {
	Date string `json:"date"`
}

// InspectionLocation struct
type InspectionLocation struct {
	LocationID string `json:"locationId"`
}

// UnpublishedReport struct
type UnpublishedReport struct {
	FirstVisitDate string `json:"firstVisitDate"`
}

// LastInspection represents the last inspection date.
type LastInspection struct {
	Date string `json:"date"`
}

// LastReport represents the last report publication date.
type LastReport struct {
	PublicationDate string `json:"publicationDate"`
}

// Relationship represents related locations.
type LocationRelationship struct {
	RelatedLocationID   string `json:"relatedLocationId"`
	RelatedLocationName string `json:"relatedLocationName"`
	Type                string `json:"type"`
	Reason              string `json:"reason"`
}

// Relationship struct
type ProviderRelationship struct {
	RelatedProviderID   string `json:"relatedProviderId"`
	RelatedProviderName string `json:"relatedProviderName"`
	Type                string `json:"type"`
	Reason              string `json:"reason"`
}

// LocationType represents types of locations.
type LocationType struct {
	Type string `json:"type"`
}

// Report struct
type Report struct {
	LinkID              string               `json:"linkId"`
	ReportDate          string               `json:"reportDate"`
	FirstVisitDate      string               `json:"firstVisitDate,omitempty"`
	ReportURI           string               `json:"reportUri"`
	ReportType          string               `json:"reportType"`
	InspectionLocations []InspectionLocation `json:"inspectionLocations,omitempty"`
	RelatedDocuments    []RelatedDocument    `json:"relatedDocuments,omitempty"`
}

// RegulatedActivity struct
type RegulatedActivity struct {
	Name                string    `json:"name"`
	Code                string    `json:"code"`
	NominatedIndividual Contact   `json:"nominatedIndividual"`
	Contacts            []Contact `json:"contacts"`
}

// InspectionArea struct
type InspectionArea struct {
	InspectionAreaID   string   `json:"inspectionAreaId"`
	InspectionAreaName string   `json:"inspectionAreaName"`
	InspectionAreaType string   `json:"inspectionAreaType,omitempty"`
	Status             string   `json:"status"`
	SupersededBy       []string `json:"supersededBy,omitempty"`
	EndDate            string   `json:"endDate,omitempty"`
}

// UseOfResources represents resource usage and combined quality rating.
type UseOfResources struct {
	OrganisationID         string `json:"organisationId"`
	Summary                string `json:"summary"`
	UseOfResourcesRating   string `json:"useOfResourcesRating"`
	CombinedQualitySummary string `json:"combinedQualitySummary"`
	CombinedQualityRating  string `json:"combinedQualityRating"`
	ReportDate             string `json:"reportDate"`
	ReportLinkID           string `json:"reportLinkId"`
}

// Contact represents contacts related to regulated activities.
type Contact struct {
	PersonTitle      string   `json:"personTitle"`
	PersonGivenName  string   `json:"personGivenName"`
	PersonFamilyName string   `json:"personFamilyName"`
	PersonRoles      []string `json:"personRoles"`
}

// GacServiceType represents general acute care services.
type GacServiceType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Specialism represents areas of specialisation.
type Specialism struct {
	Name string `json:"name"`
}

// InspectionCategory represents inspection categories.
type InspectionCategory struct {
	Primary string `json:"primary"`
	Code    string `json:"code"`
	Name    string `json:"name"`
}

// Ratings represents the overall and service ratings.
type Ratings struct {
	Overall        Rating          `json:"overall"`
	ServiceRatings []ServiceRating `json:"serviceRatings"`
}

// Rating represents an overall rating.
type Rating struct {
	OrganisationID     string              `json:"organisationId"`
	Rating             string              `json:"rating"`
	ReportDate         string              `json:"reportDate"`
	ReportLinkID       string              `json:"reportLinkId"`
	UseOfResources     UseOfResources      `json:"useOfResources"`
	KeyQuestionRatings []KeyQuestionRating `json:"keyQuestionRatings"`
}

// KeyQuestionRating represents a rating for a specific area.
type KeyQuestionRating struct {
	Name         string `json:"name"`
	Rating       string `json:"rating"`
	ReportDate   string `json:"reportDate,omitempty"`
	ReportLinkID string `json:"reportLinkId,omitempty"`
}

// ServiceRating represents ratings for specific services.
type ServiceRating struct {
	Name               string              `json:"name"`
	Rating             string              `json:"rating"`
	ReportDate         string              `json:"reportDate"`
	OrganisationID     string              `json:"organisationId,omitempty"`
	ReportLinkID       string              `json:"reportLinkId"`
	KeyQuestionRatings []KeyQuestionRating `json:"keyQuestionRatings,omitempty"`
}

// HistoricRating represents past ratings.
type HistoricRating struct {
	ReportDate     string          `json:"reportDate"`
	ReportLinkID   string          `json:"reportLinkId"`
	OrganisationID string          `json:"organisationId"`
	Overall        *Rating         `json:"overall,omitempty"`
	ServiceRatings []ServiceRating `json:"serviceRatings,omitempty"`
}

// RelatedDocument represents documents related to reports.
type RelatedDocument struct {
	DocumentURI  string `json:"documentUri"`
	DocumentType string `json:"documentType"`
}

// ProviderInspectionArea represents provider-specific inspection areas.
type ProviderInspectionArea struct {
	InspectionAreaID string     `json:"inspectionAreaId"`
	Reports          []ReportID `json:"reports"`
}

// ReportID represents a report reference in provider inspection areas.
type ReportID struct {
	InspectionID string `json:"inspectionId"`
	ReportLinkID string `json:"reportLinkId"`
	ProviderID   string `json:"providerId"`
	LocationID   string `json:"locationId"`
}
