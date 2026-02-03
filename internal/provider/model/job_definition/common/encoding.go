package common

// normalizeEncoding converts API encoding format to Terraform format.
// API returns: utf_8, utf_16LE, utf_32BE, utf_32LE.
// Terraform uses: UTF-8, UTF-16LE, UTF-32BE, UTF-32LE.
func NormalizeEncoding(apiEncoding string) string {
	switch apiEncoding {
	case "utf_8":
		return "UTF-8"
	case "utf_16LE":
		return "UTF-16LE"
	case "utf_32BE":
		return "UTF-32BE"
	case "utf_32LE":
		return "UTF-32LE"
	default:
		return apiEncoding
	}
}

// denormalizeEncoding converts Terraform encoding format to API format.
// Terraform uses: UTF-8, UTF-16LE, UTF-32BE, UTF-32LE.
// API expects: utf_8, utf_16LE, utf_32BE, utf_32LE.
func DenormalizeEncoding(terraformEncoding string) string {
	switch terraformEncoding {
	case "UTF-8":
		return "utf_8"
	case "UTF-16LE":
		return "utf_16LE"
	case "UTF-32BE":
		return "utf_32BE"
	case "UTF-32LE":
		return "utf_32LE"
	default:
		return terraformEncoding
	}
}
