package locations

type Country struct {
	Id             string         "jason: id"
	Name           string         "jason: name"
	TimeZone       string         "jason: time_zone"
	GeoInformation GeoInformation "jason: geo_information"
	States         []State        "jason: states"
}

type GeoInformation struct {
	Location GeoLocation "jason: location"
}

type GeoLocation struct {
	Latitude  float64 "jason: latitude"
	Longitude float64 "jason: logitude"
}

type State struct {
	Id   string "jason: id"
	Name string "jason: name"
}
