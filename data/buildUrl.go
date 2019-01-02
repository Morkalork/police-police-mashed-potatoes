package data

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

var CrimeTypes = [...]string{
	"Alkohollagen",
	"Anträffad död",
	"Anträffat gods",
	"Arbetsplatsolycka",
	"Bedrägeri",
	"Bombhot",
	"Brand",
	"Brand automatlarm",
	"Bråk",
	"Detonation",
	"Djur skadat/omhändertaget",
	"Ekobrott",
	"Farligt föremål, misstänkt",
	"Fjällräddning",
	"Fylleri/LOB",
	"Förfalskningsbrott",
	"Försvunnen person",
	"Gränskontroll",
	"HäleriInbrott",
	"Inbrott, försök",
	"Knivlagen",
	"Kontroll person/fordon",
	"Lagen om hundar och katter",
	"Larm Inbrott",
	"Larm Överfall",
	"Miljöbrott",
	"Missbruk av urkund",
	"Misshandel",
	"Misshandel, grov",
	"Mord/dråp",
	"Mord/dråp, försök",
	"Motorfordon, anträffat stulet",
	"Motorfordon, stöld",
	"Narkotikabrott",
	"Naturkatastrof",
	"Ofog barn/ungdom",
	"Ofredande/förargelse",
	"Olaga frihetsberövande",
	"Olaga hot",
	"Olaga intrång/hemfridsbrott",
	"Olovlig körning",
	"Ordningslagen",
	"Polisinsats/kommendering",
	"Rattfylleri",
	"Rån",
	"Rån väpnat",
	"Rån övrigt",
	"Rån, försök",
	"Räddningsinsats",
	"Sammanfattning dag",
	"Sammanfattning dygn",
	"Sammanfattning eftermiddag",
	"Sammanfattning förmiddag",
	"Sammanfattning helg",
	"Sammanfattning kväll",
	"Sammanfattning kväll och natt",
	"Sammanfattning natt",
	"Sammanfattning vecka",
	"Sedlighetsbrott",
	"Sjukdom/olycksfall",
	"Sjölagen",
	"Skadegörelse",
	"Skottlossning",
	"Skottlossning, misstänkt",
	"Spridning smittsamma kemikalier",
	"Stöld",
	"Stöld, försök",
	"Stöld, ringa",
	"Stöld/inbrott",
	"Tillfälligt obemannat",
	"Trafikbrott",
	"Trafikhinder",
	"Trafikkontroll",
	"Trafikolycka",
	"Trafikolycka, personskada",
	"Trafikolycka, singel",
	"Trafikolycka, smitning från",
	"Trafikolycka, vilt",
	"Uppdatering",
	"Utlänningslagen",
	"Vapenlagen",
	"Varningslarm/haveri",
	"Våld/hot mot tjänsteman",
	"Våldtäkt",
	"Våldtäkt, försök",
	"Vållande till kroppsskada",
}

func buildUrl(hours int, location string, crimeType string) (resultUrl string, error error) {
	now := time.Now()
	newDate := now.Add(time.Duration(-hours) * time.Hour)

	uri := fmt.Sprintf(
		"https://polisen.se/api/events?DateTime=%d-%02d-%02d%%20%d",
		newDate.Year(),
		newDate.Month(),
		newDate.Day(),
		newDate.Hour(),
	)

	if location != "" {
		uri += fmt.Sprintf("&locationname=%s", url.QueryEscape(location))
	}

	if crimeType != "" {
		found := false
		for _,crime := range CrimeTypes {
			if crime == crimeType {
				found = true
				break
			}
		}

		if found {
			uri += fmt.Sprintf("&type=%s", url.QueryEscape(crimeType))
		} else {
			error = errors.New(fmt.Sprintf("'%s' is not a valid crime type, remember that they are type sensitive.", crimeType))
		}
	}

	return uri, error
}
