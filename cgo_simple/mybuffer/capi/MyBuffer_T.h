typedef void* MyBufferHandle;

MyBufferHandle NewMyBuffer(int iSize);

void DeleteMyBuffer(MyBufferHandle p);

char* MyBufferData(MyBufferHandle p);

int MyBufferLen(MyBufferHandle p);

