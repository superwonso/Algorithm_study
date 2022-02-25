#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int t,tc;
    cin>>tc;
    for (t=1; t<=tc; t++){
    int N;
    vector<vector<int>> pascal;
    cin>>N;
    cout<<"#"<<t<<"\n"<<endl;
    for (int n = 0; n < N; n++) {
        for (int k = 0; k <= n; k++) {
            if (k == 0 || k == n) {
                pascal[n][k] = 1;
            }
            else {
                pascal[n][k] = pascal[n - 1][k - 1] + pascal[n - 1][k];
            }
            cout<<pascal[n][k]<<" ";
        }
        cout<<"\n"<<endl;
    }
    }

    return 0;
}