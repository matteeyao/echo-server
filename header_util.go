package main

// TimeFormat is the time format to use when generating times in HTTP
// headers. It is like time.RFC1123 but hard-codes GMT as the time
// zone. The time being formatted must be in UTC for Format to
// generate the correct format.
//
// For parsing this time format, see ParseTime.
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

// ValidHeaderFieldName reports whether v is a valid HTTP/1.x header name.
// HTTP/2 imposes the additional restriction that uppercase ASCII
// letters are not allowed.
//
//  RFC 7230 says:
//   header-field   = field-name ":" OWS field-value OWS
//   field-name     = token
//   token          = 1*tchar
//   tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//           "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
func ValidHeaderFieldName(v string) bool {
	if len(v) == 0 {
		return false
	}
	for _, r := range v {
		if !IsTokenRune(r) {
			return false
		}
	}
	return true
}
