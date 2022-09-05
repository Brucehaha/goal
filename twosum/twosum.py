def twoSum(nums, target: int):
    numsDict = {}
    for i, v in enumerate(nums):
        print(v)
        pvalue = target - v
        p = numsDict.get(pvalue)
        print(p, i)
        if p is not None and p != i:
            return [p, i]
        numsDict[v] = i
    return []


class TwoSum3():
    def __init__(self):
        self.map = {}

    def add(self, value):
        self.map[value] = self.map.setdefault(value, 0) + 1

    def find(self, total):
        for k, v in self.map.items():
            pair_value = total - k
            if (pair_value == v and k > 1) or (pair_value != v and self.map.get(pair_value) is not None):
                return True
        return False

    
if __name__ == "__main__":
    a = [2, 7, 11]
    print(twoSum(a, 9))

    s3 = TwoSum3()
    s3.add(1)
    s3.add(3)
    s3.add(5)
    assert s3.find(4)
    assert s3.find(6)
