package qsort

//typedef int (*qsort_cmp_func_t)(const void *a, const void *b);
import "C"

import "unsafe"

func Sort(base unsafe.Pointer, 
			num, 
			size C.size_t,
			cmp C.qsort_cmp_func_t,){
		C.qsort_cmp_func_t(base, num, size, cmp)
}