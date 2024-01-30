#pragma once

struct node {
  int key;
  struct node *left;
  struct node *right;
};

class bst {
private:
  node *root;
  node* bst::createBst(int arr[], int start, int end);
  node* bst::addNode(node* root, int key);
  node* bst::deleteNode(node* tmp, int key);
  node* bst::minValueNode(node* node);
  node* bst::maxValueNode(node* node);
  void bst::removeSubTree(node* root);
  node* bst::findNode(node* root, int key);
  node* bst::trimTreeUtil(node* root, int min, int max);
public:
  bst(int arr[]);
  void add(int key);
  void remove(int key);
  node* bst::find(int key);
  void bst::trimTree(int min, int max);
  ~bst(); 
};
