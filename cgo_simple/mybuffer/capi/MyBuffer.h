#include <iostream>

using namespace std;

class MyBuffer{
    public:
        MyBuffer(int iSize);
        ~MyBuffer();

        char * Data();
        int Len() const;
    private:
        string *m_data;
};