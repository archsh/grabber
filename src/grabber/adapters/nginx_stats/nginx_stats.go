/** Package: nginx_stats */
package nginx_stats 

type NginxStat struct {
	conns uint
	
}

func get_nginx_stats (url string) (stat NginxStat, e error){
	return
}