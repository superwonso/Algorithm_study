#include <stdio.h>
int factorial(int num);
int main()
{
int num;
printf("팩토리얼 연산을 실시할 수를 입력해주세요.(n!) :");
scanf("%d",&num);
printf("%d! = %d 입니다.\n",num,factorial(num));

return 0;
}

int factorial(int num)
{
int result;

    if(num==1)
    {
        result=1;
    } 
    else result=num*factorial(num-1);

return result;
}
