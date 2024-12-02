#include<iostream>
#include<string>
#include<sstream>
#include<vector>
using namespace std;

bool isIncreasing(vector<int>& level){

    int sz = level.size();
    for(int i=1;i<sz;i++){
        if((level[i]<=level[i-1]) || (level[i]-level[i-1]>3)){
            return false;
        }
    }

    return true;
}

bool isDecreasing(vector<int>& level){

    int sz = level.size();
    for(int i=1;i<sz;i++){
        if((level[i]>=level[i-1]) || (level[i-1]-level[i]>3)){
            return false;
        }
    }

    return true;
}

int calculateTheSafeLevels(vector<vector<int>>& store, bool dampener){

    int ans=0;
    for(auto& level : store){
        
        if (isDecreasing(level) || isIncreasing(level)){
            ans++;
            continue;
        }

        if(dampener){
            int sz = level.size();
            for(int i=0;i<sz;i++){
                vector<int> temp;
                temp.reserve(level.size()-1);
                temp.insert(temp.end(), level.begin(), level.begin() + i);
                temp.insert(temp.end(), level.begin() + i + 1, level.end());
                if (isDecreasing(temp) || isIncreasing(temp)){
                    ans++;
                    break;
                }
            }
        }

    }

    return ans;
}

int main(){

    freopen("input1.txt", "r", stdin);

    string x;

    vector<vector<int>> store;
    while(getline(cin, x)){

        istringstream iss(x);

        string k;
        vector<int> level;
        while(iss >> k){

            level.push_back(stoi(k));
        }


        store.push_back(level);
    }

    cout<<"the first start answer is "<<calculateTheSafeLevels(store, false)<<endl;
    cout<<"the second start answer is "<<calculateTheSafeLevels(store, true)<<endl;
    return 0;
}