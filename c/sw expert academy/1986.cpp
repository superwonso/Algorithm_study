#include <iostream>

using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int t,tc;
    cin>>tc;
    for (t=1;t<=tc;t++){
        int n;
        cin>>n;
        int res[2]={0,0};
        for(int i=1;i<=n; i++) {
            if(i%2==0) {
                res[0]+=i;
            } else {
                res[1]+=i;
            }
        }
        cout<<"#"<<t<<" "<<res[1]-res[0]<<endl;
    }
}