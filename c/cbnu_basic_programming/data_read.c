#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//정보통신공학부 2020037063 차동현
int main(void)
{
    FILE *fp=NULL;
    char fname[30];
    float total=0.0;
    int math,kor,eng=0;
    printf("파일 입력을 입력하세요 : ");
    scanf("%s",fname);
    if ((fp=fopen(fname,"w"))==NULL)
    {
        printf("성적 파일 %s를 열 수 없습니다.\n",fname);
        exit(1);
    }
     while(1)
    {
    printf("수학 성적을 입력해주세요 : ");
    scanf("%d",&math);
    printf("국어 성적을 입력해주세요 : ");
    scanf("%d",&kor);
    printf("영어 성적을 입력해주세요 : ");
    scanf("%d",&eng);
    fprintf(fp, "%d %d %d\n",math,kor,eng);
    break;
    }
    
    fclose(fp);
    
    if((fp=fopen(fname,"r"))==NULL)
    {
        fprintf(stderr,"성적파일 %s를 열 수 없습니다.\n",fname);
        exit(1);
    }
    while(!feof(fp))
    {
        fscanf(fp,"%d %d %d",&math,&kor,&eng);
        total=(math+kor+eng);
    }
    printf("평균 = %f\n",total/3);
    fclose(fp);
return 0;
}