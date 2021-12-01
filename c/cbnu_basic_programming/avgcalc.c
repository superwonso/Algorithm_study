#include <stdio.h> 
#include <math.h>
#define MAX 10
//정보통신공학부 2020037063 차동현
double num[MAX];
int n;
void get_data();
void get_avg_devi(double *pavg,double *psqrdevi);
int main()
{
double avg, devi;
get_avg_devi(&avg,&devi);
printf("평균은 %2.2lf 입니다.\n",avg);
printf("표준편차는 %2.2lf 입니다.",devi);
}

void get_data()
{
printf("10개 이하의 실수를 입력해주세요.\n");
printf("입력을 종료하시려면 0 이하의 숫자를 입력해주세요.\n");
n=0;
while(n<10)
{
    printf("%d번째 숫자를 입력하세요.",n+1);
    scanf("%f",&num[n]);
    if (num[n]>0.0)
        n+=1;
    else
        break;
}
}

void get_avg_devi(double *pavg,double *psqrdevi)
{
double sum,vari,sumdouble,sqrvari,avg;
int x;
sum=sumdouble=0.0;
get_data();
for(x=0;x<n;x++)
{
    sum+=num[x];
    sumdouble=num[x]*num[x];
}

avg=sum/(double)n;
vari=sumdouble/(double)n - avg*avg;
sqrvari=sqrt(vari);

*pavg=avg;
*psqrdevi=sqrvari;
}
