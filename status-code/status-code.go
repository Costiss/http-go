package statuscode

// "100"  ; Section 10.1.1: Continue
// | "101"  ; Section 10.1.2: Switching Protocols
// | "200"  ; Section 10.2.1: OK
// | "201"  ; Section 10.2.2: Created
// | "202"  ; Section 10.2.3: Accepted
// | "203"  ; Section 10.2.4: Non-Authoritative Information
// | "204"  ; Section 10.2.5: No Content
// | "205"  ; Section 10.2.6: Reset Content
// | "206"  ; Section 10.2.7: Partial Content
// | "300"  ; Section 10.3.1: Multiple Choices
// | "301"  ; Section 10.3.2: Moved Permanently
// | "302"  ; Section 10.3.3: Found
// | "303"  ; Section 10.3.4: See Other
// | "304"  ; Section 10.3.5: Not Modified
// | "305"  ; Section 10.3.6: Use Proxy
// | "307"  ; Section 10.3.8: Temporary Redirect
// | "400"  ; Section 10.4.1: Bad Request
// | "401"  ; Section 10.4.2: Unauthorized
// | "402"  ; Section 10.4.3: Payment Required
// | "403"  ; Section 10.4.4: Forbidden
// | "404"  ; Section 10.4.5: Not Found
// | "405"  ; Section 10.4.6: Method Not Allowed
// | "406"  ; Section 10.4.7: Not Acceptable
// | "407"  ; Section 10.4.8: Proxy Authentication Required
// | "408"  ; Section 10.4.9: Request Time-out
// | "409"  ; Section 10.4.10: Conflict
// | "410"  ; Section 10.4.11: Gone
// | "411"  ; Section 10.4.12: Length Required
// | "412"  ; Section 10.4.13: Precondition Failed
// | "413"  ; Section 10.4.14: Request Entity Too Large
// | "414"  ; Section 10.4.15: Request-URI Too Large
// | "415"  ; Section 10.4.16: Unsupported Media Type
// | "416"  ; Section 10.4.17: Requested range not satisfiable
// | "417"  ; Section 10.4.18: Expectation Failed
// | "500"  ; Section 10.5.1: Internal Server Error
// | "501"  ; Section 10.5.2: Not Implemented
// | "502"  ; Section 10.5.3: Bad Gateway
// | "503"  ; Section 10.5.4: Service Unavailable
// | "504"  ; Section 10.5.5: Gateway Time-out
// | "505"  ; Section 10.5.6: HTTP Version not supported
const (
	Continue                     = 100
	SwitchingProtocols           = 101
	OK                           = 200
	Created                      = 201
	Accepted                     = 202
	NonAuthoritativeInformation  = 203
	NoContent                    = 204
	ResetContent                 = 205
	PartialContent               = 206
	MultipleChoices              = 300
	MovedPermanently             = 301
	Found                        = 302
	SeeOther                     = 303
	NotModified                  = 304
	UseProxy                     = 305
	TemporaryRedirect            = 307
	BadRequest                   = 400
	Unauthorized                 = 401
	PaymentRequired              = 402
	Forbidden                    = 403
	NotFound                     = 404
	MethodNotAllowed             = 405
	NotAcceptable                = 406
	ProxyAuthenticationRequired  = 407
	RequestTimeOut               = 408
	Conflict                     = 409
	Gone                         = 410
	LengthRequired               = 411
	PreconditionFailed           = 412
	RequestEntityTooLarge        = 413
	RequestURITooLarge           = 414
	UnsupportedMediaType         = 415
	RequestedRangeNotSatisfiable = 416
	ExpectationFailed            = 417
	InternalServerError          = 500
)

func GetStatusLine(statusCode int) (int, string) {
	switch statusCode {
	case 200:
		return 200, "OK"
	case 404:
		return 404, "Not Found"
	case 500:
		return 500, "Internal Server Error"
	}
	panic("invalid status code")
}
