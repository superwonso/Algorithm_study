#include <stdio.h>
#include <math.h>
//정보통신공학부 2020037063 차동현
int check_primenum(int num);
int main()
{
    int num;
    int flag;
    printf("소수를 판정할 정수를 입력하세요.");
    scanf("%d",&num);
    if(num>1)
    {
        flag=1; // for checking
    }
    else flag=0;

    switch(flag)
    {
        case 0:
        {
            printf("소수를 판별할 수 없습니다.\n");
        }
        case 1:
        {
            printf("%d에 대한 소수를 판별하겠습니다.\n",num);
            check_primenum(num);
            if(check_primenum(num)==1) //if 'num' is the prime number, return value of check_primenum is 1
            {
                printf("%d는 소수입니다.\n",num);
            }
            else printf("%d는 소수가 아닙니다.\n",num);    
        }

    }
    return 0;
}

int check_primenum(int num)
{
    int prime;
    int from,to;
    to=(int)sqrt((double)num);
    
    for(from=2; from<=to; from++)
    {
        if( (num%to) == 0)
        {
            prime=0;    
            break;
        }
        else
        {
            prime=1;
        }
    }
return prime;
}