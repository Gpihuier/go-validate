package go_validate

import "regexp"

const (
	integer   = `^[-+]?[0-9]+$`
	float     = `^(?:[-+]?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$`
	number    = `^[-+]?[0-9]+(?:\.[0-9]+)?$`
	email     = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	alpha     = `^[A-Za-z]+$`
	alphaNum  = `^[A-Za-z0-9]+$`
	alphaDash = `^[A-Za-z0-9\\-\\_]+$`
	mobile    = `^1[3-9]\\d{9}$`
	idCard    = `(^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$)|(^[1-9]\\d{5}\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}$)`
	zip       = `\\d{6}`
)

var (
	integerRegex   = regexp.MustCompile(integer)
	floatRegex     = regexp.MustCompile(float)
	numberRegex    = regexp.MustCompile(number)
	emailRegex     = regexp.MustCompile(email)
	alphaRegex     = regexp.MustCompile(alpha)
	alphaNumRegex  = regexp.MustCompile(alphaNum)
	alphaDashRegex = regexp.MustCompile(alphaDash)
	mobileRegex    = regexp.MustCompile(mobile)
	idCardRegex    = regexp.MustCompile(idCard)
	zipRegex       = regexp.MustCompile(zip)
)
