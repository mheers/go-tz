package tz

import (
	_ "time/tzdata"
)

func TranslateMSTimezoneToIANA(tzid string) string {
	// see https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/default-time-zones?view=windows-11
	switch tzid {
	// Afghanistan Standard Time	(UTC+04:30)	Kabul
	case "Afghanistan Standard Time":
		return "Asia/Kabul"
	// Arab Standard Time	(UTC+03:00)	Kuwait, Riyadh
	case "Arab Standard Time":
		return "Asia/Riyadh"
	// Arabian Standard Time	(UTC+04:00)	Abu Dhabi, Muscat
	case "Arabian Standard Time":
		return "Asia/Dubai"
	// Arabic Standard Time	(UTC+03:00)	Baghdad
	case "Arabic Standard Time":
		return "Asia/Baghdad"
	// Argentina Standard Time	(UTC-03:00)	City of Buenos Aires
	case "Argentina Standard Time":
		return "America/Argentina/Buenos_Aires"
	// Atlantic Standard Time	(UTC-04:00)	Atlantic Time (Canada)
	case "Atlantic Standard Time":
		return "America/Halifax"
	// AUS Eastern Standard Time	(UTC+10:00)	Canberra, Melbourne, Sydney
	case "AUS Eastern Standard Time":
		return "Australia/Sydney"
	// Azerbaijan Standard Time	(UTC+04:00)	Baku
	case "Azerbaijan Standard Time":
		return "Asia/Baku"
	// Bangladesh Standard Time	(UTC+06:00)	Dhaka
	case "Bangladesh Standard Time":
		return "Asia/Dhaka"
	// Belarus Standard Time	(UTC+03:00)	Minsk
	case "Belarus Standard Time":
		return "Europe/Minsk"
	// Cape Verde Standard Time	(UTC-01:00)	Cabo Verde Is.
	case "Cape Verde Standard Time":
		return "Atlantic/Cape_Verde"
	// Caucasus Standard Time	(UTC+04:00)	Yerevan
	case "Caucasus Standard Time":
		return "Asia/Yerevan"
	// Central America Standard Time	(UTC-06:00)	Central America
	case "Central America Standard Time":
		return "America/Guatemala"
	// Central Asia Standard Time	(UTC+06:00)	Astana
	case "Central Asia Standard Time":
		return "Asia/Almaty"
	// Central Europe Standard Time	(UTC+01:00)	Belgrade, Bratislava, Budapest, Ljubljana, Prague
	case "Central Europe Standard Time":
		return "Europe/Belgrade"
	// Central European Standard Time	(UTC+01:00)	Sarajevo, Skopje, Warsaw, Zagreb
	case "Central European Standard Time":
		return "Europe/Warsaw"
	// Central Pacific Standard Time	(UTC+11:00)	Solomon Is., New Caledonia
	case "Central Pacific Standard Time":
		return "Pacific/Guadalcanal"
	// Central Standard Time (Mexico)	(UTC-06:00)	Guadalajara, Mexico City, Monterrey
	case "Central Standard Time (Mexico)":
		return "America/Mexico_City"
	// China Standard Time	(UTC+08:00)	Beijing, Chongqing, Hong Kong SAR, Urumqi
	case "China Standard Time":
		return "Asia/Shanghai"
	// E. Africa Standard Time	(UTC+03:00)	Nairobi
	case "E. Africa Standard Time":
		return "Africa/Nairobi"
	// E. Europe Standard Time	(UTC+02:00)	E. Europe
	case "E. Europe Standard Time":
		return "Europe/Bucharest"
	// E. South America Standard Time	(UTC-03:00)	Brasilia
	case "E. South America Standard Time":
		return "America/Sao_Paulo"
	// Eastern Standard Time	(UTC-05:00)	Eastern Time (US & Canada)
	case "Eastern Standard Time":
		return "America/New_York"
	// Egypt Standard Time	(UTC+02:00)	Cairo
	case "Egypt Standard Time":
		return "Africa/Cairo"
	// Fiji Standard Time	(UTC+12:00)	Fiji
	case "Fiji Standard Time":
		return "Pacific/Fiji"
	// FLE Standard Time	(UTC+02:00)	Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius
	case "FLE Standard Time":
		return "Europe/Helsinki"
	// Georgian Standard Time	(UTC+04:00)	Tbilisi
	case "Georgian Standard Time":
		return "Asia/Tbilisi"
	// GMT Standard Time	(UTC)	Dublin, Edinburgh, Lisbon, London
	case "GMT Standard Time":
		return "Europe/London"
	// Greenland Standard Time	(UTC-03:00)	Greenland
	case "Greenland Standard Time":
		return "America/Godthab"
	// Greenwich Standard Time	(UTC)	Monrovia, Reykjavik
	case "Greenwich Standard Time":
		return "Atlantic/Reykjavik"
	// GTB Standard Time	(UTC+02:00)	Athens, Bucharest
	case "GTB Standard Time":
		return "Europe/Athens"
	// Hawaiian Standard Time	(UTC-10:00)	Hawaii
	case "Hawaiian Standard Time":
		return "Pacific/Honolulu"
	// India Standard Time	(UTC+05:30)	Chennai, Kolkata, Mumbai, New Delhi
	case "India Standard Time":
		return "Asia/Kolkata"
	// Iran Standard Time	(UTC+03:30)	Tehran
	case "Iran Standard Time":
		return "Asia/Tehran"
	// Israel Standard Time	(UTC+02:00)	Middle East
	case "Israel Standard Time":
		return "Asia/Jerusalem"
	// Jordan Standard Time	(UTC+02:00)	Amman
	case "Jordan Standard Time":
		return "Asia/Amman"
	// Korea Standard Time	(UTC+09:00)	Seoul
	case "Korea Standard Time":
		return "Asia/Seoul"
	// Mauritius Standard Time	(UTC+04:00)	Port Louis
	case "Mauritius Standard Time":
		return "Indian/Mauritius"
	// Middle East Standard Time	(UTC+02:00)	Beirut
	case "Middle East Standard Time":
		return "Asia/Beirut"
	// Montevideo Standard Time	(UTC-03:00)	Montevideo
	case "Montevideo Standard Time":
		return "America/Montevideo"
	// Morocco Standard Time	(UTC)	Casablanca
	case "Morocco Standard Time":
		return "Africa/Casablanca"
	// Mountain Standard Time	(UTC-07:00)	Mountain Time (US & Canada)
	case "Mountain Standard Time":
		return "America/Denver"
	// Myanmar Standard Time	(UTC+06:30)	Yangon (Rangoon)
	case "Myanmar Standard Time":
		return "Asia/Yangon"
	// Namibia Standard Time	(UTC+01:00)	Windhoek
	case "Namibia Standard Time":
		return "Africa/Windhoek"
	// Nepal Standard Time	(UTC+05:45)	Kathmandu
	case "Nepal Standard Time":
		return "Asia/Kathmandu"
	// New Zealand Standard Time	(UTC+12:00)	Auckland, Wellington
	case "New Zealand Standard Time":
		return "Pacific/Auckland"
	// Pacific SA Standard Time	(UTC-03:00)	Santiago
	case "Pacific SA Standard Time":
		return "America/Santiago"
	// Pacific Standard Time	(UTC-08:00)	Pacific Time (US & Canada)
	case "Pacific Standard Time":
		return "America/Los_Angeles"
	// Pakistan Standard Time	(UTC+05:00)	Islamabad, Karachi
	case "Pakistan Standard Time":
		return "Asia/Karachi"
	// Paraguay Standard Time	(UTC-04:00)	Asuncion
	case "Paraguay Standard Time":
		return "America/Asuncion"
	// Romance Standard Time	(UTC+01:00)	Brussels, Copenhagen, Madrid, Paris
	case "Romance Standard Time":
		return "Europe/Paris"
	// Russian Standard Time	(UTC+03:00)	Moscow, St. Petersburg, Volgograd (RTZ 2)
	case "Russian Standard Time":
		return "Europe/Moscow"
	// SA Eastern Standard Time	(UTC-03:00)	Cayenne, Fortaleza
	case "SA Eastern Standard Time":
	// SA Pacific Standard Time	(UTC-05:00)	Bogota, Lima, Quito, Rio Branco
	case "SA Pacific Standard Time":
		return "America/Bogota"
	// SA Western Standard Time	(UTC-04:00)	Georgetown, La Paz, Manaus, San Juan
	case "SA Western Standard Time":
		return "America/Manaus"
	// Samoa Standard Time	(UTC+13:00)	Samoa
	case "Samoa Standard Time":
		return "Pacific/Apia"
	// SE Asia Standard Time	(UTC+07:00)	Bangkok, Hanoi, Jakarta
	case "SE Asia Standard Time":
		return "Asia/Bangkok"
	// Singapore Standard Time	(UTC+08:00)	Kuala Lumpur, Singapore
	case "Singapore Standard Time":
		return "Asia/Singapore"
	// South Africa Standard Time	(UTC+02:00)	Harare, Pretoria
	case "South Africa Standard Time":
		return "Africa/Johannesburg"
	// Sri Lanka Standard Time	(UTC+05:30)	Sri Jayawardenepura
	case "Sri Lanka Standard Time":
		return "Asia/Colombo"
	// Syria Standard Time	(UTC+02:00)	Damascus
	case "Syria Standard Time":
		return "Asia/Damascus"
	// Taipei Standard Time	(UTC+08:00)	Taipei
	case "Taipei Standard Time":
		return "Asia/Taipei"
	// Tokyo Standard Time	(UTC+09:00)	Osaka, Sapporo, Tokyo
	case "Tokyo Standard Time":
		return "Asia/Tokyo"
	// Tonga Standard Time	(UTC+13:00)	Nuku'alofa
	case "Tonga Standard Time":
		return "Pacific/Tongatapu"
	// Türkiye Standard Time	(UTC+02:00)	Istanbul
	case "Türkiye Standard Time":
		return "Europe/Istanbul"
	// Ulaanbaatar Standard Time	(UTC+08:00)	Ulaanbaatar
	case "Ulaanbaatar Standard Time":
		return "Asia/Ulaanbaatar"
	// UTC	(UTC)	Coordinated Universal Time
	case "UTC":
		return "UTC"
	// UTC-02	(UTC-02:00)	Coordinated Universal Time-02
	case "UTC-02":
		return "Etc/GMT+2"
	// UTC-11	(UTC-11:00)	Coordinated Universal Time-11
	case "UTC-11":
		return "Etc/GMT+11"
	// UTC+12	(UTC+12:00)	Coordinated Universal Time+12
	case "UTC+12":
		return "Etc/GMT-12"
	// Venezuela Standard Time	(UTC-04:30)	Caracas
	case "Venezuela Standard Time":
		return "America/Caracas"
	// W. Central Africa Standard Time	(UTC+01:00)	West Central Africa
	case "W. Central Africa Standard Time":
		return "Africa/Lagos"
	// W. Europe Standard Time	(UTC+01:00)	Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna
	case "W. Europe Standard Time":
		return "Europe/Berlin"
	// West Asia Standard Time	(UTC+05:00)	Ashgabat, Tashkent
	case "West Asia Standard Time":
		return "Asia/Tashkent"
	// West Pacific Standard Time	(UTC+10:00)	Guam, Port Moresby
	case "West Pacific Standard Time":
		return "Pacific/Port_Moresby"
	}
	return tzid
}
