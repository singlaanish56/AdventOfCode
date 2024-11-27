# AdventOfCode
all my advent code attempts in cpp

#### Default Command

```
g++ -std=c++11 -O2 -Wall TestFile.cpp -o test
```

#### Efficient Input and Output

```
ios:;sync_with_stdio(0);
cin.tie(0);

while(cin >> x){

}

string s;
getline(cin, s);

freopen("input.txt","r", stdin);
freopen("output.txt","w", stdout);
```