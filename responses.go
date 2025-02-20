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
	AlsoKnownAs           string                 `json:"alsoKnownAs"`
	BrandID               string                 `json:"brandId"`
	BrandName             string                 `json:"brandName"`
	CharityNumber         string                 `json:"charityNumber"`
	CompaniesHouseNumber  string                 `json:"companiesHouseNumber"`
	Constituency          string                 `json:"constituency"`
	Contacts              []Contact              `json:"contacts"`
	CurrentRatings        Ratings                `json:"currentRatings"`
	DeregistrationDate    string                 `json:"deregistrationDate"`
	HistoricRatings       []HistoricRating       `json:"historicRatings"`
	InspectionAreas       []InspectionArea       `json:"inspectionAreas"`
	InspectionCategories  []InspectionCategory   `json:"inspectionCategories"`
	InspectionDirectorate string                 `json:"inspectionDirectorate"`
	LastInspection        Inspection             `json:"lastInspection"`
	LastReport            Report                 `json:"lastReport"`
	LocalAuthority        string                 `json:"localAuthority"`
	LocationIDs           []string               `json:"locationIds"`
	MainPhoneNumber       string                 `json:"mainPhoneNumber"`
	Name                  string                 `json:"name"`
	OdsCode               string                 `json:"odsCode"`
	OnspdIcbCode          string                 `json:"onspdIcbCode"`
	OnspdIcbName          string                 `json:"onspdIcbName"`
	OnspdLatitude         float64                `json:"onspdLatitude"`
	OnspdLongitude        float64                `json:"onspdLongitude"`
	OrganisationType      string                 `json:"organisationType"`
	OwnershipType         string                 `json:"ownershipType"`
	PostalAddress         PostalAddress          `json:"-"`
	Region                string                 `json:"region"`
	RegistrationDate      string                 `json:"registrationDate"`
	RegistrationStatus    string                 `json:"registrationStatus"`
	RegulatedActivities   []RegulatedActivity    `json:"regulatedActivities"`
	Relationships         []ProviderRelationship `json:"relationships"`
	Reports               []Report               `json:"reports"`
	Type                  string                 `json:"type"`
	UnpublishedReports    []UnpublishedReport    `json:"unpublishedReports"`
	Uprn                  string                 `json:"uprn"`
	Website               string                 `json:"website"`
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
	AlsoKnownAs                 string                   `json:"alsoKnownAs"`
	BrandID                     string                   `json:"brandId"`
	BrandName                   string                   `json:"brandName"`
	CareHome                    string                   `json:"careHome"`
	Constituency                string                   `json:"constituency"`
	CurrentRatings              *Ratings                 `json:"currentRatings"`
	DeregistrationDate          string                   `json:"deregistrationDate"`
	Dormancy                    string                   `json:"dormancy"`
	DormancyEndDate             string                   `json:"dormancyEndDate"`
	DormancyStartDate           string                   `json:"dormancyStartDate"`
	GacServiceTypes             []GacServiceType         `json:"gacServiceTypes"`
	HistoricRatings             []HistoricRating         `json:"historicRatings"`
	InspectionAreas             []InspectionArea         `json:"inspectionAreas"`
	InspectionCategories        []InspectionCategory     `json:"inspectionCategories"`
	InspectionDirectorate       string                   `json:"inspectionDirectorate"`
	LastInspection              LastInspection           `json:"lastInspection"`
	LastReport                  LastReport               `json:"lastReport"`
	LocalAuthority              string                   `json:"localAuthority"`
	LocationTypes               []LocationType           `json:"locationTypes"`
	MainPhoneNumber             string                   `json:"mainPhoneNumber"`
	Name                        string                   `json:"name"`
	NumberOfBeds                int                      `json:"numberOfBeds"`
	OdsCcgCode                  string                   `json:"odsCcgCode"`
	OdsCcgName                  string                   `json:"odsCcgName"`
	OdsCode                     string                   `json:"odsCode"`
	OnspdCcgCode                string                   `json:"onspdCcgCode"`
	OnspdCcgName                string                   `json:"onspdCcgName"`
	OnspdIcbCode                string                   `json:"onspdIcbCode"`
	OnspdIcbName                string                   `json:"onspdIcbName"`
	OnspdLatitude               *float64                 `json:"onspdLatitude"`
	OnspdLongitude              *float64                 `json:"onspdLongitude"`
	OrganisationType            string                   `json:"organisationType"`
	PostalAddressCounty         string                   `json:"postalAddressCounty"`
	PostalAddressLine1          string                   `json:"postalAddressLine1"`
	PostalAddressLine2          string                   `json:"postalAddressLine2"`
	PostalAddressTownCity       string                   `json:"postalAddressTownCity"`
	PostalCode                  string                   `json:"postalCode"`
	ProviderInspectionAreas     []ProviderInspectionArea `json:"providerInspectionAreas"`
	Region                      string                   `json:"region"`
	RegisteredManagerAbsentDate string                   `json:"registeredManagerAbsentDate"`
	RegistrationDate            string                   `json:"registrationDate"`
	RegistrationStatus          string                   `json:"registrationStatus"`
	RegulatedActivities         []RegulatedActivity      `json:"regulatedActivities"`
	Relationships               []LocationRelationship   `json:"relationships"`
	Reports                     []Report                 `json:"reports"`
	ServicesProviders           *string                  `json:"servicesProviders"`
	Specialisms                 []Specialism             `json:"specialisms"`
	Type                        string                   `json:"type"`
	UnpublishedReports          []UnpublishedReport      `json:"unpublishedReports"`
	Uprn                        string                   `json:"uprn"`
	Website                     string                   `json:"website"`
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
