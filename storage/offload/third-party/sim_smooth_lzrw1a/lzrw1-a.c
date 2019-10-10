
//
// LZR cores: C model
//
// (C) 2017 IP Cores, Inc.
//
// Confidential information provided under an NDA clause in the license agreement
//
// Rev. 1.00 - initial release
//      1.01 - fixed boolean constants
//
#include <string.h>
#include "lzrw.h"

#define SMOOTH_THRESHOLD        4096
#define SMOOTH_BUF_SIZE         256

#if LZRW1A_WITH_ADLER32_ENABLE
#define INCR_GEN_ADLER32_DECL(a, b)		int a, b
#define INCR_GEN_ADLER32_START(res, a, b)	incr_gen_adler32_start(a, b)
#define INCR_GEN_ADLER32_CONT(res, a, b, m, c)	if ((res)) incr_gen_adler32_cont(a, b, m, c)
#define INCR_GEN_ADLER32_CONT1(res, a, b, m)	if ((res)) incr_gen_adler32_cont1(a, b, m)
#define INCR_GEN_ADLER32_END(res, a, b)		if ((res)) *(res) = incr_gen_adler32_end(a, b)
#else
#define INCR_GEN_ADLER32_DECL(a, b)
#define INCR_GEN_ADLER32_START(res, a, b)
#define INCR_GEN_ADLER32_CONT(res, a, b, m, c)
#define INCR_GEN_ADLER32_CONT1(res, a, b, m)
#define INCR_GEN_ADLER32_END(res, a, b)		if ((res)) *(res) = 0
#endif

static uint8_t cp_smooth_buf0[SMOOTH_BUF_SIZE];
static uint8_t cp_smooth_buf1[SMOOTH_BUF_SIZE];

/*
 * Incremental adler32
 */
static inline void incr_gen_adler32_start(int *adler32_a, int *adler32_b)
{
	*adler32_a = 1;
	*adler32_b = 0;
}

static inline void incr_gen_adler32_cont1(int *adler32_a, int *adler32_b, uint8_t msg_byte)
{
	*adler32_a += (int)(msg_byte);
	*adler32_a %= 65521;
	*adler32_b += *adler32_a;
	*adler32_b %= 65521;
}

static void incr_gen_adler32_cont(int *adler32_a, int *adler32_b, const uint8_t *msg, uint32_t bytes)
{
	for (int i = 0; i < bytes; i++)
	{
		incr_gen_adler32_cont1(adler32_a, adler32_b, msg[i]);
	}
}

static inline uint32_t incr_gen_adler32_end(int adler32_a, int adler32_b)
{
	return (uint32_t)((adler32_b << 16) | (adler32_a & 0xffff));
}

static int Hash(byte a, byte b, byte c)
{
	return (((40543 * ((a << 8) ^ (b << 4) ^ (c))) >> 4) & 0xfff);
}

static inline uint8_t cp_input_byte(const uint8_t *cp_smooth_buf, const uint8_t *data, uint32_t pos, uint32_t input_size)
{
	if (input_size <= SMOOTH_THRESHOLD)
	{
		return data[pos];
	}
	return pos < SMOOTH_BUF_SIZE ? cp_smooth_buf[pos] : data[pos - SMOOTH_BUF_SIZE];
}

static int non_smooth_match(int *hashTable, byte* data, int len, int pos, int* idx0)
{
	int size = len;
	int posBits = pos & ((1 << addrBits) - 1);
	int posMSB = pos & ~((1 << addrBits) - 1);
	*idx0 = 0;
	if (size < pos + 3) return 0; // Done

	int hash = Hash(data[pos], data[pos + 1], data[pos + 2]);
	*idx0 = hashTable[hash];
	hashTable[hash] = posBits;
	if (*idx0 >= posBits)
	{
		if (posMSB == 0)
			return 0; // Uninitialized
		*idx0 += (posMSB - (1 << addrBits));
	}
	else
		*idx0 += posMSB;
	int matched;
	for (matched = 0;
	pos + matched < size && data[pos + matched] == data[*idx0 + matched] && matched < ((1 << (16 - addrBits)) + 2);
		++matched);
	return matched;
}

static int post_smooth_match(const uint8_t *cp_smooth_buf, int *hashTable, byte* data, int len, int pos, int* idx0, uint32_t input_size)
{
	int size = len;
	int posBits = pos & ((1 << addrBits) - 1);
	int posMSB = pos & ~((1 << addrBits) - 1);
	*idx0 = 0;
	if (size < pos + 3) return 0; // Done

        int posIdx = pos - SMOOTH_BUF_SIZE;
	int hash = Hash(data[posIdx], data[posIdx + 1], data[posIdx + 2]);
	*idx0 = hashTable[hash];
	hashTable[hash] = posBits;
	if (*idx0 >= posBits)
	{
		if (posMSB == 0)
			return 0; // Uninitialized
		*idx0 += (posMSB - (1 << addrBits));
	}
	else
		*idx0 += posMSB;
	int matched;
	int _idx0 = *idx0;
	if (_idx0 >= SMOOTH_BUF_SIZE)
	{
		_idx0 -= SMOOTH_BUF_SIZE;
		for (matched = 0;
		pos + matched < size && data[posIdx + matched] == data[_idx0 + matched] && matched < ((1 << (16 - addrBits)) + 2);
			++matched);
		return matched;
	}

	for (matched = 0;
	pos + matched < size && cp_input_byte(cp_smooth_buf, data, pos + matched, input_size) == 
		cp_input_byte(cp_smooth_buf, data, _idx0 + matched, input_size) && matched < ((1 << (16 - addrBits)) + 2);
		++matched);
	return matched;
}

static int Match(const uint8_t *cp_smooth_buf, int *hashTable, byte* data, int len, int pos, int* idx0, uint32_t input_size)
{
	if (input_size <= SMOOTH_THRESHOLD)
		return non_smooth_match(hashTable, data, len, pos, idx0);

	if (pos >= SMOOTH_BUF_SIZE)
		return post_smooth_match(cp_smooth_buf, hashTable, data, len, pos, idx0, input_size);

	int size = len;
	int posBits = pos & ((1 << addrBits) - 1);
	int posMSB = pos & ~((1 << addrBits) - 1);
	*idx0 = 0;
	if (size < pos + 3) return 0; // Done
        
	int hash = Hash(cp_input_byte(cp_smooth_buf, data, pos + 0, input_size), 
                        cp_input_byte(cp_smooth_buf, data, pos + 1, input_size),
                        cp_input_byte(cp_smooth_buf, data, pos + 2, input_size));
	*idx0 = hashTable[hash];
	hashTable[hash] = posBits;
	if (*idx0 >= posBits)
	{
		if (posMSB == 0)
			return 0; // Uninitialized
		*idx0 += (posMSB - (1 << addrBits));
	}
	else
		*idx0 += posMSB;
	int matched;
	int _idx0 = *idx0;
	for (matched = 0;
	pos + matched < size && cp_input_byte(cp_smooth_buf, data, pos + matched, input_size) == 
		cp_input_byte(cp_smooth_buf, data, _idx0 + matched, input_size) && matched < ((1 << (16 - addrBits)) + 2);
		++matched);
	return matched;
}

static inline uint8_t dc_result_byte(const uint8_t *dc_smooth_buf0, const uint8_t *result, uint32_t pos, uint32_t thresh)
{
	if (thresh <= SMOOTH_THRESHOLD)
	{
		return result[pos];
	}
	return pos < SMOOTH_BUF_SIZE ? dc_smooth_buf0[pos] : result[pos - SMOOTH_BUF_SIZE];
}

int lzrw1a_compress(uint32_t action, uint8_t *hash, uint8_t *data, uint32_t size, uint8_t *result, uint32_t* p_dst_len, uint32_t thresh, uint32_t *p_adler32_chksum)
{
	int *hashTable = (int *)hash;
	uint8_t *p_data;
	uint8_t *p_result;
	uint32_t buf_size;
	uint8_t dc_smooth_buf0[SMOOTH_BUF_SIZE];
        INCR_GEN_ADLER32_DECL(adler32_a, adler32_b);

	if (cp_smooth_buf1[0] == 0)
	{
		memset(cp_smooth_buf1, 0xff, sizeof(cp_smooth_buf1));
	}
        INCR_GEN_ADLER32_START(p_adler32_chksum, &adler32_a, &adler32_b);
	if (action == COMPRESS_ACTION_COMPRESS)
	{
		int cf = FALSE; int pos = 0; int rescnt = 0;
		p_data = data;
		buf_size = size;
		if (size > SMOOTH_THRESHOLD)
		{
			p_data = data[SMOOTH_THRESHOLD - SMOOTH_BUF_SIZE] ? cp_smooth_buf0 : cp_smooth_buf1;
			buf_size += SMOOTH_BUF_SIZE;
                        INCR_GEN_ADLER32_CONT(p_adler32_chksum, &adler32_a, &adler32_b, p_data, SMOOTH_BUF_SIZE);
		}
                INCR_GEN_ADLER32_CONT(p_adler32_chksum, &adler32_a, &adler32_b, data, size);
                INCR_GEN_ADLER32_END(p_adler32_chksum, adler32_a, adler32_b);

		while (pos < buf_size)
		{
			int groupPtr = rescnt;
			int needHeader = TRUE; // Header output postponed for thresholding
			for (int item = 0; item < 16; ++item)
			{
				if (pos >= buf_size)
					break; // No more items
				int cpyIdx;
				int msize = Match(p_data, hashTable, data, buf_size, pos, &cpyIdx, size);
				int cpyOff = pos - cpyIdx;
				if (msize >= 3)
				{
					// If copy item plus header do not fit, terminate
					if (thresh != 0 && thresh < rescnt + (needHeader ? 2 : 0) + 2)
					{
						cf = TRUE;
						goto finish;
					}

					if (needHeader)
					{
						// Output header template
						result[rescnt++] = 0; // Assume all literals
						result[rescnt++] = 0;
						needHeader = FALSE;
					}

					// Output the copy item
					result[rescnt++] = (byte)(((cpyOff & 0xf00) >> (addrBits - 8)) | (msize - 3));
					result[rescnt++] = (byte)(cpyOff);
					// Update hash for all matches
					++pos;
					int dummy;
					for (int p = 1; p < msize; ++p, ++pos)
						Match(p_data, hashTable, data, buf_size, pos, &dummy, size);
					// Fix the group descriptor 
					if (item < 8)
						result[groupPtr] |= (byte)(0x01 << item);
					else
						result[groupPtr + 1] |= (byte)(0x01 << (item - 8));
				}
				else
				{
					// If copy item plus header do not fit, terminate
					if (thresh != 0 && thresh < rescnt + (needHeader ? 2 : 0) + 1)
					{
						cf = TRUE;
						goto finish;
					}

					if (needHeader)
					{
						// Output header template
						result[rescnt++] = 0; // Assume all literals
						result[rescnt++] = 0;
						needHeader = FALSE;
					}

					// Literal
					result[rescnt++] = cp_input_byte(p_data, data, pos++, size);
				}
			}
		}
	finish:
		*p_dst_len = rescnt;
		return TRUE^cf;
	}
	else
	{
		int pos = 0; int rescnt = 0, w_rescnt = 0;
		uint8_t byte_result;

		if (thresh <= SMOOTH_THRESHOLD)
		{
			while (pos < size)
			{
				uint32_t group = (uint32_t)((data[pos + 1] << 8) | (data[pos]));
				pos += 2;

				for (int item = 0; item < 16; ++item)
				{
					if (pos >= size) break;
					if ((group & (0x0001 << item)) == 0)
					{
						byte_result = result[rescnt++] = data[pos++];
						INCR_GEN_ADLER32_CONT1(p_adler32_chksum, &adler32_a, &adler32_b, byte_result);
					}
					else
					{
						int cpyOff = (data[pos + 1] | ((data[pos] << (addrBits - 8)) & 0xf00));
						if (cpyOff == 0)
							cpyOff += (1 << addrBits);
						int cpyLen = (data[pos] & (0xff >> (addrBits - 8))) + 3;
						int cpyPos = rescnt - cpyOff;
						for (int c = 0; c < cpyLen; ++c)
						{
							byte_result = result[rescnt++] = result[cpyPos + c];
							INCR_GEN_ADLER32_CONT1(p_adler32_chksum, &adler32_a, &adler32_b, byte_result);
						}
						pos += 2;
					}
				}
			}
			*p_dst_len = rescnt;
		}
		else
		{
			p_result = dc_smooth_buf0;
			while (pos < size)
			{
				uint32_t group = (uint32_t)((data[pos + 1] << 8) | (data[pos]));
				pos += 2;

				for (int item = 0; item < 16; ++item)
				{
					if (pos >= size) break;
					if ((group & (0x0001 << item)) == 0)
					{
						byte_result = p_result[w_rescnt++] = data[pos++];
	                                        INCR_GEN_ADLER32_CONT1(p_adler32_chksum, &adler32_a, &adler32_b, byte_result);
						rescnt++;
						if ((w_rescnt >= SMOOTH_BUF_SIZE) && (p_result != result))
						{
							p_result = result;
							w_rescnt = 0;
						}
					}
					else
					{
						int cpyOff = (data[pos + 1] | ((data[pos] << (addrBits - 8)) & 0xf00));
						if (cpyOff == 0)
							cpyOff += (1 << addrBits);
						int cpyLen = (data[pos] & (0xff >> (addrBits - 8))) + 3;
						int cpyPos = rescnt - cpyOff;
						if (cpyPos >= SMOOTH_BUF_SIZE) {
							int cpyIdx = cpyPos - SMOOTH_BUF_SIZE;
							for (int c = 0; c < cpyLen; ++c)
							{
								byte_result = p_result[w_rescnt++] = result[cpyIdx + c];
								INCR_GEN_ADLER32_CONT1(p_adler32_chksum, &adler32_a, &adler32_b, byte_result);
								rescnt++;
								if ((w_rescnt >= SMOOTH_BUF_SIZE) && (p_result != result))
								{
									p_result = result;
									w_rescnt = 0;
								}
							}
						}
						else
						{
							for (int c = 0; c < cpyLen; ++c)
							{
								byte_result = p_result[w_rescnt++] = dc_result_byte(dc_smooth_buf0, result, cpyPos + c, thresh);
								INCR_GEN_ADLER32_CONT1(p_adler32_chksum, &adler32_a, &adler32_b, byte_result);
								rescnt++;
								if ((w_rescnt >= SMOOTH_BUF_SIZE) && (p_result != result))
								{
									p_result = result;
									w_rescnt = 0;
								}
							}
						}
						pos += 2;
					}
				}
			}
			*p_dst_len = rescnt > SMOOTH_BUF_SIZE ? rescnt - SMOOTH_BUF_SIZE : rescnt;
		}
		INCR_GEN_ADLER32_END(p_adler32_chksum, adler32_a, adler32_b);
		return 0;
	}
}

