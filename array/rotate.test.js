const rotate = require('./rotate'); // Assume your function is in a file named 'rotate.js'

describe('rotate function', () => {
  test('rotates an array by 1 position', () => {
    const arr = [1, 2, 3, 4, 5];
    rotate(arr, 1);
    expect(arr).toEqual([2, 3, 4, 5, 1]);
  });

  test('rotates an array by 2 positions', () => {
    const arr = [1, 2, 3, 4, 5];
    rotate(arr, 2);
    expect(arr).toEqual([3, 4, 5, 1, 2]);
  });

  test('rotates an array by more than its length', () => {
    const arr = [1, 2, 3, 4, 5];
    rotate(arr, 7); // 7 rotations is effectively the same as 2 rotations for this array
    expect(arr).toEqual([3, 4, 5, 1, 2]);
  });

  test('handles empty array', () => {
    const arr = [];
    rotate(arr, 3);
    expect(arr).toEqual([]);
  });

  test('handles rotation of 0 positions', () => {
    const arr = [1, 2, 3, 4, 5];
    rotate(arr, 0);
    expect(arr).toEqual([1, 2, 3, 4, 5]);
  });
});
