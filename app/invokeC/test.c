#include <stdio.h>
#include <stdlib.h>

void print(char *str) {
  printf("test.c print(char *str)\n");
  printf("%s\n", str);
}
void print1(char *str) {
  printf("test.c print1(char *str)\n");
  printf("%s\n", str);
}
void print2(char *str)
{
  printf("test C testPrint %s\n",str);
}
void ttargbyte(char *dst, char *src){
  printf("test C argbyte\n");
  printf("dst:%s,src:%s\n",dst, src);
}

int isboolok(){
  return 1;
}
