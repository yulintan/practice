# testList = [[1, 2, 3, 4],
#             [5, 6, 7, 8],
#             [9, 10, 11, 12],
#             [13, 14, 15, 16]]

testList = [[1, 2, 3],
            [4, 5, 6],
            [7, 8, 9]]

r = len(testList)
c = len(testList[0])


def sum(array, m, n, path):
    # when it's on the last column, it can only go down
    if n == r-1:
        for e in array[m:]:
            path.append(e[n])

        print(path)
        return

    # when it's on the last row, it can only go right
    if m == c-1:
        for e in array[m][n:]:
            path.append(e)

        print(path)
        return

    path.append(array[m][n])
    pre = path[:]
    sum(array, m, n+1, path)
    sum(array, m+1, n, pre)


sum(testList, 0, 0, [])
