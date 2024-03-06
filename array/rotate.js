function rotate(arr, k) {
  let n = arr.length;
  for (let i = 0; i < k; i++) {
    let tmp = arr[0];
    for (let j = 0; j < n - 1; j++) {
      arr[j] = arr[j + 1];
    }
    arr[n - 1] = tmp;
  }
}
