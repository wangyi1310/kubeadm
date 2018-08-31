/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cpumanager

import (
	"k8s.io/kubeadm/pkg/kubelet/cm/cpumanager/topology"
)

var (
	topoSingleSocketHT = &topology.CPUTopology{
		NumCPUs:    8,
		NumSockets: 1,
		NumCores:   4,
		CPUDetails: map[int]topology.CPUInfo{
			0: {CoreID: 0, SocketID: 0},
			1: {CoreID: 1, SocketID: 0},
			2: {CoreID: 2, SocketID: 0},
			3: {CoreID: 3, SocketID: 0},
			4: {CoreID: 0, SocketID: 0},
			5: {CoreID: 1, SocketID: 0},
			6: {CoreID: 2, SocketID: 0},
			7: {CoreID: 3, SocketID: 0},
		},
	}

	topoDualSocketHT = &topology.CPUTopology{
		NumCPUs:    12,
		NumSockets: 2,
		NumCores:   6,
		CPUDetails: map[int]topology.CPUInfo{
			0:  {CoreID: 0, SocketID: 0},
			1:  {CoreID: 1, SocketID: 1},
			2:  {CoreID: 2, SocketID: 0},
			3:  {CoreID: 3, SocketID: 1},
			4:  {CoreID: 4, SocketID: 0},
			5:  {CoreID: 5, SocketID: 1},
			6:  {CoreID: 0, SocketID: 0},
			7:  {CoreID: 1, SocketID: 1},
			8:  {CoreID: 2, SocketID: 0},
			9:  {CoreID: 3, SocketID: 1},
			10: {CoreID: 4, SocketID: 0},
			11: {CoreID: 5, SocketID: 1},
		},
	}

	topoDualSocketNoHT = &topology.CPUTopology{
		NumCPUs:    8,
		NumSockets: 2,
		NumCores:   8,
		CPUDetails: map[int]topology.CPUInfo{
			0: {CoreID: 0, SocketID: 0},
			1: {CoreID: 1, SocketID: 0},
			2: {CoreID: 2, SocketID: 0},
			3: {CoreID: 3, SocketID: 0},
			4: {CoreID: 4, SocketID: 1},
			5: {CoreID: 5, SocketID: 1},
			6: {CoreID: 6, SocketID: 1},
			7: {CoreID: 7, SocketID: 1},
		},
	}

	/*
		Topology from https://www.open-mpi.org/projects/hwloc/lstopo/images/KNL.SNC4.H50.v1.11.png.
		Socket0:
		0-2,9-10,13-14,21-22,25-26,33-34,38-39,46-47,50,57-58,71-72,79-80,87-88,95-96,103-104,109-110,117-118,
		131-132,139-140,147-148,155-156,163-164,169-170,177-178,191-192,199-200,207-208,215-216,223-224,229-230,
		237-238,251-252,259-260,267-268,275-276,283-284
		Socket1:
		3-4,11-12,15-16,23-24,27-28,35-36,40-41,48-49,51-52,59-60,65-66,73-74,81-82,89-90,97-98,111-112,119-120,125-126,
		133-134,141-142,149-150,157-158,171-172,179-180,185-186,193-194,201-202,209-210,217-218,231-232,239-240,245-246,
		253-254,261-262,269-270,277-278
		Socket2:
		5-6,17-18,29-30,42-43,53-54,61-62,67-68,75-76,83-84,91-92,99-100,105-106,113-114,121-122,127-128,135-136,
		143-144,151-152,159-160,165-166,173-174,181-182,187-188,195-196,203-204,211-212,219-220,225-226,233-234,241-242,
		247-248,255-256,263-264,271-272,279-280,285-286
		Socket3:
		7-8,19-20,31-32,37,44-45,55-56,63-64,69-70,77-78,85-86,93-94,101-102,107-108,115-116,123-124,129-130,137-138,
		145-146,153-154,161-162,167-168,175-176,183-184,189-190,197-198,205-206,213-214,221-222,227-228,235-236,243-244,
		249-250,257-258,265-266,273-274,281-282,287
	*/
	topoQuadSocketFourWayHT = &topology.CPUTopology{
		NumCPUs:    288,
		NumSockets: 4,
		NumCores:   72,
		CPUDetails: map[int]topology.CPUInfo{
			0:   {CoreID: 0, SocketID: 0},
			169: {CoreID: 0, SocketID: 0},
			109: {CoreID: 0, SocketID: 0},
			229: {CoreID: 0, SocketID: 0},
			50:  {CoreID: 1, SocketID: 0},
			170: {CoreID: 1, SocketID: 0},
			110: {CoreID: 1, SocketID: 0},
			230: {CoreID: 1, SocketID: 0},
			1:   {CoreID: 64, SocketID: 0},
			25:  {CoreID: 64, SocketID: 0},
			13:  {CoreID: 64, SocketID: 0},
			38:  {CoreID: 64, SocketID: 0},
			2:   {CoreID: 65, SocketID: 0},
			26:  {CoreID: 65, SocketID: 0},
			14:  {CoreID: 65, SocketID: 0},
			39:  {CoreID: 65, SocketID: 0},
			9:   {CoreID: 72, SocketID: 0},
			33:  {CoreID: 72, SocketID: 0},
			21:  {CoreID: 72, SocketID: 0},
			46:  {CoreID: 72, SocketID: 0},
			10:  {CoreID: 73, SocketID: 0},
			34:  {CoreID: 73, SocketID: 0},
			22:  {CoreID: 73, SocketID: 0},
			47:  {CoreID: 73, SocketID: 0},
			57:  {CoreID: 8, SocketID: 0},
			177: {CoreID: 8, SocketID: 0},
			117: {CoreID: 8, SocketID: 0},
			237: {CoreID: 8, SocketID: 0},
			58:  {CoreID: 9, SocketID: 0},
			178: {CoreID: 9, SocketID: 0},
			118: {CoreID: 9, SocketID: 0},
			238: {CoreID: 9, SocketID: 0},
			71:  {CoreID: 24, SocketID: 0},
			191: {CoreID: 24, SocketID: 0},
			131: {CoreID: 24, SocketID: 0},
			251: {CoreID: 24, SocketID: 0},
			72:  {CoreID: 25, SocketID: 0},
			192: {CoreID: 25, SocketID: 0},
			132: {CoreID: 25, SocketID: 0},
			252: {CoreID: 25, SocketID: 0},
			79:  {CoreID: 32, SocketID: 0},
			199: {CoreID: 32, SocketID: 0},
			139: {CoreID: 32, SocketID: 0},
			259: {CoreID: 32, SocketID: 0},
			80:  {CoreID: 33, SocketID: 0},
			200: {CoreID: 33, SocketID: 0},
			140: {CoreID: 33, SocketID: 0},
			260: {CoreID: 33, SocketID: 0},
			87:  {CoreID: 40, SocketID: 0},
			207: {CoreID: 40, SocketID: 0},
			147: {CoreID: 40, SocketID: 0},
			267: {CoreID: 40, SocketID: 0},
			88:  {CoreID: 41, SocketID: 0},
			208: {CoreID: 41, SocketID: 0},
			148: {CoreID: 41, SocketID: 0},
			268: {CoreID: 41, SocketID: 0},
			95:  {CoreID: 48, SocketID: 0},
			215: {CoreID: 48, SocketID: 0},
			155: {CoreID: 48, SocketID: 0},
			275: {CoreID: 48, SocketID: 0},
			96:  {CoreID: 49, SocketID: 0},
			216: {CoreID: 49, SocketID: 0},
			156: {CoreID: 49, SocketID: 0},
			276: {CoreID: 49, SocketID: 0},
			103: {CoreID: 56, SocketID: 0},
			223: {CoreID: 56, SocketID: 0},
			163: {CoreID: 56, SocketID: 0},
			283: {CoreID: 56, SocketID: 0},
			104: {CoreID: 57, SocketID: 0},
			224: {CoreID: 57, SocketID: 0},
			164: {CoreID: 57, SocketID: 0},
			284: {CoreID: 57, SocketID: 0},
			3:   {CoreID: 66, SocketID: 1},
			27:  {CoreID: 66, SocketID: 1},
			15:  {CoreID: 66, SocketID: 1},
			40:  {CoreID: 66, SocketID: 1},
			4:   {CoreID: 67, SocketID: 1},
			28:  {CoreID: 67, SocketID: 1},
			16:  {CoreID: 67, SocketID: 1},
			41:  {CoreID: 67, SocketID: 1},
			11:  {CoreID: 74, SocketID: 1},
			35:  {CoreID: 74, SocketID: 1},
			23:  {CoreID: 74, SocketID: 1},
			48:  {CoreID: 74, SocketID: 1},
			12:  {CoreID: 75, SocketID: 1},
			36:  {CoreID: 75, SocketID: 1},
			24:  {CoreID: 75, SocketID: 1},
			49:  {CoreID: 75, SocketID: 1},
			51:  {CoreID: 2, SocketID: 1},
			171: {CoreID: 2, SocketID: 1},
			111: {CoreID: 2, SocketID: 1},
			231: {CoreID: 2, SocketID: 1},
			52:  {CoreID: 3, SocketID: 1},
			172: {CoreID: 3, SocketID: 1},
			112: {CoreID: 3, SocketID: 1},
			232: {CoreID: 3, SocketID: 1},
			59:  {CoreID: 10, SocketID: 1},
			179: {CoreID: 10, SocketID: 1},
			119: {CoreID: 10, SocketID: 1},
			239: {CoreID: 10, SocketID: 1},
			60:  {CoreID: 11, SocketID: 1},
			180: {CoreID: 11, SocketID: 1},
			120: {CoreID: 11, SocketID: 1},
			240: {CoreID: 11, SocketID: 1},
			65:  {CoreID: 18, SocketID: 1},
			185: {CoreID: 18, SocketID: 1},
			125: {CoreID: 18, SocketID: 1},
			245: {CoreID: 18, SocketID: 1},
			66:  {CoreID: 19, SocketID: 1},
			186: {CoreID: 19, SocketID: 1},
			126: {CoreID: 19, SocketID: 1},
			246: {CoreID: 19, SocketID: 1},
			73:  {CoreID: 26, SocketID: 1},
			193: {CoreID: 26, SocketID: 1},
			133: {CoreID: 26, SocketID: 1},
			253: {CoreID: 26, SocketID: 1},
			74:  {CoreID: 27, SocketID: 1},
			194: {CoreID: 27, SocketID: 1},
			134: {CoreID: 27, SocketID: 1},
			254: {CoreID: 27, SocketID: 1},
			81:  {CoreID: 34, SocketID: 1},
			201: {CoreID: 34, SocketID: 1},
			141: {CoreID: 34, SocketID: 1},
			261: {CoreID: 34, SocketID: 1},
			82:  {CoreID: 35, SocketID: 1},
			202: {CoreID: 35, SocketID: 1},
			142: {CoreID: 35, SocketID: 1},
			262: {CoreID: 35, SocketID: 1},
			89:  {CoreID: 42, SocketID: 1},
			209: {CoreID: 42, SocketID: 1},
			149: {CoreID: 42, SocketID: 1},
			269: {CoreID: 42, SocketID: 1},
			90:  {CoreID: 43, SocketID: 1},
			210: {CoreID: 43, SocketID: 1},
			150: {CoreID: 43, SocketID: 1},
			270: {CoreID: 43, SocketID: 1},
			97:  {CoreID: 50, SocketID: 1},
			217: {CoreID: 50, SocketID: 1},
			157: {CoreID: 50, SocketID: 1},
			277: {CoreID: 50, SocketID: 1},
			98:  {CoreID: 51, SocketID: 1},
			218: {CoreID: 51, SocketID: 1},
			158: {CoreID: 51, SocketID: 1},
			278: {CoreID: 51, SocketID: 1},
			5:   {CoreID: 68, SocketID: 2},
			29:  {CoreID: 68, SocketID: 2},
			17:  {CoreID: 68, SocketID: 2},
			42:  {CoreID: 68, SocketID: 2},
			6:   {CoreID: 69, SocketID: 2},
			30:  {CoreID: 69, SocketID: 2},
			18:  {CoreID: 69, SocketID: 2},
			43:  {CoreID: 69, SocketID: 2},
			53:  {CoreID: 4, SocketID: 2},
			173: {CoreID: 4, SocketID: 2},
			113: {CoreID: 4, SocketID: 2},
			233: {CoreID: 4, SocketID: 2},
			54:  {CoreID: 5, SocketID: 2},
			174: {CoreID: 5, SocketID: 2},
			114: {CoreID: 5, SocketID: 2},
			234: {CoreID: 5, SocketID: 2},
			61:  {CoreID: 12, SocketID: 2},
			181: {CoreID: 12, SocketID: 2},
			121: {CoreID: 12, SocketID: 2},
			241: {CoreID: 12, SocketID: 2},
			62:  {CoreID: 13, SocketID: 2},
			182: {CoreID: 13, SocketID: 2},
			122: {CoreID: 13, SocketID: 2},
			242: {CoreID: 13, SocketID: 2},
			67:  {CoreID: 20, SocketID: 2},
			187: {CoreID: 20, SocketID: 2},
			127: {CoreID: 20, SocketID: 2},
			247: {CoreID: 20, SocketID: 2},
			68:  {CoreID: 21, SocketID: 2},
			188: {CoreID: 21, SocketID: 2},
			128: {CoreID: 21, SocketID: 2},
			248: {CoreID: 21, SocketID: 2},
			75:  {CoreID: 28, SocketID: 2},
			195: {CoreID: 28, SocketID: 2},
			135: {CoreID: 28, SocketID: 2},
			255: {CoreID: 28, SocketID: 2},
			76:  {CoreID: 29, SocketID: 2},
			196: {CoreID: 29, SocketID: 2},
			136: {CoreID: 29, SocketID: 2},
			256: {CoreID: 29, SocketID: 2},
			83:  {CoreID: 36, SocketID: 2},
			203: {CoreID: 36, SocketID: 2},
			143: {CoreID: 36, SocketID: 2},
			263: {CoreID: 36, SocketID: 2},
			84:  {CoreID: 37, SocketID: 2},
			204: {CoreID: 37, SocketID: 2},
			144: {CoreID: 37, SocketID: 2},
			264: {CoreID: 37, SocketID: 2},
			91:  {CoreID: 44, SocketID: 2},
			211: {CoreID: 44, SocketID: 2},
			151: {CoreID: 44, SocketID: 2},
			271: {CoreID: 44, SocketID: 2},
			92:  {CoreID: 45, SocketID: 2},
			212: {CoreID: 45, SocketID: 2},
			152: {CoreID: 45, SocketID: 2},
			272: {CoreID: 45, SocketID: 2},
			99:  {CoreID: 52, SocketID: 2},
			219: {CoreID: 52, SocketID: 2},
			159: {CoreID: 52, SocketID: 2},
			279: {CoreID: 52, SocketID: 2},
			100: {CoreID: 53, SocketID: 2},
			220: {CoreID: 53, SocketID: 2},
			160: {CoreID: 53, SocketID: 2},
			280: {CoreID: 53, SocketID: 2},
			105: {CoreID: 60, SocketID: 2},
			225: {CoreID: 60, SocketID: 2},
			165: {CoreID: 60, SocketID: 2},
			285: {CoreID: 60, SocketID: 2},
			106: {CoreID: 61, SocketID: 2},
			226: {CoreID: 61, SocketID: 2},
			166: {CoreID: 61, SocketID: 2},
			286: {CoreID: 61, SocketID: 2},
			7:   {CoreID: 70, SocketID: 3},
			31:  {CoreID: 70, SocketID: 3},
			19:  {CoreID: 70, SocketID: 3},
			44:  {CoreID: 70, SocketID: 3},
			8:   {CoreID: 71, SocketID: 3},
			32:  {CoreID: 71, SocketID: 3},
			20:  {CoreID: 71, SocketID: 3},
			45:  {CoreID: 71, SocketID: 3},
			37:  {CoreID: 63, SocketID: 3},
			168: {CoreID: 63, SocketID: 3},
			108: {CoreID: 63, SocketID: 3},
			228: {CoreID: 63, SocketID: 3},
			107: {CoreID: 62, SocketID: 3},
			227: {CoreID: 62, SocketID: 3},
			167: {CoreID: 62, SocketID: 3},
			287: {CoreID: 62, SocketID: 3},
			55:  {CoreID: 6, SocketID: 3},
			175: {CoreID: 6, SocketID: 3},
			115: {CoreID: 6, SocketID: 3},
			235: {CoreID: 6, SocketID: 3},
			56:  {CoreID: 7, SocketID: 3},
			176: {CoreID: 7, SocketID: 3},
			116: {CoreID: 7, SocketID: 3},
			236: {CoreID: 7, SocketID: 3},
			63:  {CoreID: 14, SocketID: 3},
			183: {CoreID: 14, SocketID: 3},
			123: {CoreID: 14, SocketID: 3},
			243: {CoreID: 14, SocketID: 3},
			64:  {CoreID: 15, SocketID: 3},
			184: {CoreID: 15, SocketID: 3},
			124: {CoreID: 15, SocketID: 3},
			244: {CoreID: 15, SocketID: 3},
			69:  {CoreID: 22, SocketID: 3},
			189: {CoreID: 22, SocketID: 3},
			129: {CoreID: 22, SocketID: 3},
			249: {CoreID: 22, SocketID: 3},
			70:  {CoreID: 23, SocketID: 3},
			190: {CoreID: 23, SocketID: 3},
			130: {CoreID: 23, SocketID: 3},
			250: {CoreID: 23, SocketID: 3},
			77:  {CoreID: 30, SocketID: 3},
			197: {CoreID: 30, SocketID: 3},
			137: {CoreID: 30, SocketID: 3},
			257: {CoreID: 30, SocketID: 3},
			78:  {CoreID: 31, SocketID: 3},
			198: {CoreID: 31, SocketID: 3},
			138: {CoreID: 31, SocketID: 3},
			258: {CoreID: 31, SocketID: 3},
			85:  {CoreID: 38, SocketID: 3},
			205: {CoreID: 38, SocketID: 3},
			145: {CoreID: 38, SocketID: 3},
			265: {CoreID: 38, SocketID: 3},
			86:  {CoreID: 39, SocketID: 3},
			206: {CoreID: 39, SocketID: 3},
			146: {CoreID: 39, SocketID: 3},
			266: {CoreID: 39, SocketID: 3},
			93:  {CoreID: 46, SocketID: 3},
			213: {CoreID: 46, SocketID: 3},
			153: {CoreID: 46, SocketID: 3},
			273: {CoreID: 46, SocketID: 3},
			94:  {CoreID: 47, SocketID: 3},
			214: {CoreID: 47, SocketID: 3},
			154: {CoreID: 47, SocketID: 3},
			274: {CoreID: 47, SocketID: 3},
			101: {CoreID: 54, SocketID: 3},
			221: {CoreID: 54, SocketID: 3},
			161: {CoreID: 54, SocketID: 3},
			281: {CoreID: 54, SocketID: 3},
			102: {CoreID: 55, SocketID: 3},
			222: {CoreID: 55, SocketID: 3},
			162: {CoreID: 55, SocketID: 3},
			282: {CoreID: 55, SocketID: 3},
		},
	}
)
