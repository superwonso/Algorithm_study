#include <stdio.h>
#include <time.h>
#include <stdlib.h>
#include <windows.h>
//정보통신공학부 2020037063 차동현
void swap(int *a,int *b);
void start(); // 전역변수 리셋
void printarr(int(*arr)[5]); // 숫자판
void computerselect();//숫자 선택,컴퓨터
void userselect();//숫자 선택,유저 
void selectednumber();//숫자 처리
int checkwinner(int (*arr)[5]);//승패 확인

int user[5][5];//사용자 숫자판
int computer[5][5];//컴퓨터 숫자판
int checked[25];//중복 숫자 저장
int count;//중복 숫자 개수

int main()
{
    int i,j,turn=0;
    start();
    do{
    if (turn%2==0)// 사용자 턴
    {
    system("cls");
    printarr(user);
    printf("\n 숫자를 선택하세요:");
    userselect(); // 사용자 숫자 선택
    selectednumber(); // 선택된 숫자의 처리
    system("cls"); // cmd cls... 리눅스나 맥에서는 "clear"
    printarr(user);
    printf("\n 사용자가 [%d]를 선택하였습니다. \n",checked[count-1]);
    turn ++;
    }
    else  // 컴퓨터 턴
    {
        computerselect(); // 컴퓨터 숫자 선택
        selectednumber(); // 선택된 숫자의 처리
        printf("컴퓨터가 [%d]를 선택했습니다. \n",checked[count-1]);
        system("pause");
        turn ++;
    }
    } while ( (checkwinner(computer)+checkwinner(user)) ==0);
    system("cls");
    printf("사용자 : \n");
    printarr(user);
    printf("\n컴퓨터 : \n");
    printarr(computer);
    if ((checkwinner(computer) + checkwinner(user)) == 2) 
    {
        printf("무승부 입니다.\n");
    }
    else
    {
        if (checkwinner(computer) == 1)
        printf("컴퓨터가 이겼습니다.\n");
        else
        printf("사용자가 이겼습니다.\n");
    }
return 0;
}

void start()
{
    int i,j,k;
    int randarr[25];//임시 배열
    srand(time(NULL));
    for(i=0; i<25; i++)
    {
        randarr[i] = i+1; // 임시 배열에 1~25의 수를 넣음
    }

    //유저 숫자판 초기화
    for (i=0; i<25; i++)
    {
        swap(&randarr[i],&randarr[rand()%25]);
    }
    for(i=0,k=0; i<5; i++)
    {
        for(j=0;j<5;j++)
        user[i][j]=randarr[k++];
    }
    //컴퓨터 숫자판 초기화
    for (i=0; i<25; i++)
    {
        swap(&randarr[i],&randarr[rand()%25]);
    }
    for(int m=0,n=0; m<5; m++)
    {
        for(int o=0;o<5;o++)
        computer[m][o]=randarr[n++];
    }
    //숫자 저장판 초기화
    for(i=0; i<25; i++)
    checked[i]=0;
    count =0;
}

void swap(int *a,int *b)
{
    int temp;
    temp=*a;
    *a=*b;
    *b=temp;
}

void printarr(int (*arr)[5])
{
    int i,j;
    for(i=0; i<5; i++)
    {
        for (j=0; j<5; j++)
        {
            if (arr[i][j]!=-1)
            printf("%5d",arr[i][j]);
            else
            printf("%5c",'X');
        }
        printf("\n");
    }
}

void computerselect() // 컴퓨터 숫자 선택
{
    int i,computerselectnum,flag;
    while(1)
    {
        computerselectnum=(rand()%25)+1;//숫자를 임의로 발생시키고
        flag=0;
        for(i=0;i<count;i++)
        {
        if (checked[i] == computerselectnum)//이미 존재하면 표시
        flag=1;
        }

        if(flag==0)//존재하지 않으면
        {
            checked[count++]=computerselectnum;
            //입력되는 숫자는 checked 배열의 끝에 추가
            break;
        }
    }
}

void userselect()//유저의 숫자 선택
{
    int i,userselectnum,flag;
    while (1)
    {
        scanf("%d",&userselectnum);
        flag=0; // 플레그 초기화
        if(userselectnum<0 || userselectnum>25)
        flag=1;

        for(i=0;i<count;i++) // 이미 존재하는지 확인
        {
        if(checked[i]==userselectnum)
        flag=1;
        }
        if(flag==0)//존재하지 않으면
        {
            checked[count++]=userselectnum; //숫자를 넣고 count 증가
            //입력되는 숫자는 checked 배열의 끝에 추가
            break;
        }
        printf("\n 알맞은 숫자를 선택하세요:");
    }
}

void selectednumber()
{
    int i,j;
    //컴퓨터 숫자판 처리
    for (i= 0; i< 5; i++)
    for (j = 0; j < 5; j++)
    if (computer[i][j] == checked[count -1])
    computer[i][j] = -1;
    
    //사용자 숫자판 처리
    for (i= 0; i< 5; i++)
    for (j = 0; j < 5; j++)
    if (user[i][j] == checked[count -1])
    user[i][j] = -1;
}

int checkwinner(int(*arr)[5])
{
    int i,j,bingo,win=0;
    //행 확인
    for(i=0;i<5;i++)
    {
        bingo=0;
        for(j=0; j<5; j++)
        if(arr[i][j]==-1)
        bingo++;

        if (bingo==5) // 한 행에 5개가 존재하면
        {
            win=1;
            break;
        }
    }
    //열 확인
    for(j=0;j<5;j++)
    {
        bingo=0;
        for(i=0; i<5; i++)
        if(arr[i][j]==-1)
        bingo++;
        if (bingo==5) // 한 열에 5개가 존재하면
        {
            win=1;
            break;
        }
    }
    //대각선 확인,좌측 위 시작
    bingo=0;
    for(i=0,j=0; i<5&&j<5; i++,j++)
    {
        if(arr[i][j]==-1)
        bingo++;
        if(bingo==5)//5개가 존재하면
        win=1;
    }
    
    //대각선 확인,우측 위 시작
    bingo=0;
    for(i=0,j=4; i<5&&j>-1; i++,j--)
    {
        if(arr[i][j]==-1)
        bingo++;
        if(bingo==5)//5개가 존재하면
        win=1;
    }
return win;
}