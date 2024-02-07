#include "trie.hh"

bool Trie::Insert(string word) {
  int l = word.size();
  Node* tmp = root;
  for (int i = 0; i < l; i++) {
    char c = word[i];
    if (tmp->children.find(c) == tmp->children.end()) {
      Node* n = new(Node);
      tmp->children[c] = n;
    }
    tmp = tmp->children[c];
  }
  tmp->isWord = true;
  return true;
}

bool Trie::Search(string word) {
  int l = word.size();
  Node* tmp = root;
  for (int i = 0; i < l; i++) {
    char c = word[i];
    if (tmp->children.find(c) == tmp->children.end()) {
      return false;
    }
    tmp = tmp->children[c];
  }
  return tmp->isWord;
}

bool Trie::SearchPrefix(string prefix) {
  int l = prefix.size();
  Node* tmp = root;
  for (int i = 0; i < l; i++) {
    char c = prefix[i];
    if (tmp->children.find(c) == tmp->children.end()) {
      return false;
    }
    tmp = tmp->children[c];
  }
  return true;
}