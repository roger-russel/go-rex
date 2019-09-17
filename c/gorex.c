#include <errno.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <hs/hs.h>

typedef struct comp
{
  hs_database_t *database;
  const char *error;
} comp;

comp Comp(const char *pattern)
{
  hs_database_t *database;
  hs_compile_error_t *compile_err;
  if (
      hs_compile(
        pattern,
        HS_FLAG_DOTALL,
        HS_MODE_BLOCK,
        NULL,
        &database,
        &compile_err)
      != HS_SUCCESS)
  {

    comp error = {
      error : strdup(compile_err->message)
    };

    hs_free_compile_error(compile_err);

    return error;
  }

}

void main()
{
  comp cmp = Comp("/a(}bc/");
}
