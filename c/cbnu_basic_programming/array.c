#include <stdio.h>
#include <stdlib.h>
#include <time.h>
//버블 정렬 예제, 2020037063 차동현

void generate_random(int array[], int size);
void point_array(char *str,int array[], int size);
void bubble_sort(int array[],int size);
void bubble(int array[],int last);
void swap(int *xp,int *yp);

int main(void)
{
int *data, size;
srand(time(NULL));
printf("배열의 크기를 입력하세요 : ");
scanf("%d",&size);
data = (int *)malloc(sizeof(int)*size);
generate_random(data,size);
point_array("원본 : ",data,size); // 버블 정렬
bubble_sort(data,size);
point_array("정렬 : ",data,size); // 정렬 데이터 출력
//파일에 저장
free(data);
return 0;
}

void generate_random(int array[],int size)
{
    int n;
    for(n=0; n<size; n++)
        array[n]=(rand()%1000)+1;
}

void point_array(char *str,int array[],int size)
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

void bubble_sort(int array[],int size)
{
    int n;
    for(n=0; n<size-1; n++)
    bubble(array,size-n-1);
}

void bubble(int array[],int last)
{
    int m;
    for(m=0; m<last; m++)
    if(array[m]>array[m+1])
    swap(&array[m],&array[m+1]);
}

void swap(int *xp,int *yp)
{
int tmp=*xp;
*xp=*yp;
*yp=tmp;
}