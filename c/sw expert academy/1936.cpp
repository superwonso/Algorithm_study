#include <iostream>
#include <vector>

using namespace std;

int main(){
    int a,b;
    ios_base::sync_with_stdio(false);
    cin>>a>>b;
    if(a==1&&b==2) {
        cout<<"B";
    } 
    if(a==2&&b==3){
        cout<<"B";
    }
    if(a==1&&b==3){
        cout<<"A";
    }
    if(a==2&&b==1){
        cout<<"A";
    }
    if(a==3&&b==1){
        cout<<"B";
    }
    if(a==3&&b==2){
        cout<<"A";
    }
}