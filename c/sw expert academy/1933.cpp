#include <iostream>
#include <vector>
using namespace std;

int main() {
    int n;
    ios_base :: sync_with_stdio(false);
    cin>>n;
    for(int i=1; i<=n; i++){
        if(n%i==0){
            cout<<i<<" ";
        }
    }
    return 0;
}