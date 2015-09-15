/** Package: nginx_log */
package nginx_log 

type RequestInfo struct {
	timestamp string
	clientIp string
	requestUri string
	userAgent string
	responseCode uint16
	requestTime uint16
	bytes uint
}
