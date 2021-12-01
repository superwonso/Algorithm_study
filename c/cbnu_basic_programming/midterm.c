#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <math.h>

int primecheck(int a) // 소수 판별 함수
{
    int from,to; 
    to=(int)sqrt((double)a); // 형 변환과 마지막에 나눌 값
    from=2;
    while (from<=to) // 2 부터 마지막에 나눌 값까지
    {
        if((a%from)==0) // 나머지가 0일때는
        {
            return 0; // 소수 아님
        }
        from +=1;
    }
    return 1; // 나머지가 0이 아닐때는 소수임
}

int main() // 메인 함수
{
    int prime=primecheck(int randomnum); // 소수 판단
    int randomnum=srand((unsigned int)TIME(NULL));
    int random=randomnum;
    while (1) // 무한 루프
    {
       if(rand>=2)
       {
        printf("%d에 대해 소수인가를 판단하겠습니다.",random);
        if(prime==1) // randomnum이 소수이면
        printf("%d는 소수입니다.\n",randomnum);
        else // randomnum이 소수가 아니면
        printf("%d는 소수가 아닙니다.\n",randomnum);
       }
        else
        {
        printf("소수를 판단할 수 없습니다.");
        }
        
    }
}