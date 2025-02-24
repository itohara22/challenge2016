package main

import (
	"fmt"
	"strings"
)

type (
	City struct {
		CityName string
	}
	Province struct {
		Cities       map[string]City
		ProvinceName string
	}
	Country struct {
		Provinces  map[string]Province
		CountyName string
	}
)

type Places struct {
	Countries map[string]Country
}

func LoadPlace(data [][]string) *Places {
	p := &Places{Countries: make(map[string]Country)}

	for _, v := range data {
		county_name, province_name, city_name := v[5], v[4], v[3]
		county_code, province_code, city_code := v[2], v[1], v[0]

		_, county_exist := p.Countries[county_code]
		if !county_exist {
			p.Countries[county_code] = Country{
				// CountryCode: county_code,
				CountyName: county_name,
				Provinces:  make(map[string]Province),
			}
		}
		country := p.Countries[county_code]

		_, province_exist := country.Provinces[province_code]
		if !province_exist {
			country.Provinces[province_code] = Province{
				// ProvinceCode: province_code,
				ProvinceName: province_name,
				Cities:       make(map[string]City),
			}
		}
		province := country.Provinces[province_code]

		province.Cities[city_code] = City{
			CityName: city_name,
			// CityCode: city_code,
		}
	}

	return p
}

func (p *Places) Place_Exist(place_code string) bool {
	codes := strings.Split(place_code, "-")

	switch len(codes) {
	case 1:
		_, e := p.Countries[codes[0]]
		return e

	case 2:
		c, e := p.Countries[codes[1]]
		if !e {
			return false
		}
		_, e = c.Provinces[codes[0]]
		return e

	case 3:
		c, e := p.Countries[codes[2]]
		if !e {
			return false
		}
		p, e := c.Provinces[codes[1]]
		if !e {
			return false
		}
		_, e = p.Cities[codes[0]]
		return e

	default:
		return false
	}
}

func (p *Places) Get_Name_From_Codes(code string) {
	codes := strings.Split(code, "-")

	switch len(codes) {
	case 1:
		c := p.Countries[codes[0]]
		fmt.Println(c.CountyName)

	case 2:
		c := p.Countries[codes[1]]
		p := c.Provinces[codes[0]]

		fmt.Println(p.ProvinceName, ",", c.CountyName)

	case 3:
		c := p.Countries[codes[2]]
		p := c.Provinces[codes[1]]
		ci := p.Cities[codes[0]]
		fmt.Println(ci.CityName, ",", p.ProvinceName, ",", c.CountyName)

	default:
		fmt.Println("Invalid code")
	}
}

func (p *Places) Print() {
	for _, country := range p.Countries {
		fmt.Println("county:", country.CountyName)
		for _, province := range country.Provinces {
			fmt.Println("  province:", province.ProvinceName)
			for _, city := range province.Cities {
				fmt.Println("    city:", city.CityName)
			}
		}
	}
}
