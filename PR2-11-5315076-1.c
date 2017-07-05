#include <stdio.h>
#include <stdlib.h>

struct player {
  char name[30];
  double ave;
  int homerun;
};

int main(void) {
  FILE *fp;
  struct player batter[3] = {
    {"Maru", 0.336, 13},
    {"Oshima", 0.335, 1},
    {"Miyazaki", 0.333, 6}
  };
  if ((fp = fopen("abc.txt", "w")) == NULL) {
    printf("Can't open file.\n");
    exit(1);
  }
  for(int i = 0; i < 3; i++) {
    fprintf(fp, "%s %f %d \n", batter[i].name, batter[i].ave, batter[i].homerun);
  }
  fclose(fp);
  return 0;
}
