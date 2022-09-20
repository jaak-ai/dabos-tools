package transclude

import (
	"encoding/xml"
)

func ByXML(origin interface{}, destine interface{}) {
	bData, err := xml.Marshal(origin)
	if err != nil {
		return
	}

	err = xml.Unmarshal(bData, destine)
	if err != nil {
		return
	}
}

func ListByXML[T any, U any](originList []T, destineList *[]U) {

	for _, origin := range originList {
		newDestine := new(U)

		bOrigin, err := xml.Marshal(origin)
		if err != nil {
			return
		}

		err = xml.Unmarshal(bOrigin, newDestine)
		if err != nil {
			return
		}

		*destineList = append(*destineList, *newDestine)
	}

}
