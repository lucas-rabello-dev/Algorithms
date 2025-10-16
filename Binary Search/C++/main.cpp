#include <iostream>
#include <utility> // para std::pair
#include <vector>

using namespace std;

pair<int, int> BinarySearch(vector<int>, int target);

int main() {

    vector<int> nums;

    for (int i = 0; i < 101; i++) {
        nums.push_back(i);
    }

    auto result = BinarySearch(nums, 1);

    cout << result.first << " - " << result.second << endl;

    return 0;
}

pair<int, int> BinarySearch(vector<int> nums, int target) {
    int menor = 0,  maior = nums.size() - 1, tentativas = 0;

    while (menor <= maior) {
        tentativas += 1;
        int meio = (menor + maior) / 2;
        int chute = nums[meio];
        

        if (chute == target) {
            return {chute, tentativas};
        }
        if (chute < target) {
            menor = meio + 1;
        } else {
            maior = meio - 1;
        }
    }
    return {0, tentativas};
}