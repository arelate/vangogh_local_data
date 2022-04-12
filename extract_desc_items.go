package vangogh_local_data

import "regexp"

var (
	descItems     = regexp.MustCompile(`https://items.gog.com/([-\pL0-9()!@:%_,'™\ \+\[\].~#?&\/\/=]*)`)
	descGameLinks = regexp.MustCompile(`https://www.gog.com/game/([-\pL0-9()!@:%_,'™\ \+\[\].~#?&\/\/=]*)`)
)

func ExtractDescItems(desc string) []string {
	return descItems.FindAllString(desc, -1)
}

func ExtractGameLinks(desc string) []string {
	return descGameLinks.FindAllString(desc, -1)
}
