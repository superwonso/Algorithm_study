#include <stdio.h>
//정보통신공학부 2020037063 차동현
#define N 40
int main()
{
    char dept[N],name[N];
    int grade=0,joiny=0;
//input

printf("정보를 입력해주세요.\n\n");

printf("이름? ");
scanf("%s",name);

printf("학과는? ");
scanf("%s",dept);

printf("입학 년도는? ");
scanf("%d",&joiny);
int claof=joiny%100;

printf("학년은? ");
scanf("%d",&grade);

//output

printf("\n저는 %d학번 %d학년 %s %s입니다.",claof,grade,dept,name);

return 0;
}