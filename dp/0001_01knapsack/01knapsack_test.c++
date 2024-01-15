#include <iostream>
#include <vector>
#include <gtest/gtest.h>
#include "01knapsack.h++"

TEST(Kanapsack, Basic) {
  vector<int> W = {10, 20, 30};
	vector<int> V = {60, 100, 120};
	int capacity = 50;
	int expected = 220;

  EXPECT_EQ(expected, knapsack(V, W, capacity));
}

TEST(Kanapsack, EmptyKnapsack) {
  vector<int> W = {10, 20, 30};
	vector<int> V = {60, 100, 120};
	int capacity = 0;
	int expected = 0;

  EXPECT_EQ(expected, knapsack(V, W, capacity));
}

TEST(Kanapsack, SingleItemFits) {
  vector<int> W = {4, 5, 6};
	vector<int> V = {1, 2, 3};
	int capacity = 5;
	int expected = 2;

  EXPECT_EQ(expected, knapsack(V, W, capacity));
}

TEST(Kanapsack, AllItemsFits) {
  vector<int> W = {4, 5, 6};
	vector<int> V = {1, 2, 3};
	int capacity = 15;
	int expected = 6;

  EXPECT_EQ(expected, knapsack(V, W, capacity));
}

TEST(Kanapsack, LargeCapacity) {
  vector<int> W = {19,
		10,
		24,
		6,
		29,
		20,
		32,
		29,
		44,
		1,
		43,
		36,
		28,
		34,
		46,
		29,
		25,
		41,
		49,
		23,
		15,
		13,
		2,
		49,
		45,
		10,
		5,
		8,
		50,
		11,
		37,
		44,
		2,
		41,
		26,
		38,
		27,
		35,
		7,
		28,
		6,
		7,
		47,
		6,
		50,
		18,
		44,
		23,
		2,
		18,
		42,
		38,
		8,
		9,
		18,
		17,
		16,
		20,
		35,
		8,
		49,
		15,
		37,
		1,
		43,
		47,
		32,
		2,
		11,
		42,
		22,
		2,
		49,
		12,
		28,
		44,
		33,
		13,
		5,
		47,
		37,
		25,
		15,
		30,
		20,
		35,
		43,
		13,
		21,
		35,
		25,
		45,
		30,
		50,
		6,
		20,
		10,
		27,
		38,
		30};
	vector<int> V = {6, 63, 38, 55, 59, 94, 54, 86, 3, 42, 58, 35, 60, 87, 16, 45, 25, 78, 50, 5, 90, 62, 22, 3, 24, 61, 46, 2, 66, 71, 90, 12, 96, 41, 40, 80, 26, 94, 1, 30, 22, 56, 40, 60, 90, 55, 77, 96, 1, 59, 81, 47, 63, 53, 13, 97, 81, 96, 65, 84, 62, 67, 77, 27, 82, 63, 75, 98, 31, 47, 4, 79, 66, 29, 49, 46, 17, 18, 31, 75, 90, 28, 42, 11, 26, 72, 36, 45, 61, 80, 33, 76, 64, 74, 96, 74, 3, 76, 52, 85};
	int capacity = 5000;
	int expected = 5319;

  EXPECT_EQ(expected, knapsack(V, W, capacity));
}

int main(int argc, char** argv) {
  testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}