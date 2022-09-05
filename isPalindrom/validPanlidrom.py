import unittest
def validPalindrom(s):
    i, j = 0, len(s) - 1
    while i < j:
        if not s[i].isalnum():
            i += 1
            continue
        if not s[j].isalnum():
            j -= 1
            continue
        if str(s[i]).lower() != str(s[j]).lower():
            return False
        i += 1
        j -= 1
    return True


if __name__ == "__main__":
    s = "A man, a plan, a canal: Panama", True
    s1 = "race a car", False
    s2 = " ", True
    test = [s, s1, s2]
    for t in test:
        assert validPalindrom(t[0]) == t[1], f"not a valid palindrom: {t}"
