#include <iostream>
#include <cstring>
#include <vector>
using namespace std;

int is_palindrome(char *str); 

int main() {
	ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int T;
	cin>>T;
	char *str;
	for (int tc=1; tc<=T; tc++) {
		cin>>str;
        printf("#%d %d\n", tc, is_palindrome(str));
    }
}

int is_palindrome(char *str) {
	int len = strlen(str);
	for(int i=0; i<len/2; i++) {
		if( str[i] != str[len-1-i] ) 
        {
        return 0;
        }
	}
	return 1;
}