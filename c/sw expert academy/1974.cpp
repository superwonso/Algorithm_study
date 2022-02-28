#include <iostream>
#include <vector>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int t, tc;
    cin >> tc;
    int input[9][9];
    for(t=1; t<=tc; t++){
        for(int i=0; i<9; i++){
            for (int j=0; j<9; j++){
                    cin >> input[i][j];
            }
        }
    bool flag = true;
    for(int i=0; i<9; i++){
        for(int j=0; j<9; j++){
            if(input[i][j]!=0){
                for(int k=0; k<9; k++){
                    if(k!=j && input[i][k]==input[i][j]){
                        flag = false;
                    }
                }
                for(int k=0; k<9; k++){
                    if(k!=i && input[k][j]==input[i][j]){
                        flag = false;
                    }
                }
            }
        }
    }
    for(int i = 0; i < 3; i++) { 
        if(!flag) break; 
        for(int j = 0; j < 3; j++) { 
            int check[9]={0,}; 
        for(int k = 0; k < 3; k++) { 
            for(int s = 0; s < 3; s++) {
                if(check[input[3*i + k][3*j + s]-1]>0){
                   flag=false; 
                   break; 
                   } else check[input[3*i + k][3*j + s]-1]+=1; 
                   } 
                   } 
                   } 
                   }
    cout<<"#"<<t<<" "<<(flag?"1":"0")<<endl;
    }
}