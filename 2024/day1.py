def preprocessing():

    with open("input1.txt", "r") as reader:
        l = reader.readlines()
        a = []
        b = []
        safe_l = []
        for i in range(len(l)):
            safe_l.append(int(l[i].split(" ")[0].strip()))
            safe_l.append(int(l[i].split(" ")[3].strip()))

        for i in range(len(safe_l)):
            a.append(safe_l[i]) if i % 2 else b.append(safe_l[i])
    return a, b


def partition(arr, lo, hi):
    pivot = arr[hi]
    i = lo - 1
    j = lo
    while j < hi:
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
        j += 1
    arr[i+1], arr[hi] = arr[hi], arr[i+1]
    return i + 1


def quick_sort(arr, lo, hi):
    if lo < hi:
        pivot = partition(arr, lo, hi)
        quick_sort(arr, lo, pivot - 1)
        quick_sort(arr, pivot + 1, hi)


def subtraction(a, b):
    # obviously it will be equals to b as given
    c = []
    for i in range(len(a)):
        c_el = abs(a[i] - b[i])
        c.append(c_el)
    return c


def second_problem(a, b):
    d = []
    for i in range(len(a)):
        x = b.count(a[i])
        mult = x * a[i]
        d.append(mult)
    return d


def main():
    a, b = preprocessing()
    quick_sort(a, 0, len(a) - 1)
    quick_sort(b, 0, len(b) - 1)
    c = subtraction(a, b)
    print("first answer", sum(c))
    d = second_problem(a, b)
    print("second answer", sum(d))


if __name__ == "__main__":
    main()
