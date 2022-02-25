#include<iostream>

using namespace std; 

int main() {
int rows;
int t,tc;
cin>>tc;
for (t=1; t<=tc; t++){
cin>>rows;
cout<<"#"<<t<<endl;
for (int i = 0; i < rows; i++) {
    int val = 1; 
    for (int k = 0; k <= i; k++) {
        cout<<val<<" "; 
        val = val * (i - k) / (k + 1); 
        }
        cout<<endl;
    } 
} 
return 0; 
}