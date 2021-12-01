#include <stdio.h>
//정보통신공학부 2020037063 차동현
// 기말고사 9번
struct students_info // 학생 정보의 구조체 선언
{
    char name[20]; // 이름
    double kor; // 국어 점수
    double eng; // 영어 점수
    double math; // 수학 점수
};


int main()
{
double avg_kor,avg_eng,avg_math=0; // 국어,영어,수학 평균 실수형 변수 선언 
struct students_info s[2]; //구조체 배열
for (int x=0; x<3; x++) // s[0],s[1],s[2]에 정보 입력
{
    printf("학생의 이름, 국어 점수, 영어 점수, 수학 점수를 순서대로 적어주세요\n");
    scanf("%s %lf %lf %lf",s[x].name,&s[x].kor,&s[x].eng,&s[x].math);
}
avg_kor = (s[0].kor+s[1].kor+s[2].kor)/3;
avg_eng = (s[0].eng+s[1].eng+s[2].eng)/3;
avg_math = (s[0].math+s[1].math+s[2].math)/3;
printf("국어 점수의 평균점수는 %.02lf 입니다.\n",avg_kor);
printf("영어 점수의 평균점수는 %.02lf 입니다.\n",avg_eng);
printf("수학 점수의 평균점수는 %.02lf 입니다.\n",avg_math);
}