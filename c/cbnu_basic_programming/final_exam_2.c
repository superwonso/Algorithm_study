#include <stdio.h>
#include <ctype.h>
//정보통신공학부 2020037063 차동현
//기말고사 10번
int alphabet[26];
void main(void) {
   char buffer[256], * pch;
   char ch;
   int n;
   char temp[100];
   FILE* fp = fopen("C:/Users/user/Desktop/b_b/Clang/hw/malloc.txt", "rt"); //쓰기모드 파일 열기
   while(1)
   {
      printf("영어 단어를 입력해주세요. 공백을 입력하면 종료됩니다.");
      fscanf(fp,"%s",temp); // 영단어 입력받아서 파일에 쓰기
      if(temp==EOF) break; // 공백일 때 종료
      fclose(fp);
   }
    FILE* fp = fopen("C:/Users/user/Desktop/b_b/Clang/hw/malloc.txt", "r"); //쓰기모드 파일 열기
   for (n = 0;n < 26;n++)
      alphabet[n] = 0;
   while (fgets(buffer, 255, fp) != NULL)
   {
      pch = buffer;
      while (*pch !=NULL)
      {
         if (isupper(*pch))
            alphabet[*pch - 'A'] += 1;
         else if (islower(*pch))
            alphabet[*pch - 'a'] += 1;
         pch++;
      }
   }
   fclose(fp);
   printf("파일의 영문자수\n\n");
   for (n = 0;n < 26;n++)
    printf("%c: %d\t", n + 'A', alphabet[n]);
}