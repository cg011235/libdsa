#pragma once

#include <map>
#include <string>

using namespace std;

typedef struct node {
  bool isWord;
  map<char, struct node*> children;
  node(): isWord(true) {};
} Node;

class Trie {
  public:
  bool Insert(string);
  bool Search(string);
  bool SearchPrefix(string);
  Trie() {this->root = new(Node);}
  private:
   Node* root;
};