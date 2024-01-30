
#include "bst.hh"

bst::bst(int arr[]) {
  int len = sizeof(arr)/sizeof(arr[0]);
  this->root = createBst(arr, 0, len - 1);
}

bst::~bst() {
  this->root = nullptr;
}

node* bst::createBst(int arr[], int start, int end) {
  if (start <= end) {
    int mid = (start + end) / 2;
    root = new(node);
    root->key = arr[mid];
    root->left = createBst(arr, start, mid - 1);
    root->right = createBst(arr, mid + 1, end);
    return root;
  } else {
    return nullptr;
  }
}

node* bst::addNode(node* root, int key) {
  if (root == nullptr) {
    node* tmp = new(node);
    tmp->key = key;
    tmp->left = nullptr;
    tmp->right = nullptr;
    return tmp;
  }
  if (key < root->key) {
    return addNode(root->left, key);
  } else {
    return addNode(root->right, key);
  }
}

void bst::add(int key) {
  this->root = addNode(this->root, key);
}

node* bst::maxValueNode(node* n) {
  node* current = n;
  while (current && current->right != nullptr) {
    current = current->right;
  }
  return current;
}

node* bst::minValueNode(node* n) {
  node* current = n;
  while (current && current->left != nullptr) {
    current = current->left;
  }
  return current;
}

/*
 * possiblities:
 * 1. Key does not exist
 * 2. Node is leaf node (no children)
 * 3. Node has one child (eighter left or right)
 * 4. Node has both childrens.
*/
node* bst::deleteNode(node* root, int key) {
  if (root == nullptr) {
    return root;
  }

  if (key < root->key) {
    root->left = deleteNode(root->left, key);
  } else if (key > root->key) {
    root->right = deleteNode(root->right, key);
  } else {
    if (root->left == nullptr) {
      node* temp = root->right;
      delete root;
      return temp;
    } else if (root->right == nullptr) {
      node* temp = root->left;
      delete root;
      return temp;
    }
    // Node with two children so
    // get the inorder successor
    // (smallest in the right subtree)
    node* temp = minValueNode(root->right);
    root->key = temp->key;
    root->right = deleteNode(root->right, temp->key);
  }
  return root;
}

void bst::remove(int key) {
  this->root = deleteNode(this->root, key);
}

void bst::removeSubTree(node* root) {
  if(root == nullptr) {
    return;
  }
  removeSubTree(root->left);
  removeSubTree(root->right);
  delete(root);
}

node* bst::findNode(node* root, int key) {
  if (root == nullptr) {
    return nullptr;
  }
  if (root->key == key) {
    return root;
  }
  if (key < root->key) {
    return findNode(root->left, key);
  } else {
    return findNode(root->right, key);
  }
}

node* bst::find(int key) {
  return findNode(this->root, key);
}

node* bst::trimTreeUtil(node* root, int min, int max) {
  if (root == nullptr) return nullptr;

  root->left = trimTreeUtil(root->left, min, max);
  root->right = trimTreeUtil(root->right, min, max);

  if (root->key < min) {
    node* rightSub = root->right;
    delete root;
    return rightSub;
  } else if (root->key > max) {
    node* leftSub = root->left;
    delete root;
    return leftSub;
  }
  return root;
}

void bst::trimTree(int min, int max) {
  this->root = trimTreeUtil(this->root, min, max);
}