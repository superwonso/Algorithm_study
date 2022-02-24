#include <iostream>

using namespace std;

int main() {
    int n;
    cin>>n;
    int res=1;
    for(int i=0; i<n; i++){
        printf("%d ",res);
        res*=2;
    }
        printf("%d",res);
}