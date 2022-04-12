package vangogh_local_data

import "regexp"

var re = regexp.MustCompile(`https://items.gog.com/([-a-zA-Z0-9()!@:%_,'\+.~#?&\/\/=]*)`)

func ExtractDescItems(desc string) []string {
	return re.FindAllString(desc, -1)
}
