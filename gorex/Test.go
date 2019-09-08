package gorex

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include <regex.h>
//
// #define DEFAULT_MAX_MATCHES 1 //The maximum number of matches allowed in a single string
// #define FALSE 0
// #define TRUE 1
//
// typedef struct sct_result {
//    int match;
//    const char *error;
// } sct_result;
//
// typedef struct sct_comp {
//    regex_t rgx;
//    const char *error;
// } sct_comp;
//
// sct_comp comp(const char *pattern){
//
//    regex_t re;
//
//    int rc = regcomp(&re, pattern, REG_EXTENDED|REG_NOSUB);
//
//    if (rc != 0) {
//
//        size_t errcomp = regerror(rc, &re, (char *)NULL, (size_t)0);
//        char * errcomp_str = malloc(errcomp * sizeof(char*));
//        regerror(rc, &re, errcomp_str, errcomp);
//
//        sct_comp response = {
//            rgx: re,
//            error: errcomp_str
//        };
//
//        return response;
//    }
//
//    sct_comp response = {
//        rgx: re,
//        error: NULL
//    };
//
//    return response;
// }
//
// sct_result test(const char *pattern, const char *subject) {
//
//    sct_comp r = comp(pattern);
//
//    if ( r.error != NULL ) {
//
//        sct_result response = {
//            match: FALSE,
//            error: r.error
//        };
//
//        return response;
//    }
//
//    int status = regexec(&r.rgx, subject, 0, NULL, 0);
//    regfree(&r.rgx);
//
//    sct_result response = {
//        match: status == 0, // If status is equal to zero, then regexec matched the result
//        error: NULL,
//    };
//
//     return response;
//
// }
import "C"
import (
	"errors"
	"unsafe"
)

//Test match a string and a subject
func Test(pattern string, subject string) (bool, error) {

	cpattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cpattern))

	csubject := C.CString(subject)
	defer C.free(unsafe.Pointer(csubject))

	ptr := C.malloc(C.sizeof_int * 1024)
	defer C.free(unsafe.Pointer(ptr))

	result := C.test(cpattern, csubject)
	return result.match == 1, errors.New(C.GoString(result.error))
}
