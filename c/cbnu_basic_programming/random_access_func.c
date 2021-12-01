#include <stdio.h>
#include <stdlib.h>
//정보통신공학부 2020037063 차동현
#define SIZE 100

void init_table(int table[],int size);

int main(void)
{
    int table[SIZE];
    int n,data;
    long pos;
    FILE *fp=NULL;
    init_table(table,SIZE);
    if( (fp=fopen("sample.dat","wb"))==NULL)
    {
        fprintf(stderr,"출력을 위한 파일을 열 수 없습니다.\n");
        exit(1);
    }
    fwrite(table,sizeof(int),SIZE,fp);
    fclose(fp);
    if((fp=fopen("sample.dat","rb"))==NULL)
    {
        fprintf(stderr,"입력을 위한 파일을 열 수 없습니다.\n");
        exit(1);
    }
    while(1)
    {
        printf("파일에서의 위치를 입력하세요(단, 0에서 %d, 종료 -1)\n",SIZE-1);
        scanf("%d",&n);
        if(n==-1) break;
        pos = (long)n*sizeof(int);
        fseek(fp,pos,SEEK_SET);
        fread(&data,sizeof(int),1,fp);
        printf("%d 위치의 값은 %d입니다.\n",n,data);
    }
    fclose(fp);
    return 0;
}

void init_table(int table[],int size)
{
    int i;
    for(i=0; i<size; i++)
    table[i]=i*i;
}