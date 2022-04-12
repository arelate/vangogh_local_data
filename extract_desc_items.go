package vangogh_local_data

import "regexp"

var re = regexp.MustCompile(`https://items.gog.com/([^"']*?).*?(?="|$)`)

func ExtractDescItems(desc string) []string {
	return re.FindAllString(desc, -1)
}
