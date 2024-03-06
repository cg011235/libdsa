#!/usr/bin/python

"""
Rotate given array by k position
"""
def rotate(arr, k):
  n = len(arr)
  for i in range(k):
    tmp = arr[0]
    for j in range(n - 1):
      arr[j] = arr[j + 1]
    arr[n - 1] = tmp
