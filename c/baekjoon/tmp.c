#include <stdio.h>
#include <stdbool.h>

bool is_he_cbnu_student(char* name);
bool is_he_telecommunication_student(char* name);
void c_b_v(int a,int b);
void c_b_r(int *a,int *b);
int main() {

    int my_age=22;
    int korean_age=my_age+1;

    int my_IQ = -138;
    int your_IQ = +138;

    your_IQ++;
    your_IQ--;
    ++your_IQ;

    printf("%d\n", your_IQ);
    if(my_IQ>your_IQ)
        printf("My IQ is bigger than your IQ\n");
    else if(my_IQ<your_IQ)
        printf("My IQ is smaller than your IQ\n");
    else if(my_IQ==your_IQ)
        printf("My IQ is equal to your IQ\n");

int x=0;
int y=200;
int z=300;
int w = x ? y : z;

int ex1,ex2;
ex1=x*y,ex2=x+y;
printf("%d %d\n",ex1,ex2);

    if((your_IQ&1)==0) // 139(10) = 10001011(2)
        printf("Your IQ is even\n");
    else
        printf("Your IQ is odd\n");


char *name = "cdh";
if(is_he_cbnu_student(name)==true) {
    if(is_he_telecommunication_student(name)==true)
        printf("He is a telecommunication student\n");
    else
        printf("He is not a telecommunication student\n");
} else if(is_he_cbnu_student(name)==false) {
    printf("He is not a student of cbnu\n");
}

printf("%d\n",my_IQ);

switch_start:
switch(my_IQ) {
    case +138:
        printf("My IQ is positive\n");
        break;
    case -138:
        printf("My IQ is negative\n");
    default:
        printf("Shall we study?\n");
        my_IQ+=138;
        goto switch_start;


}
for (int i=1; i<=5; ++i) {
    for(int j=i; j<=5;j++) printf("*");
    printf("\n");
}

int index=0;
do 
{
    printf("do-while \n");
    index++;
} while(index<10);

index=0;
while(1) {
    printf("while \n");
    index++;
    if(index==10) break;
}
for(int i=0; i<6; i++) {
to_static_var();
}
}
bool is_he_cbnu_student(char* name) {
    if(*name>0)
        return true;
    else
        return false;
}

bool is_he_telecommunication_student(char* name) {
    if(*name>0)
        return true;
    else
        return false;
}

void c_b_v(int a,int b) {
a=1,b=2;
int temp=a;
a=b;
b=temp;
}

void c_b_r(int *a,int *b) {
    a=1,b=2;
    int temp=a;
    a=b;
    b=temp;
}

void to_static_var(){
    static int age_of_jiwon=18;
    printf("%d\n",age_of_jiwon);
    age_of_jiwon++;
}