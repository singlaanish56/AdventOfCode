#include<iostream>
#include<vector>
#include<unordered_map>
#include<string>
#include<sstream>
#include<algorithm>
using namespace std;




int main(){

    
    freopen("input1.txt", "r", stdin);

    vector<int> left;
    vector<int> right;
    unordered_map<int,int> store;
    string x;
    while (getline(cin, x)){

        std::istringstream iss(x);
        string l, r;

        iss >> l >>r;

        int li = stoi(l);
        int ri = stoi(r);

        left.push_back(li);
        right.push_back(ri);
        store[ri]++;

    }

    sort(left.begin(), left.end());
    sort(right.begin(), right.end());   

    int ans1=0, ans2=0;
    int sz = left.size();
    for(int i=0;i<sz;i++){

        ans1+=abs(left[i]-right[i]);
        if(store.find(left[i])!=store.end()){
            ans2+=(left[i]*store[left[i]]);
        }
    }

    cout<<"Ans 1 "<<ans1<<endl;
    cout<<"Ans 2 "<<ans2<<endl;

    return 0;
}