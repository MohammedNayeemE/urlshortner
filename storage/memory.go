package storage

var Map = make(map[string](string));

func Save(code string , url string) {	
	Map[code] = url;
}

func GetUrl(code string) string {
	return Map[code];
}
