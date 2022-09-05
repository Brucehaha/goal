from collections import namedtuple
def rotateString(s: str, goal: str) -> bool:
    for i in range(len(s)):
        if goal == s[i+1:0] + s[:i+1]:
            return True

    return False


if __name__ == "__main__":

    """
    Example 1:

    Input: s = "abcde", goal = "cdeab"
    Output: true
    Example 2:
    
    Input: s = "abcde", goal = "abced"
    Output: false
"""
    inputs = [("abcde", "cdeab"), ("abcde", "abced")]
    for i in inputs:
        r = rotateString(i[0], i[1])
        assert r == True, f"{i} is not rotated string"
