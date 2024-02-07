
#include "gtest/gtest.h"
#include "trie.hh"

TEST (TestTrie, InsertSearch) {
  Trie* t = new(Trie);
  ASSERT_EQ(true, t->Insert("cat"));
  ASSERT_EQ(true, t->Insert("car"));
  ASSERT_EQ(true, t->Insert("call"));

  ASSERT_EQ(true, t->Search("call"));
  ASSERT_EQ(false, t->Search("caller"));
}

int main(int argc, char* argv[]) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}