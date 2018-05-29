package convert

func urlLen(row []interface{}) interface{} {
	return len(row[0].(string))
}

func priority(row []interface{}) interface{} {
	return DEFAULT_PRIORITY
}

func loc(row []interface{}) interface{} {
	return row[0].(string)
}

func lastmod(row []interface{}) interface{} {
	return DEFAULT_LASTMOD
}

func changefreq(row []interface{}) interface{} {
	return DEFAULT_CHANGE_FREQ
}
