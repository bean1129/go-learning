#include "MyBuffer.h"

extern "C" {
    #include "MyBuffer_T.h"
}

class  MyBuffer_T:public MyBuffer
{
public:
    MyBuffer_T(int iSize):MyBuffer(iSize){}
    ~MyBuffer_T(){}
};


MyBufferHandle NewMyBuffer(int iSize){
    return new MyBuffer_T(iSize);
}

void DeleteMyBuffer(MyBufferHandle p){
    delete (MyBuffer_T*)p;
}

char * MyBufferData(MyBufferHandle p){
    return ((MyBuffer_T*)p)->Data();
}

int MyBufferLen(MyBufferHandle p){
    return  ((MyBuffer_T*)p)->Len();
}


