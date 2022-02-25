#include<iostream>

using namespace std;

int stock[1000000];

int main(){
    int i, test_case, T, N;
    cin>> T;
    for (test_case = 1; test_case <= T; ++test_case) {
        scanf("%d", &N);
        long balance = 0;
        for (i = 0; i < N; i++){
            scanf("%d", stock + i);
        }
        int max = i - 1;
        for (i -= 2; i >= 0; --i) {
            if (stock[i] <= stock[max]){
                balance += stock[max] - stock[i];
            }
            else{
                max = i;
            }
        }
        printf("#%d %ld\n", test_case, balance);
    }
    return 0;
}