#include "MyBuffer.h"

MyBuffer::MyBuffer(int iSize){
    this->m_data = new std::string(iSize,'\0');
}

MyBuffer::~MyBuffer(){

}

char* MyBuffer::Data(){
    return (char*)this->m_data->data();
}

int MyBuffer::Len() const {
    return this->m_data->length();
}