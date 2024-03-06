package array

object rotatedArr extends App { 
  val arr = Array(1, 2, 3, 4, 5)
  val k = 2
  val rotatedArr = rotate(arr, k)
  println(rotatedArr.mkString(" "))
}

def rotate(arr: Array[Int], k: Int): Array[Int] = {
  val n = arr.length
  val actualK = k % n // In case k is greater than n, to avoid unnecessary rotations
  (arr.drop(actualK) ++ arr.take(actualK))
}
