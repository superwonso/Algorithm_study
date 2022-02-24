#include <iostream>

using namespace std;

int main() {
    int T;
    cin >> T;
    for(int i=1; i<=T; i++) {
        int n, k;
        cin >> n;
        cin >> k;
        cout << "#" << i << " " << n/k << " " << n%k << endl;
    }
}