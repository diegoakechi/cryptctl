package ttlv

var SampleCreateRequest = WiresharkDumpToBytes(`
0000   42 00 78 01 00 00 01 a8 42 00 77 01 00 00 00 80
0010   42 00 69 01 00 00 00 20 42 00 6a 02 00 00 00 04
0020   00 00 00 01 00 00 00 00 42 00 6b 02 00 00 00 04
0030   00 00 00 02 00 00 00 00 42 00 0c 01 00 00 00 40
0040   42 00 23 01 00 00 00 38 42 00 24 05 00 00 00 04
0050   00 00 00 01 00 00 00 00 42 00 25 01 00 00 00 20
0060   42 00 99 07 00 00 00 08 74 65 73 74 75 73 65 72
0070   42 00 a1 07 00 00 00 08 74 65 73 74 70 61 73 73
0080   42 00 0d 02 00 00 00 04 00 00 00 01 00 00 00 00
0090   42 00 0f 01 00 00 01 18 42 00 5c 05 00 00 00 04
00a0   00 00 00 01 00 00 00 00 42 00 79 01 00 00 01 00
00b0   42 00 57 05 00 00 00 04 00 00 00 02 00 00 00 00
00c0   42 00 91 01 00 00 00 e8 42 00 08 01 00 00 00 30
00d0   42 00 0a 07 00 00 00 17 43 72 79 70 74 6f 67 72
00e0   61 70 68 69 63 20 41 6c 67 6f 72 69 74 68 6d 00
00f0   42 00 0b 05 00 00 00 04 00 00 00 03 00 00 00 00
0100   42 00 08 01 00 00 00 30 42 00 0a 07 00 00 00 14
0110   43 72 79 70 74 6f 67 72 61 70 68 69 63 20 4c 65
0120   6e 67 74 68 00 00 00 00 42 00 0b 02 00 00 00 04
0130   00 00 00 80 00 00 00 00 42 00 08 01 00 00 00 30
0140   42 00 0a 07 00 00 00 18 43 72 79 70 74 6f 67 72
0150   61 70 68 69 63 20 55 73 61 67 65 20 4d 61 73 6b
0160   42 00 0b 02 00 00 00 04 00 00 00 0c 00 00 00 00
0170   42 00 08 01 00 00 00 38 42 00 0a 07 00 00 00 04
0180   4e 61 6d 65 00 00 00 00 42 00 0b 01 00 00 00 20
0190   42 00 55 07 00 00 00 07 41 45 53 4b 65 79 31 00
01a0   42 00 54 05 00 00 00 04 00 00 00 01 00 00 00 00`)

var SampleCreateResponse = WiresharkDumpToBytes(`
0000   42 00 7b 01 00 00 00 a0 42 00 7a 01 00 00 00 48
0010   42 00 69 01 00 00 00 20 42 00 6a 02 00 00 00 04
0020   00 00 00 01 00 00 00 00 42 00 6b 02 00 00 00 04
0030   00 00 00 02 00 00 00 00 42 00 92 09 00 00 00 08
0040   00 00 00 00 58 ef 52 8d 42 00 0d 02 00 00 00 04
0050   00 00 00 01 00 00 00 00 42 00 0f 01 00 00 00 48
0060   42 00 5c 05 00 00 00 04 00 00 00 01 00 00 00 00
0070   42 00 7f 05 00 00 00 04 00 00 00 00 00 00 00 00
0080   42 00 7c 01 00 00 00 20 42 00 57 05 00 00 00 04
0090   00 00 00 02 00 00 00 00 42 00 94 07 00 00 00 01
00a0   31 00 00 00 00 00 00 00`)

var SampleGetRequest = WiresharkDumpToBytes(`
0000   42 00 78 01 00 00 00 b8 42 00 77 01 00 00 00 80
0010   42 00 69 01 00 00 00 20 42 00 6a 02 00 00 00 04
0020   00 00 00 01 00 00 00 00 42 00 6b 02 00 00 00 04
0030   00 00 00 02 00 00 00 00 42 00 0c 01 00 00 00 40
0040   42 00 23 01 00 00 00 38 42 00 24 05 00 00 00 04
0050   00 00 00 01 00 00 00 00 42 00 25 01 00 00 00 20
0060   42 00 99 07 00 00 00 08 74 65 73 74 75 73 65 72
0070   42 00 a1 07 00 00 00 08 74 65 73 74 70 61 73 73
0080   42 00 0d 02 00 00 00 04 00 00 00 01 00 00 00 00
0090   42 00 0f 01 00 00 00 28 42 00 5c 05 00 00 00 04
00a0   00 00 00 0a 00 00 00 00 42 00 79 01 00 00 00 10
00b0   42 00 94 07 00 00 00 01 31 00 00 00 00 00 00 00`)

var SampleGetResponse = WiresharkDumpToBytes(`
0000   42 00 7b 01 00 00 01 00 42 00 7a 01 00 00 00 48
0010   42 00 69 01 00 00 00 20 42 00 6a 02 00 00 00 04
0020   00 00 00 01 00 00 00 00 42 00 6b 02 00 00 00 04
0030   00 00 00 02 00 00 00 00 42 00 92 09 00 00 00 08
0040   00 00 00 00 58 ef 52 8d 42 00 0d 02 00 00 00 04
0050   00 00 00 01 00 00 00 00 42 00 0f 01 00 00 00 a8
0060   42 00 5c 05 00 00 00 04 00 00 00 0a 00 00 00 00
0070   42 00 7f 05 00 00 00 04 00 00 00 00 00 00 00 00
0080   42 00 7c 01 00 00 00 80 42 00 57 05 00 00 00 04
0090   00 00 00 02 00 00 00 00 42 00 94 07 00 00 00 01
00a0   31 00 00 00 00 00 00 00 42 00 8f 01 00 00 00 58
00b0   42 00 40 01 00 00 00 50 42 00 42 05 00 00 00 04
00c0   00 00 00 01 00 00 00 00 42 00 45 01 00 00 00 18
00d0   42 00 43 08 00 00 00 10 40 4a 70 1c 3c ae ea af
00e0   04 74 25 69 cf 64 b2 7a 42 00 28 05 00 00 00 04
00f0   00 00 00 03 00 00 00 00 42 00 2a 02 00 00 00 04
0100   00 00 00 80 00 00 00 00`)

var SampleDestroyRequest = WiresharkDumpToBytes(`
0000   42 00 78 01 00 00 00 b8 42 00 77 01 00 00 00 80
0010   42 00 69 01 00 00 00 20 42 00 6a 02 00 00 00 04
0020   00 00 00 01 00 00 00 00 42 00 6b 02 00 00 00 04
0030   00 00 00 02 00 00 00 00 42 00 0c 01 00 00 00 40
0040   42 00 23 01 00 00 00 38 42 00 24 05 00 00 00 04
0050   00 00 00 01 00 00 00 00 42 00 25 01 00 00 00 20
0060   42 00 99 07 00 00 00 08 74 65 73 74 75 73 65 72
0070   42 00 a1 07 00 00 00 08 74 65 73 74 70 61 73 73
0080   42 00 0d 02 00 00 00 04 00 00 00 01 00 00 00 00
0090   42 00 0f 01 00 00 00 28 42 00 5c 05 00 00 00 04
00a0   00 00 00 14 00 00 00 00 42 00 79 01 00 00 00 10
00b0   42 00 94 07 00 00 00 01 31 00 00 00 00 00 00 00`)

var SampleDestroyResponse = WiresharkDumpToBytes(`
0000   42 00 7b 01 00 00 00 90 42 00 7a 01 00 00 00 48
0010   42 00 69 01 00 00 00 20 42 00 6a 02 00 00 00 04
0020   00 00 00 01 00 00 00 00 42 00 6b 02 00 00 00 04
0030   00 00 00 02 00 00 00 00 42 00 92 09 00 00 00 08
0040   00 00 00 00 58 ef 52 8d 42 00 0d 02 00 00 00 04
0050   00 00 00 01 00 00 00 00 42 00 0f 01 00 00 00 38
0060   42 00 5c 05 00 00 00 04 00 00 00 14 00 00 00 00
0070   42 00 7f 05 00 00 00 04 00 00 00 00 00 00 00 00
0080   42 00 7c 01 00 00 00 10 42 00 94 07 00 00 00 01
0090   31 00 00 00 00 00 00 00`)