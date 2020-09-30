package _6_extended_switch

func getCupSize(milliliters int) string {
	if milliliters <= 300 {
		return "small"
	} else if milliliters <= 500 {
		return "medium"
	} else if milliliters <= 650 {
		return "large"
	} else {
		return "bucket"
	}
}

func getCupSizeImproved(milliliters int) string {
	switch {
	case milliliters <= 300:
		return "small"

	case milliliters <= 500:
		return "medium"

	case milliliters <= 650:
		return "large"

	default:
		return "bucket"
	}
}
