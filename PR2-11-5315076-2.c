#include <stdio.h>
#include <stdlib.h>

int main() {
  FILE *file = fopen("abc.txt", "r");
  int c;
  while((c = getc(file)) != EOF) {
    putchar(c);
  }
  fclose(file);
  return 0;
}
