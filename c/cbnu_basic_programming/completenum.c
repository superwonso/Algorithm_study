#include <stdio.h>

int calc(int comp);
int main(void)
{
    int num,input;
    printf("몇 이하의 완전수를 계산할까요?\n");
    scanf("%d",&input);
    for(num=1;num<input;num++)
{
    if(calc(num)==1)
    printf("%d",num);
}
return 0;
}

int calc(int comp)
{
    int a;
    int sum=0;
    for(a=1; a<comp; a++)
    {
        if(comp%a==0)
        {
            sum=sum+a;
        }
        else if(sum==a)
        {
            return 1;
        }
        else return 0;
    }
}