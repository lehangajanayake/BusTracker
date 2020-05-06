#include "SIMComAT.h"
#include <errno.h>

void SIMComAT::begin(Stream& port)
{
	_port = &port;
	_output.begin(LOG_LEVEL_VERBOSE, this, false);
#if _SIM808_DEBUG
	_debug.begin(LOG_LEVEL_VERBOSE, &Serial, false);
#endif // _SIM808_DEBUG
}

void SIMComAT::flushInput() {
	uint16_t timeout = 0;
	while(readNext(replyBuffer, BUFFER_SIZE, &timeout));
	memset(replyBuffer, 0, BUFFER_SIZE);
}


size_t SIMComAT::readNext(char * buffer, size_t size, uint16_t * timeout, char stop)
{
	size_t i = 0;
	bool exit = false;

	do {
		while(!exit && i < size - 1 && available()) {
			char c = read();
			buffer[i] = c;
			i++;

			exit |= stop && c == stop;
		}

		if(timeout) {			
			if(*timeout) {
				delay(1);
				(*timeout)--;
			}

			if(!(*timeout)) break;
		}
	} while(!exit && i < size - 1);

	buffer[i] = '\0';

	if(i) {
		RECEIVEARROW;
		SIM808_PRINT(buffer);
	}

	return i > 0 ? i - 1 : i;
}

int8_t SIMComAT::waitResponse(uint16_t timeout, 
	ATConstStr s1, 
	ATConstStr s2,
	ATConstStr s3,
	ATConstStr s4)
{
	ATConstStr wantedTokens[4] = { s1, s2, s3, s4 };
	size_t length;

	do {
		memset(replyBuffer, 0, BUFFER_SIZE);
		length = readNext(replyBuffer, BUFFER_SIZE, &timeout, '\n');

		if(!length) continue; 					//read nothing
		if(wantedTokens[0] == NULL) return 0;	//looking for a line with any content

		for(uint8_t i = 0; i < 4; i++) {
			if(wantedTokens[i]) {
				char *p = strstr_P(replyBuffer, TO_P(wantedTokens[i]));
				if(replyBuffer == p) return i;				
			}
		}
	} while(timeout);

	return -1;
}

size_t SIMComAT::copyCurrentLine(char *dst, size_t dstSize, uint16_t shift)
{
	char *p = dst;
	char *p1;
	
	p += safeCopy(replyBuffer + shift, p, dstSize); // copy the current buffer content	
	// copy the rest of the line if any
	if(!strchr(dst, '\n')) p += readNext(p, dstSize - (p - dst), NULL, '\n');

	// terminating the string no matter what
	p1 = strchr(dst, '\n');
	p = p1 ? p1 : p;
	*p = '\0';

	return strlen(dst);
}

size_t SIMComAT::safeCopy(const char *src, char *dst, size_t dstSize)
{
	size_t len = strlen(src);
	if (dst != NULL) {
		size_t maxLen = min(len + 1, dstSize);
		strlcpy(dst, src, maxLen);
	}

	return len;
}

char* SIMComAT::find(const char* str, char divider, uint8_t index)
{
	char* p = strchr(str, ':');
	if (p == NULL) p = strchr(str, str[0]); //ditching eventual response header

	p++;
	for (uint8_t i = 0; i < index; i++)
	{
		p = strchr(p, divider);
		if (p == NULL) return NULL;
		p++;
	}

	return p;
}

bool SIMComAT::parse(const char* str, char divider, uint8_t index, uint8_t* result)
{
	uint16_t tmpResult;
	if (!parse(str, divider, index, &tmpResult)) return false;

	*result = (uint8_t)tmpResult;
	return true;
}

bool SIMComAT::parse(const char* str, char divider, uint8_t index, int8_t* result)
{
	int16_t tmpResult;
	if (!parse(str, divider, index, &tmpResult)) return false;

	*result = (int8_t)tmpResult;
	return true;
}

bool SIMComAT::parse(const char* str, char divider, uint8_t index, uint16_t* result)
{
	char* p = find(str, divider, index);
	if (p == NULL) return false;

	errno = 0;
	*result = strtoul(p, NULL, 10);

	return errno == 0;
}

#if defined(NEED_SIZE_T_OVERLOADS)
bool SIMComAT::parse(const char* str, char divider, uint8_t index, size_t* result) { 
	char* p = find(str, divider, index);
	if (p == NULL) return false;

	errno = 0;
	*result = strtoull(p, NULL, 10);
	
	return errno == 0; 
}
#endif

bool SIMComAT::parse(const char* str, char divider, uint8_t index, int16_t* result)
{	
	char* p = find(str, divider, index);
	if (p == NULL) return false;

	errno = 0;
	*result = strtol(p, NULL, 10);
	
	return errno == 0;
}

bool SIMComAT::parse(const char* str, char divider, uint8_t index, float* result)
{
	char* p = find(str, divider, index);
	if (p == NULL) return false;

	errno = 0;
	*result = strtod(p, NULL);

	return errno == 0;
}
