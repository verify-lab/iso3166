package iso3166

import "strings"

const (
	alpha2CodeLength = 2
	alpha3CodeLength = 3
)

type country struct {
	num    int
	name   string
	alpha2 string
	alpha3 string
}

var countries = []country{
	{4, "Afghanistan", "af", "afg"},
	{8, "Albania", "al", "alb"},
	{12, "Algeria", "dz", "dza"},
	{20, "Andorra", "ad", "and"},
	{24, "Angola", "ao", "ago"},
	{28, "Antigua and Barbuda", "ag", "atg"},
	{32, "Argentina", "ar", "arg"},
	{51, "Armenia", "am", "arm"},
	{36, "Australia", "au", "aus"},
	{40, "Austria", "at", "aut"},
	{31, "Azerbaijan", "az", "aze"},
	{44, "Bahamas", "bs", "bhs"},
	{48, "Bahrain", "bh", "bhr"},
	{50, "Bangladesh", "bd", "bgd"},
	{52, "Barbados", "bb", "brb"},
	{112, "Belarus", "by", "blr"},
	{56, "Belgium", "be", "bel"},
	{84, "Belize", "bz", "blz"},
	{204, "Benin", "bj", "ben"},
	{64, "Bhutan", "bt", "btn"},
	{68, "Bolivia (Plurinational State of)", "bo", "bol"},
	{70, "Bosnia and Herzegovina", "ba", "bih"},
	{72, "Botswana", "bw", "bwa"},
	{76, "Brazil", "br", "bra"},
	{96, "Brunei Darussalam", "bn", "brn"},
	{100, "Bulgaria", "bg", "bgr"},
	{854, "Burkina Faso", "bf", "bfa"},
	{108, "Burundi", "bi", "bdi"},
	{132, "Cabo Verde", "cv", "cpv"},
	{116, "Cambodia", "kh", "khm"},
	{120, "Cameroon", "cm", "cmr"},
	{124, "Canada", "ca", "can"},
	{140, "Central African Republic", "cf", "caf"},
	{148, "Chad", "td", "tcd"},
	{152, "Chile", "cl", "chl"},
	{156, "China", "cn", "chn"},
	{170, "Colombia", "co", "col"},
	{174, "Comoros", "km", "com"},
	{178, "Congo", "cg", "cog"},
	{180, "Congo, Democratic Republic of the", "cd", "cod"},
	{188, "Costa Rica", "cr", "cri"},
	{384, "Côte d'Ivoire", "ci", "civ"},
	{191, "Croatia", "hr", "hrv"},
	{192, "Cuba", "cu", "cub"},
	{196, "Cyprus", "cy", "cyp"},
	{203, "Czechia", "cz", "cze"},
	{208, "Denmark", "dk", "dnk"},
	{262, "Djibouti", "dj", "dji"},
	{212, "Dominica", "dm", "dma"},
	{214, "Dominican Republic", "do", "dom"},
	{218, "Ecuador", "ec", "ecu"},
	{818, "Egypt", "eg", "egy"},
	{222, "El Salvador", "sv", "slv"},
	{226, "Equatorial Guinea", "gq", "gnq"},
	{232, "Eritrea", "er", "eri"},
	{233, "Estonia", "ee", "est"},
	{748, "Eswatini", "sz", "swz"},
	{231, "Ethiopia", "et", "eth"},
	{242, "Fiji", "fj", "fji"},
	{246, "Finland", "fi", "fin"},
	{250, "France", "fr", "fra"},
	{266, "Gabon", "ga", "gab"},
	{270, "Gambia", "gm", "gmb"},
	{268, "Georgia", "ge", "geo"},
	{276, "Germany", "de", "deu"},
	{288, "Ghana", "gh", "gha"},
	{300, "Greece", "gr", "grc"},
	{308, "Grenada", "gd", "grd"},
	{320, "Guatemala", "gt", "gtm"},
	{324, "Guinea", "gn", "gin"},
	{624, "Guinea-Bissau", "gw", "gnb"},
	{328, "Guyana", "gy", "guy"},
	{332, "Haiti", "ht", "hti"},
	{340, "Honduras", "hn", "hnd"},
	{348, "Hungary", "hu", "hun"},
	{352, "Iceland", "is", "isl"},
	{356, "India", "in", "ind"},
	{360, "Indonesia", "id", "idn"},
	{364, "Iran (Islamic Republic of)", "ir", "irn"},
	{368, "Iraq", "iq", "irq"},
	{372, "Ireland", "ie", "irl"},
	{376, "Israel", "il", "isr"},
	{380, "Italy", "it", "ita"},
	{388, "Jamaica", "jm", "jam"},
	{392, "Japan", "jp", "jpn"},
	{400, "Jordan", "jo", "jor"},
	{398, "Kazakhstan", "kz", "kaz"},
	{404, "Kenya", "ke", "ken"},
	{296, "Kiribati", "ki", "kir"},
	{408, "Korea (Democratic People's Republic of)", "kp", "prk"},
	{410, "Korea, Republic of", "kr", "kor"},
	{414, "Kuwait", "kw", "kwt"},
	{417, "Kyrgyzstan", "kg", "kgz"},
	{418, "Lao People's Democratic Republic", "la", "lao"},
	{428, "Latvia", "lv", "lva"},
	{422, "Lebanon", "lb", "lbn"},
	{426, "Lesotho", "ls", "lso"},
	{430, "Liberia", "lr", "lbr"},
	{434, "Libya", "ly", "lby"},
	{438, "Liechtenstein", "li", "lie"},
	{440, "Lithuania", "lt", "ltu"},
	{442, "Luxembourg", "lu", "lux"},
	{450, "Madagascar", "mg", "mdg"},
	{454, "Malawi", "mw", "mwi"},
	{458, "Malaysia", "my", "mys"},
	{462, "Maldives", "mv", "mdv"},
	{466, "Mali", "ml", "mli"},
	{470, "Malta", "mt", "mlt"},
	{584, "Marshall Islands", "mh", "mhl"},
	{478, "Mauritania", "mr", "mrt"},
	{480, "Mauritius", "mu", "mus"},
	{484, "Mexico", "mx", "mex"},
	{583, "Micronesia (Federated States of)", "fm", "fsm"},
	{498, "Moldova, Republic of", "md", "mda"},
	{492, "Monaco", "mc", "mco"},
	{496, "Mongolia", "mn", "mng"},
	{499, "Montenegro", "me", "mne"},
	{504, "Morocco", "ma", "mar"},
	{508, "Mozambique", "mz", "moz"},
	{104, "Myanmar", "mm", "mmr"},
	{516, "Namibia", "na", "nam"},
	{520, "Nauru", "nr", "nru"},
	{524, "Nepal", "np", "npl"},
	{528, "Netherlands", "nl", "nld"},
	{554, "New Zealand", "nz", "nzl"},
	{558, "Nicaragua", "ni", "nic"},
	{562, "Niger", "ne", "ner"},
	{566, "Nigeria", "ng", "nga"},
	{807, "North Macedonia", "mk", "mkd"},
	{578, "Norway", "no", "nor"},
	{512, "Oman", "om", "omn"},
	{586, "Pakistan", "pk", "pak"},
	{585, "Palau", "pw", "plw"},
	{591, "Panama", "pa", "pan"},
	{598, "Papua New Guinea", "pg", "png"},
	{600, "Paraguay", "py", "pry"},
	{604, "Peru", "pe", "per"},
	{608, "Philippines", "ph", "phl"},
	{616, "Poland", "pl", "pol"},
	{620, "Portugal", "pt", "prt"},
	{634, "Qatar", "qa", "qat"},
	{642, "Romania", "ro", "rou"},
	{643, "Russian Federation", "ru", "rus"},
	{646, "Rwanda", "rw", "rwa"},
	{659, "Saint Kitts and Nevis", "kn", "kna"},
	{662, "Saint Lucia", "lc", "lca"},
	{670, "Saint Vincent and the Grenadines", "vc", "vct"},
	{882, "Samoa", "ws", "wsm"},
	{674, "San Marino", "sm", "smr"},
	{678, "Sao Tome and Principe", "st", "stp"},
	{682, "Saudi Arabia", "sa", "sau"},
	{686, "Senegal", "sn", "sen"},
	{688, "Serbia", "rs", "srb"},
	{690, "Seychelles", "sc", "syc"},
	{694, "Sierra Leone", "sl", "sle"},
	{702, "Singapore", "sg", "sgp"},
	{703, "Slovakia", "sk", "svk"},
	{705, "Slovenia", "si", "svn"},
	{90, "Solomon Islands", "sb", "slb"},
	{706, "Somalia", "so", "som"},
	{710, "South Africa", "za", "zaf"},
	{728, "South Sudan", "ss", "ssd"},
	{724, "Spain", "es", "esp"},
	{144, "Sri Lanka", "lk", "lka"},
	{729, "Sudan", "sd", "sdn"},
	{740, "Suriname", "sr", "sur"},
	{752, "Sweden", "se", "swe"},
	{756, "Switzerland", "ch", "che"},
	{760, "Syrian Arab Republic", "sy", "syr"},
	{762, "Tajikistan", "tj", "tjk"},
	{834, "Tanzania, United Republic of", "tz", "tza"},
	{764, "Thailand", "th", "tha"},
	{626, "Timor-Leste", "tl", "tls"},
	{768, "Togo", "tg", "tgo"},
	{776, "Tonga", "to", "ton"},
	{780, "Trinidad and Tobago", "tt", "tto"},
	{788, "Tunisia", "tn", "tun"},
	{792, "Turkey", "tr", "tur"},
	{795, "Turkmenistan", "tm", "tkm"},
	{798, "Tuvalu", "tv", "tuv"},
	{800, "Uganda", "ug", "uga"},
	{804, "Ukraine", "ua", "ukr"},
	{784, "United Arab Emirates", "ae", "are"},
	{826, "United Kingdom of Great Britain and Northern Ireland", "gb", "gbr"},
	{840, "United States of America", "us", "usa"},
	{858, "Uruguay", "uy", "ury"},
	{860, "Uzbekistan", "uz", "uzb"},
	{548, "Vanuatu", "vu", "vut"},
	{862, "Venezuela (Bolivarian Republic of)", "ve", "ven"},
	{704, "Viet Nam", "vn", "vnm"},
	{887, "Yemen", "ye", "yem"},
	{894, "Zambia", "zm", "zmb"},
	{716, "Zimbabwe", "zw", "zwe"},
}

func IsAlpha2CountryCode(code string) bool {
	if len(code) != alpha2CodeLength {
		return false
	}

	code = strings.ToLower(code)

	for _, c := range countries {
		if c.alpha2 == code {
			return true
		}
	}

	return false
}

func IsAlpha3CountryCode(code string) bool {
	if len(code) != alpha3CodeLength {
		return false
	}

	code = strings.ToLower(code)

	for _, c := range countries {
		if c.alpha3 == code {
			return true
		}
	}

	return false
}

func CountryByCode(code string) (num int, name, alpha2, alpha3 string) {
	code = strings.ToLower(code)

	for _, c := range countries {
		if c.alpha3 != code && c.alpha2 != code {
			continue
		}

		num = c.num
		name = c.name
		alpha2 = c.alpha2
		alpha3 = c.alpha3

		break
	}

	return num, name, alpha2, alpha3
}
