#include <stdio.h>
#include <stdlib.h>
#include <time.h>
//정보통신공학부 2020037063 차동현
void print_array(char *str, int array[],int size);
int binary_search(int array[],int size,int key);

int main()
{
int size=5;
int data[size-1];
srand(time(NULL));
for(int n=0;n<size;n++)
data[n]=(rand()%4)+1;

int index,key=1;
print_array("원본 :",data,size);
while(key>0)
{
    printf("검색할 수를 입력하세요:");
    scanf("%d",&key);
    if(key>0)
    {
        int index=binary_search(data,size,key);
        printf("%d를 탐색합니다. \n ",key);
        if(index!=-1)
        printf("데이터가 %d번째에 있습니다.\n",index+1);
        else
        printf("데이터가 파일에 없습니다.\n");
        
    }
}
printf("프로그램을 종료합니다.\n");
}

void print_array(char *str, int array[],int size)
{
    int n;
    printf("%s %d\n",str,size);
    for (n=0; n<size; n++)
    {
        printf("%5d",array[n]);
        if((n+1)%10==0)
        printf("\n");
    }
    printf("\n");
}

int binary_search(int array[],int size,int key)
{
    int n;
    for(n=0;n<size;n++)
    {
        if(key==array[n])
        return n;
        if(key<array[n])
        break;
    }
    return -1;
}