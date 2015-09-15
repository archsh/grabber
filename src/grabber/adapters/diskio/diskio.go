/** Package: diskio */
package diskio 

type IOStatInfo struct {
	part string
	
}

type DFInfo struct {
	part string
	available uint64
	used uint64
	total uint64
}

func get_iostat_info() (info IOStatInfo, e error){
	return
}

func get_df_info()(info DFInfo, e error){
	return
}