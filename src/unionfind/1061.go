package unionfind

// 1061. 按字典序排列最小的等效字符串
// 给出长度相同的两个字符串s1 和 s2 ，还有一个字符串 baseStr
// 对于每个索引 i ，如果 s1[i] == s2[i] ，那么两个字符串中的这一对应字母相同。
// 否则，它们不同，需要找到一个字符 x ，使得 s1[i] == x 和 s2[i] == x ，且 x 在 baseStr 中的索引最小。
// 例如，"a" 和 "b" 的等效字符串为 "a" ，因为 'a' == 'a' 且 'b' != 'a' 。
// 返回长度与 baseStr 相同的字符串，该字符串转译自 s1 和 s2 ，并满足上述条件。
// 1 <= s1.length, s2.length, baseStr.length <= 1000
// s1.length == s2.length

func smallestEquivalentString(s1, s2, baseStr string) string {
    fa := [26]byte{}
    for i := range fa {
        fa[i] = byte(i)
    }

    var find func(byte) byte
    find = func(x byte) byte {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y byte) {
        small, big := find(x), find(y)
        if small > big {
            small, big = big, small
        }
        // 把大的代表元指向小的代表元
        fa[big] = small
    }

    for i, x := range s1 {
        merge(byte(x)-'a', s2[i]-'a')
    }

    s := make([]byte, len(baseStr))
    for i, c := range baseStr {
        s[i] = find(byte(c)-'a') + 'a'
    }
    return string(s)
}