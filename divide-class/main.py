import copy


def combination(n):
    if n == 1:
        return [
            [0],
            [1],
            [0, 1],
        ]

    pre = combination(n-1)

    result = copy.copy(pre)
    print("n =========================", n)
    print("pre = ", pre)
    print("result = ", result)

    for v in pre:
        duplicate = v[:]
        print("v ===", v)
        duplicate.append(n)

        result.append(duplicate)

    return result


result = combination(4)

print("result == ", result)
