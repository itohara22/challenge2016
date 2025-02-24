package main

import (
	"strings"
)

type Distributor struct {
	Name           string
	IncludeRegions map[string]bool
	ExcludeRegions map[string]bool
	Parent         *Distributor
}

func (d *Distributor) Has_Permission(place_code string) bool {
	// exact code in exlude map
	_, exclude_exist := d.ExcludeRegions[place_code]
	if exclude_exist {
		return false
	}

	codes := strings.Split(place_code, "-")

	switch len(codes) {
	case 1:
		_, include_exist := d.IncludeRegions[codes[0]]
		// _, e_exist := d.ExcludeRegions[codes[0]] // this case redundant when providing only country code
		return include_exist

	case 2:
		province_code := codes[0] + "-" + codes[1]
		county_code := codes[1]
		_, inlude_exist := d.IncludeRegions[province_code]
		if inlude_exist {
			return true
		}
		_, include_country_exist := d.IncludeRegions[county_code]
		_, exclude_province_exist := d.ExcludeRegions[province_code]

		if include_country_exist {
			return !exclude_province_exist
		}
		return false

	case 3:
		_, include_city_exist := d.IncludeRegions[place_code]
		if include_city_exist {
			return true
		}

		province_code := codes[1] + "-" + codes[2]
		_, include_province_exist := d.IncludeRegions[province_code]
		_, exclude_city_exist := d.ExcludeRegions[place_code]
		if include_province_exist {
			return !exclude_city_exist
		}

		country_code := codes[2]
		_, include_country_exist := d.IncludeRegions[country_code]
		_, exclude_province_exist := d.ExcludeRegions[province_code]
		if include_country_exist {
			return !exclude_province_exist
		}

	default:
		return false
	}
	return false
}
