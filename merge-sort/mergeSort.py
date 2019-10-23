def mergeSort(arr):
    if len(arr) > 1:
        midIndex = len(arr) / 2
        left = arr[:midIndex]
        right = arr[midIndex:]

        mergeSort(left)
        mergeSort(right)

        i = j = 0
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                arr.append(left[i])
                i += 1
            else:
                arr.append(right[j])
                j += 1

        while i < len(left):
            arr.append(left[i])
            i += 1

        while j < len(right):
            arr.append(right[j])
            j += 1


testList = [38, 27, 43, 3, 9, 82, 10]

mergeSort(testList)

print(testList)
