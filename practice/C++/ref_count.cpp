//
// Created by caipengbo on 2021/2/4.
//
#include <string>
using namespace std;

class HasPtr {
public:
    HasPtr(const string &s = string()): ps(new string(s)), i(0), use(new size_t(1)) {}
    HasPtr(const HasPtr &p): ps(p.ps), i(p.i), use(p.use) {
        ++(*use);
    }
    HasPtr& operator=(const HasPtr&);
    ~HasPtr();
private:
    string *ps;  // 底层string （共享）
    int i;  // id
    size_t *use;  // 引用计数器 （共享）
};

HasPtr& HasPtr::operator=(const HasPtr &rhs) {
    // 处理自赋值
    if (&rhs == this) {
        return *this;
    }
    ++*rhs.use;  // 递增右侧引用计数
    if (--*use == 0) {  // 递减左侧引用计数
        delete ps;
        delete use;
    }
    ps = rhs.ps;
    i = rhs.i;
    use = rhs.use;
    return *this;
}

HasPtr::~HasPtr() {
    if (--(*use) == 0) {
        delete ps;
        delete use;
    }
}

int main() {
    HasPtr p1("Content");
    HasPtr p2(p1);
    HasPtr p3(p1);
}