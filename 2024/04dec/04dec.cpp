#include<iostream>
#include<string>
#include<sstream>
#include<vector>
using namespace std;

void part2(vector<string>& store, int m , int n){
    int ans=0;

    for(int i=0;i<m;i++){
        for(int j=0;j<n;j++){

                if(store[i][j]=='A'){
				int mcount=0;
				int scount=0;

				if(i-1>=0 && j-1>=0 && i+1<m && j+1<n){
					if(store[i-1][j-1]=='M' && store[i+1][j+1]!='M'){
						mcount++;
					}else if(store[i-1][j-1]=='S'  && store[i+1][j+1]!='S'){
						scount++;
					}
					if(store[i-1][j+1]=='M'  && store[i+1][j-1]!='M'){
						mcount++;
					}else if(store[i-1][j+1]=='S'  && store[i+1][j-1]!='S'){
						scount++;
					}
					if(store[i+1][j-1]=='M'  && store[i-1][j+1]!='M'){
						mcount++;
					}else if(store[i+1][j-1]=='S'  && store[i-1][j+1]!='S'){
						scount++;
					}
					if(store[i+1][j+1]=='M'  && store[i-1][j-1]!='M'){
						mcount++;
					}else if(store[i+1][j+1]=='S'  && store[i-1][j-1]!='S'){
						scount++;
					}
				
					if (mcount==2 && scount==2){
						ans++;
					}
				}
			}
        }


    }

cout<<ans<<"\n";


}



void part1(vector<string>& store, int m , int n){
    int ans=0;

    for(int i=0;i<m;i++){
        for(int j=0;j<n;j++){

            if(store[i][j]=='X'){
				//up
				if(i-3>=0 && (store[i-1][j]=='M' && store[i-2][j]=='A' && store[i-3][j]=='S')){
					ans++;
				}
				//down
				if(i+3<m && (store[i+1][j]=='M' && store[i+2][j]=='A' && store[i+3][j]=='S')){
					ans++;
				}
				//left
				if(j-3>=0 && (store[i][j-1]=='M' && store[i][j-2]=='A' && store[i][j-3]=='S')){
					ans++;
				}
				//right
				if(j+3<n && (store[i][j+1]=='M' && store[i][j+2]=='A' && store[i][j+3]=='S')){
					ans++;
				}
				//up left diagnal
				if(i-3>=0 && j-3>=0 && (store[i-1][j-1]=='M' && store[i-2][j-2]=='A' && store[i-3][j-3]=='S')){
					ans++;
				}				
				//up right diagnal
				if(i-3>=0 && j+3<n && (store[i-1][j+1]=='M' && store[i-2][j+2]=='A' && store[i-3][j+3]=='S')){
					ans++;
				}	
				//bottom left diagnal
				if(i+3<m && j-3>=0 && (store[i+1][j-1]=='M' && store[i+2][j-2]=='A' && store[i+3][j-3]=='S')){
					ans++;
				}	
				//bottom  right diagnal
				if(i+3<m && j+3<n && (store[i+1][j+1]=='M' && store[i+2][j+2]=='A' && store[i+3][j+3]=='S')){
					ans++;
				}	
			}
        }
    }

    cout<<ans<<"\n";
}

int main(){

    freopen("input1.txt", "r", stdin);

    string x;

    vector<string> store;
    while(getline(cin, x)){

        store.push_back(x);
    }

    int m = store.size();
    int n = store[0].size();

    part1(store,m,n);
    part2(store,m,n);
    return 0;
}