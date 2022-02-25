#include <iostream>
#include <cstdio>
#include <string>
using namespace std;

int main(){
    int tc;
    string result[]={"NA","MMYY","YYMM","AMBIGUOUS"};
    int t;
    scanf("%d",&t);

    for(tc = 1; tc<= t; ++tc) {
    char s[4];
    int r_n=0;
    scanf("%s",s);
    int fN=10*(s[0]-'0') +(s[1]-'0');
    int bN=10*(s[2]-'0') +(s[3]-'0');

    if (fN<13 && bN>0) r_n+=1;
    if (bN<13 && fN>0) r_n+=2;

    printf("#%d %s\n",tc,result[r_n].c_str());
    }
}