package comp

// #include <stdlib.h>
// #include <stdio.h>
// #include <hs/hs.h>
// #include <hs/hs_common.h>
// #include <hs/hs_compile.h>
import "C"

type cmp struct {
	rgx *interface{}
	err *error
}

//Comp is a simple compile with default setup, for custom use Compile
func Comp(pattern string) {

    expression := C.CString()

    hs_error_t C.hs_compile(expression, HS_FLAG_ALLOWEMPTY | HS_FLAG_UTF8 | HS_FLAG_MULTILINE, unsigned int mode, const hs_platform_info_t * platform, hs_database_t ** db, hs_compile_error_t ** error)

    
}
