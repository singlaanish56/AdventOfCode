#include<iostream>
#include<string>
#include<sstream>
#include<vector>
#include<unordered_map>
#include<unordered_set>
#include<algorithm>
#include <fstream>
using namespace std;


bool part1(vector<int>& update,unordered_map<int, vector<int>>& constraints) {
    unordered_set<int> seen; 

    for (int num : update) {
     
        if (constraints.count(num)) {
            const auto& dependencies = constraints.at(num);
          
            for (int dep : dependencies) {
                if (seen.find(dep) != seen.end()) {
                    return false;
                }
            }
        }
       
        seen.insert(num);
    }

    return true; 
}

int part2(vector<int>& update,unordered_map<int, vector<int>>& constraints){

    vector<int> fixedList = update;
    bool changesMade;

    do {
        changesMade = false;
        
        int sz = fixedList.size();
        for (int i = 0; i < sz; ++i) {
            int num = fixedList[i];

            if (constraints.count(num)) {
                const auto& dependencies = constraints.at(num);

            
                for (int dep : dependencies) {
                    auto it_num = find(fixedList.begin(), fixedList.end(), num);
                    auto it_dep = find(fixedList.begin(), fixedList.end(), dep);

                 
                    if (it_dep != fixedList.end() && it_num != fixedList.end() && it_dep < it_num) {
                        
                        fixedList.erase(it_num); 
                        fixedList.insert(it_dep, num); 
                        changesMade = true; 
                        break; 
                    }
                }
            }

            if (changesMade) {
                break;
            }
        }

    } while (changesMade); // Keep checking until no changes are made

    int sz  = fixedList.size();
    int middleindex = sz/2;
    return fixedList[middleindex];
}

void checkValidOrNot(unordered_map<int, vector<int>>& constraints,vector<vector<int>>& updates){

    int ans=0;
    int ans2=0;
    for(auto& update : updates){
        if (part1(update, constraints)){
    
            int sz  = update.size();
            int middleindex = sz/2;
            ans+=update[middleindex];
        }else{

            //update get the write order
            ans2+=part2(update, constraints);
        }
    }

    cout<<ans<<"\n";
    cout<<ans2<<"\n";
}
int main(){

    ifstream file1("input1-1.txt");
    if (!file1.is_open()) {
        cerr << "Error opening input2-1.txt" << endl;
        return 1;
    }

    string x;

    unordered_map<int, vector<int>> adj;
    while(getline(file1, x)){

        istringstream iss(x);

        string k;
        
        
        vector<int> level;
        while(getline(iss, k, '|')){

            level.push_back(stoi(k));
        }

        adj[level[0]].push_back(level[1]);

    }

    // for(auto& m : adj){
    //     cout<<m.first<<"\n";
    //     cout<<":[";
    //     for (int v : m.second){
    //         cout<<v<<" ";
    //     }

    //     cout<<"]"<<"\n";
    // }

    file1.close();

    ifstream file2("input1-2.txt");
    if (!file2.is_open()) {
        cerr << "Error opening input2-2.txt" << endl;
        return 1;
    }


    string x1;

    vector<vector<int>> updates;

    while(getline(file2, x1)){

        istringstream iss(x1);

        string k;

        vector<int> update;
        while(getline(iss, k , ',')){
            
            update.push_back(stoi(k));
        }
        updates.push_back(update);

        // cout<<"[";
        // for(int v : update){
        //     cout<<v<<" ";
        // }

        // cout<<"]"<<"\n";
    }
    
    file2.close();
    checkValidOrNot(adj , updates);

    return 0;
}