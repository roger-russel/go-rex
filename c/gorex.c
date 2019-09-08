#include <stdlib.h>
#include <regex.h>
#include <sys/types.h>
#include <stdio.h>

#define DEFAULT_MAX_MATCHES 1 //The maximum number of matches allowed in a single string
#define FALSE 0
#define TRUE 1

typedef struct sct_result {
    int match;
    const char *error;
} sct_result;

typedef struct sct_comp {
    regex_t rgx;
    const char *error;
} sct_comp;

sct_comp comp(const char *pattern){

    regex_t re;
    
    int rc = regcomp(&re, pattern, REG_EXTENDED|REG_NOSUB);

    if (rc != 0) {

        size_t errcomp = regerror(rc, &re, (char *)NULL, (size_t)0);
        char * errcomp_str = malloc(errcomp * sizeof(char*));
        regerror(rc, &re, errcomp_str, errcomp);

        sct_comp response = {
            rgx: re,
            error: errcomp_str
        };

        return response;
    }

    sct_comp response = {
        rgx: re,
        error: NULL
    };

    return response;
}

sct_result test(const char *pattern, const char *subject) {
    
    sct_comp r = comp(pattern);

    if ( r.error != NULL ) {

        sct_result response = {
            match: FALSE,
            error: r.error
        };

        return response;
    }

    int status = regexec(&r.rgx, subject, 0, NULL, 0);
    regfree(&r.rgx);

    sct_result response = {
        match: status == 0, // If status is equal to zero, then regexec matched the result
        error: NULL,
    };

    return response;

}

regmatch_t match(regex_t *pexp, char *subject) {
	regmatch_t matches[DEFAULT_MAX_MATCHES]; //A list of the matches in the string (a list of 1)
	//Compare the string to the expression
	//regexec() returns 0 on match, otherwise REG_NOMATCH
	regexec(pexp, subject, DEFAULT_MAX_MATCHES, matches, 0);
	
    return *matches;
	
}

int main() {

    char * pattern = "^([[:alpha:]]+)([[:digit:]]+)$";
    char * subject = "abc123";

    sct_result result = test(pattern, subject);

    printf("%s -> %s = %d | %s \n", pattern, subject, result.match, result.error);
}

void  trash(){    
	int rv;
	regex_t exp; //Our compiled expression
	//1. Compile our expression.
	//Our regex is "-?[0-9]+(\\.[0-9]+)?". I will explain this later.
	//REG_EXTENDED is so that we can use Extended regular expressions
	rv = regcomp(&exp, "([[:alpha:]]+)([[:digit:]])", REG_EXTENDED);
	if (rv != 0) {
		printf("regcomp failed with %d\n", rv);
	}
	//2. Now run some tests on it
	regmatch_t m1 = match(&exp, "a0");



	//3. Free it
	regfree(&exp);
	//return 0;
}