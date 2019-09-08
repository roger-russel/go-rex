package gorex

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include <regex.h>
// int match(const char *pattern, const char *subject) {
//     regex_t re;
//     if (regcomp(&re, pattern, REG_EXTENDED|REG_NOSUB) != 0) return 0;
//     int status = regexec(&re, subject, 0, NULL, 0);
//     regfree(&re);
//     if (status != 0) return 0;
//     return 1;
// }
import "C"
import (
	"unsafe"
)

//Match match a string and a subject
func Match(pattern string, subject string) bool {

	cpattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cpattern))

	csubject := C.CString(subject)
	defer C.free(unsafe.Pointer(csubject))

	ptr := C.malloc(C.sizeof_int * 1024)
	defer C.free(unsafe.Pointer(ptr))

	result := C.match(cpattern, csubject)

	return result == 0
}
