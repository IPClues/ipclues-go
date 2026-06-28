package ipclues

// IPVersion is the IP protocol version of the queried address.
type IPVersion uint8

const (
	// IPv4 indicates a 32-bit IPv4 address.
	IPv4 IPVersion = 4
	// IPv6 indicates a 128-bit IPv6 address.
	IPv6 IPVersion = 6
)

// AddressType classifies an IP address by its IANA-designated scope or purpose.
type AddressType string

const (
	// InvalidAddrType indicates the address could not be classified.
	InvalidAddrType AddressType = "invalid"
	// LoopbackAddrType covers 127.0.0.0/8 and ::1.
	LoopbackAddrType AddressType = "loopback"
	// LinkLocalAddrType covers 169.254.0.0/16 and fe80::/10.
	LinkLocalAddrType AddressType = "link_local"
	// MulticastAddrType covers 224.0.0.0/4 and ff00::/8.
	MulticastAddrType AddressType = "multicast"
	// PrivateAddrType covers RFC 1918 ranges and IPv6 ULA (fc00::/7).
	PrivateAddrType AddressType = "private"
	// UnspecifiedAddrType covers IANA special-purpose blocks not matched
	// by the more specific types above, including the unspecified address
	// (0.0.0.0 / ::).
	UnspecifiedAddrType AddressType = "unspecified"
	// ReservedAddrType indicates the address is classified as reserved, might be public but not allocated/assigned
	ReservedAddrType AddressType = "reserved"
	// PublicAddrType indicates a globally routable unicast address.
	PublicAddrType AddressType = "public"
)

// LookupResult is the top-level response from a single IP lookup.
// Fields are populated based on the caller's plan tier; pointer fields
// are nil when not included in the response.
type LookupResult struct {
	// IP is the queried address string, normalised to its canonical form.
	IP string `json:"ip"`
	// Version is the IP protocol version: 4 or 6.
	Version IPVersion `json:"version"`
	// AddressType classifies the address by IANA scope.
	AddressType AddressType `json:"address_type"`
	// Continent is populated at GeoI3-Country tier and above.
	Continent *Continent `json:"continent,omitempty"`
	// Country is populated at GeoI3-Country tier and above.
	Country *Country `json:"country,omitempty"`
	// Currency is the primary currency of the country the IP is located in.
	// Populated at GeoI3-Country tier and above.
	Currency *Currency `json:"currency,omitempty"`
}

// Country contains geographic and demographic attributes of the country
// associated with an IP address.
type Country struct {
	// Code is the ISO 3166-1 alpha-2 country code (e.g. "NG").
	Code string `json:"code,omitempty"`
	// GeoNameID is the GeoNames database identifier for this country.
	GeoNameID uint32 `json:"geoname_id,omitempty"`
	// Name is the English country name (e.g. "Nigeria").
	Name string `json:"name,omitempty"`
	// Neighbors lists ISO 3166-1 alpha-2 codes of bordering countries.
	Neighbors []string `json:"neighbors,omitempty"`
	// CallingCode is the ITU country calling code (e.g. "234" for Nigeria).
	CallingCode string `json:"calling_code,omitempty"`
	// Population is the approximate national population.
	Population uint32 `json:"population,omitempty"`
	// Languages lists BCP 47 / ISO 639-1 language codes spoken in the country.
	Languages []string `json:"languages,omitempty"`
	// TLD is the country-code top-level domain (e.g. ".ng").
	TLD string `json:"tld,omitempty"`
}

// Continent contains the continent associated with an IP address.
type Continent struct {
	// Code is the two-letter continent code (AF, AN, AS, EU, NA, OC, SA).
	Code string `json:"code,omitempty"`
	// Name is the English continent name (e.g. "Africa").
	Name string `json:"name,omitempty"`
}

// Currency contains the primary currency of the country associated with an IP address.
type Currency struct {
	// Code is the ISO 4217 currency code (e.g. "NGN").
	Code string `json:"code,omitempty"`
	// Name is the English currency name (e.g. "Nigerian Naira").
	Name string `json:"name,omitempty"`
}
