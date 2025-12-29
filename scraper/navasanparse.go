package scraper

func ParseNavasan(raw map[string]interface{}) map[string]string {
	out := make(map[string]string)
	for k, v := range raw {
		out[k] = fmt.Sprintf("%v", v)
	}
	return out
}
