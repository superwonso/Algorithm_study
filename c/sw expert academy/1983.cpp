#include <iostream>
#include <algorithm>

using namespace std;
int main() {
	ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int T;
	cin>>T;
	int a[100];
	int N, K;
	int midterm, finalterm, homework;
	char grades[10][3] = {"A+","A0","A-","B+","B0","B-","C+","C0","C-","D0"};
	for(int i=1; i<=T; i++) {
		cin>>N>>K;
		for(int j=0; j<N; j++) {
			cin>>midterm>>finalterm>>homework;
			a[j] = midterm*35 + finalterm*45 + homework*20;
		}
		int score = a[K-1];
		int rank;
		sort(a, a+N);
		for(int j=0; j<N; j++) {
			if(a[j] == score) {
				rank = N-1-j;
				break;
			}
		}
		cout<<"#"<<i<<" "<<grades[rank*10/N]<<endl;
	}
}