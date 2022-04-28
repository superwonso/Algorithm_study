#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int rcm[110] = {0, }; // 추천수 배열
int dur[110] = {0, }; // 게시 기간 배열
	
bool comp(const int o1, const int o2){
	if(rcm[o1] == rcm[o2]){
		if(dur[o1] < dur[o2]){
			return true;
		}
		else{
			return false;
		}
	}
	else if(rcm[o1] > rcm[o2]){
		return true;
	}
	else{
		return false;
	}
}

int main() {
	int n, c;
	vector<int> photo; // 사진틀
	
	scanf("%d %d", &n, &c);
	
	for(int i = 0; i < c; i++){
		int x;
		scanf("%d ", &x); // 후보 입력
		
		rcm[x]++;
		
		for(int j = 0; j < photo.size(); j++){
			dur[photo[j]]++; // 후보 입력마다 사진틀에 게시된 후보 기간 증가
		}
		
		bool exist = false;
		
		for(int j = 0; j < photo.size(); j++){
			if(photo[j] == x){
				exist = true;
				break;
			}
		}
		
		if(!exist){
			if(photo.size() < n){
				photo.push_back(x);
			}else{
				sort(photo.begin(), photo.end(), comp);
                		// 재정의한 정렬 기준에 따라 사진틀 정렬
				rcm[photo.back()] = 0;
				dur[photo.back()] = 0;
				photo.back() = x;
			}
		}
		
	}
	
	sort(photo.begin(), photo.end()); // 후보 오름차순 정렬
	
	for(int i = 0; i < photo.size(); i++){
		printf("%d ", photo[i]);
	}
	
}

// 출처 : https://velog.io/@hyez_dev/백준-1713-후보-추천하기-C