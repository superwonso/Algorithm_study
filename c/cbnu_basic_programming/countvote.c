#include <stdio.h>
#define max 5
//정보통신공학부 2020037063 차동현
int main(void)
{
int ip[max];
int n=0;
int xdata=0;

    for(int i=0; i<max; i++) //ip[] data reset
    {
        ip[i]=0;
    }  
printf("1번에서 4번중에 하나를 고르세요(종료하려면 숫자 외의 다른 문자를 넣어주세요) : ");
    while(scanf("%d",&xdata)==1)
    {
        n++;
        if(xdata<1 || xdata>=max)
        {
            ++ip[0];
        }
        else
        {
            ++ip[xdata];
        } 
        printf("1번에서 4번중에 하나를 고르세요 : ");
    }
    printf("\n득표수\n");
        for(int i=1; i<max; i++)
        {
        printf("%d번 : %d표\n",i,ip[i]);
        }
    printf("무효표 : %d표\n",ip[0]);
    printf("총 투표수 : %d표\n",n);
}