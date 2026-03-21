package http

func getPostTitleFontSize(title string) string {
	titleLen := len(title)
	switch true {
	case titleLen <= 10:
		return "5em"
	case titleLen < 20:
		return "4.6em"
	case titleLen < 30:
		return "4.2em"
	case titleLen < 40:
		return "3.6em"
	case titleLen < 50:
		return "3.2em"
	case titleLen < 60:
		return "3em"
	case titleLen < 70:
		return "2.6em"
	default:
		return "2em"
	}
}
