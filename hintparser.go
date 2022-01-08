// Code generated by goyacc - DO NOT EDIT.

// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import __yyfmt__ "fmt"

import (
	"strconv"

	"github.com/dubbogo/parser/ast"
	"github.com/dubbogo/parser/model"
)

type yyhintSymType struct {
	yys    int
	ident  string
	number uint64
	hint   *ast.TableOptimizerHint
	hints  []*ast.TableOptimizerHint
	table  ast.HintTable
}

type yyhintXError struct {
	state, xsym int
}

const (
	yyhintDefault           = 57389
	yyhintEofCode           = 57344
	yyhintErrCode           = 57345
	hintBKA                 = 57354
	hintBNL                 = 57356
	hintDupsWeedOut         = 57385
	hintFalse               = 57381
	hintFirstMatch          = 57386
	hintGB                  = 57384
	hintHashJoin            = 57358
	hintIdentifier          = 57347
	hintIndexMerge          = 57362
	hintIntLit              = 57346
	hintJoinFixedOrder      = 57350
	hintJoinOrder           = 57351
	hintJoinPrefix          = 57352
	hintJoinSuffix          = 57353
	hintLooseScan           = 57387
	hintMB                  = 57383
	hintMRR                 = 57364
	hintMaterialization     = 57388
	hintMaxExecutionTime    = 57372
	hintMerge               = 57360
	hintNoBKA               = 57355
	hintNoBNL               = 57357
	hintNoHashJoin          = 57359
	hintNoICP               = 57366
	hintNoIndexMerge        = 57363
	hintNoMRR               = 57365
	hintNoMerge             = 57361
	hintNoRangeOptimization = 57367
	hintNoSemijoin          = 57371
	hintNoSkipScan          = 57369
	hintOLAP                = 57377
	hintOLTP                = 57378
	hintQBName              = 57375
	hintResourceGroup       = 57374
	hintSemijoin            = 57370
	hintSetVar              = 57373
	hintSingleAtIdentifier  = 57348
	hintSkipScan            = 57368
	hintStringLit           = 57349
	hintTiFlash             = 57380
	hintTiKV                = 57379
	hintTrue                = 57382
	hintXID                 = 57376

	yyhintMaxDepth = 200
	yyhintTabOfs   = -101
)

var (
	yyhintPrec = map[int]int{}

	yyhintXLAT = map[int]int{
		41:    0,  // ')' (81x)
		57354: 1,  // hintBKA (77x)
		57356: 2,  // hintBNL (77x)
		57362: 3,  // hintIndexMerge (77x)
		57350: 4,  // hintJoinFixedOrder (77x)
		57351: 5,  // hintJoinOrder (77x)
		57352: 6,  // hintJoinPrefix (77x)
		57353: 7,  // hintJoinSuffix (77x)
		57372: 8,  // hintMaxExecutionTime (77x)
		57360: 9,  // hintMerge (77x)
		57364: 10, // hintMRR (77x)
		57355: 11, // hintNoBKA (77x)
		57357: 12, // hintNoBNL (77x)
		57359: 13, // hintNoHashJoin (77x)
		57366: 14, // hintNoICP (77x)
		57363: 15, // hintNoIndexMerge (77x)
		57361: 16, // hintNoMerge (77x)
		57365: 17, // hintNoMRR (77x)
		57367: 18, // hintNoRangeOptimization (77x)
		57371: 19, // hintNoSemijoin (77x)
		57369: 20, // hintNoSkipScan (77x)
		57375: 21, // hintQBName (77x)
		57374: 22, // hintResourceGroup (77x)
		57370: 23, // hintSemijoin (77x)
		57373: 24, // hintSetVar (77x)
		57368: 25, // hintSkipScan (77x)
		44:    26, // ',' (73x)
		57385: 27, // hintDupsWeedOut (64x)
		57386: 28, // hintFirstMatch (64x)
		57387: 29, // hintLooseScan (64x)
		57388: 30, // hintMaterialization (64x)
		57381: 31, // hintFalse (61x)
		57384: 32, // hintGB (61x)
		57358: 33, // hintHashJoin (61x)
		57347: 34, // hintIdentifier (61x)
		57383: 35, // hintMB (61x)
		57377: 36, // hintOLAP (61x)
		57378: 37, // hintOLTP (61x)
		57380: 38, // hintTiFlash (61x)
		57379: 39, // hintTiKV (61x)
		57382: 40, // hintTrue (61x)
		57348: 41, // hintSingleAtIdentifier (48x)
		46:    42, // '.' (40x)
		61:    43, // '=' (40x)
		40:    44, // '(' (31x)
		57376: 45, // hintXID (17x)
		57344: 46, // $end (15x)
		57395: 47, // Identifier (12x)
		57401: 48, // QueryBlockOpt (9x)
		57346: 49, // hintIntLit (5x)
		57392: 50, // HintTable (4x)
		57390: 51, // CommaOpt (2x)
		57349: 52, // hintStringLit (2x)
		57393: 53, // HintTableList (2x)
		57398: 54, // JoinOrderOptimizerHintName (2x)
		57399: 55, // NullaryHintName (2x)
		57403: 56, // SubqueryOptimizerHintName (2x)
		57406: 57, // SubqueryStrategy (2x)
		57408: 58, // TableOptimizerHintOpt (2x)
		57409: 59, // UnsupportedIndexLevelOptimizerHintName (2x)
		57410: 60, // UnsupportedTableLevelOptimizerHintName (2x)
		57411: 61, // Value (2x)
		57391: 62, // HintIndexList (1x)
		57394: 63, // HintTableListOpt (1x)
		57396: 64, // IndexNameList (1x)
		57397: 65, // IndexNameListOpt (1x)
		57400: 66, // OptimizerHintList (1x)
		57402: 67, // Start (1x)
		57404: 68, // SubqueryStrategies (1x)
		57405: 69, // SubqueryStrategiesOpt (1x)
		57389: 70, // $default (0x)
		57345: 71, // error (0x)
		57407: 72, // SupportedTableLevelOptimizerHintName (0x)
	}

	yyhintSymNames = []string{
		"')'",
		"hintBKA",
		"hintBNL",
		"hintIndexMerge",
		"hintJoinFixedOrder",
		"hintJoinOrder",
		"hintJoinPrefix",
		"hintJoinSuffix",
		"hintMaxExecutionTime",
		"hintMerge",
		"hintMRR",
		"hintNoBKA",
		"hintNoBNL",
		"hintNoHashJoin",
		"hintNoICP",
		"hintNoIndexMerge",
		"hintNoMerge",
		"hintNoMRR",
		"hintNoRangeOptimization",
		"hintNoSemijoin",
		"hintNoSkipScan",
		"hintQBName",
		"hintResourceGroup",
		"hintSemijoin",
		"hintSetVar",
		"hintSkipScan",
		"','",
		"hintDupsWeedOut",
		"hintFirstMatch",
		"hintLooseScan",
		"hintMaterialization",
		"hintFalse",
		"hintGB",
		"hintHashJoin",
		"hintIdentifier",
		"hintMB",
		"hintOLAP",
		"hintOLTP",
		"hintTiFlash",
		"hintTiKV",
		"hintTrue",
		"hintSingleAtIdentifier",
		"'.'",
		"'='",
		"'('",
		"hintXID",
		"$end",
		"Identifier",
		"QueryBlockOpt",
		"hintIntLit",
		"HintTable",
		"CommaOpt",
		"hintStringLit",
		"HintTableList",
		"JoinOrderOptimizerHintName",
		"NullaryHintName",
		"SubqueryOptimizerHintName",
		"SubqueryStrategy",
		"TableOptimizerHintOpt",
		"UnsupportedIndexLevelOptimizerHintName",
		"UnsupportedTableLevelOptimizerHintName",
		"Value",
		"HintIndexList",
		"HintTableListOpt",
		"IndexNameList",
		"IndexNameListOpt",
		"OptimizerHintList",
		"Start",
		"SubqueryStrategies",
		"SubqueryStrategiesOpt",
		"$default",
		"error",
		"SupportedTableLevelOptimizerHintName",
	}

	yyhintTokenLiteralStrings = map[int]string{
		57354: "BKA",
		57356: "BNL",
		57362: "INDEX_MERGE",
		57350: "JOIN_FIXED_ORDER",
		57351: "JOIN_ORDER",
		57352: "JOIN_PREFIX",
		57353: "JOIN_SUFFIX",
		57372: "MAX_EXECUTION_TIME",
		57360: "MERGE",
		57364: "MRR",
		57355: "NO_BKA",
		57357: "NO_BNL",
		57359: "NO_HASH_JOIN",
		57366: "NO_ICP",
		57363: "NO_INDEX_MERGE",
		57361: "NO_MERGE",
		57365: "NO_MRR",
		57367: "NO_RANGE_OPTIMIZATION",
		57371: "NO_SEMIJOIN",
		57369: "NO_SKIP_SCAN",
		57375: "QB_NAME",
		57374: "RESOURCE_GROUP",
		57370: "SEMIJOIN",
		57373: "SET_VAR",
		57368: "SKIP_SCAN",
		57385: "DUPSWEEDOUT",
		57386: "FIRSTMATCH",
		57387: "LOOSESCAN",
		57388: "MATERIALIZATION",
		57381: "FALSE",
		57384: "GB",
		57358: "HASH_JOIN",
		57383: "MB",
		57377: "OLAP",
		57378: "OLTP",
		57380: "TIFLASH",
		57379: "TIKV",
		57382: "TRUE",
		57348: "identifier with single leading at",
		57376: "XID",
		57346: "a 64-bit unsigned integer",
	}

	yyhintReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {67, 1},
		2:   {66, 1},
		3:   {66, 3},
		4:   {58, 4},
		5:   {58, 4},
		6:   {58, 4},
		7:   {58, 4},
		8:   {58, 5},
		9:   {58, 5},
		10:  {58, 6},
		11:  {58, 4},
		12:  {58, 4},
		13:  {58, 4},
		14:  {58, 4},
		15:  {48, 0},
		16:  {48, 1},
		17:  {51, 0},
		18:  {51, 1},
		19:  {63, 1},
		20:  {63, 1},
		21:  {53, 2},
		22:  {53, 3},
		23:  {50, 2},
		24:  {50, 4},
		25:  {62, 4},
		26:  {65, 0},
		27:  {65, 1},
		28:  {64, 1},
		29:  {64, 3},
		30:  {69, 0},
		31:  {69, 1},
		32:  {68, 1},
		33:  {68, 3},
		34:  {61, 1},
		35:  {61, 1},
		36:  {61, 1},
		37:  {54, 1},
		38:  {54, 1},
		39:  {54, 1},
		40:  {60, 1},
		41:  {60, 1},
		42:  {60, 1},
		43:  {60, 1},
		44:  {60, 1},
		45:  {60, 1},
		46:  {60, 1},
		47:  {72, 1},
		48:  {59, 1},
		49:  {59, 1},
		50:  {59, 1},
		51:  {59, 1},
		52:  {59, 1},
		53:  {59, 1},
		54:  {59, 1},
		55:  {56, 1},
		56:  {56, 1},
		57:  {57, 1},
		58:  {57, 1},
		59:  {57, 1},
		60:  {57, 1},
		61:  {55, 1},
		62:  {47, 1},
		63:  {47, 1},
		64:  {47, 1},
		65:  {47, 1},
		66:  {47, 1},
		67:  {47, 1},
		68:  {47, 1},
		69:  {47, 1},
		70:  {47, 1},
		71:  {47, 1},
		72:  {47, 1},
		73:  {47, 1},
		74:  {47, 1},
		75:  {47, 1},
		76:  {47, 1},
		77:  {47, 1},
		78:  {47, 1},
		79:  {47, 1},
		80:  {47, 1},
		81:  {47, 1},
		82:  {47, 1},
		83:  {47, 1},
		84:  {47, 1},
		85:  {47, 1},
		86:  {47, 1},
		87:  {47, 1},
		88:  {47, 1},
		89:  {47, 1},
		90:  {47, 1},
		91:  {47, 1},
		92:  {47, 1},
		93:  {47, 1},
		94:  {47, 1},
		95:  {47, 1},
		96:  {47, 1},
		97:  {47, 1},
		98:  {47, 1},
		99:  {47, 1},
		100: {47, 1},
	}

	yyhintXErrors = map[yyhintXError]string{}

	yyhintParseTab = [145][]uint16{
		// 0
		{1: 119, 121, 126, 105, 116, 117, 118, 110, 124, 127, 120, 122, 123, 129, 135, 125, 128, 130, 134, 132, 113, 112, 133, 111, 131, 45: 115, 54: 106, 114, 109, 58: 104, 108, 107, 66: 103, 102},
		{46: 101},
		{1: 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 217, 45: 84, 100, 51: 244},
		{1: 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 45: 99, 99},
		{44: 241},
		// 5
		{44: 237},
		{44: 229},
		{44: 212},
		{44: 200},
		{44: 196},
		// 10
		{44: 191},
		{44: 188},
		{44: 185},
		{44: 181},
		{44: 136},
		// 15
		{44: 64},
		{44: 63},
		{44: 62},
		{44: 61},
		{44: 60},
		// 20
		{44: 59},
		{44: 58},
		{44: 57},
		{44: 56},
		{44: 55},
		// 25
		{44: 53},
		{44: 52},
		{44: 51},
		{44: 50},
		{44: 49},
		// 30
		{44: 48},
		{44: 47},
		{44: 46},
		{44: 45},
		{44: 40},
		// 35
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 139, 49: 140, 52: 138, 61: 137},
		{180},
		{67},
		{66},
		{65},
		// 40
		{39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39},
		{38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38},
		{37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37, 37},
		{36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36},
		{35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35},
		// 45
		{34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34},
		{33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33},
		{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32},
		{31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31},
		{30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30},
		// 50
		{29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29},
		{28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28},
		{27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27},
		{26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26},
		{25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25},
		// 55
		{24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24, 24},
		{23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23, 23},
		{22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22},
		{21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21},
		{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20},
		// 60
		{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19},
		{18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18, 18},
		{17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17, 17},
		{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16},
		{15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15},
		// 65
		{14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14},
		{13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		{12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		// 70
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		// 75
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1: 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 87, 45: 87, 87},
		// 80
		{86, 41: 183, 48: 182},
		{184},
		{85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 85, 49: 85},
		{1: 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 45: 88, 88},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 186},
		// 85
		{187},
		{1: 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 89, 45: 89, 89},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 189},
		{190},
		{1: 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 45: 90, 90},
		// 90
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 192},
		{43: 193},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 139, 49: 140, 52: 138, 61: 194},
		{195},
		{1: 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 45: 91, 91},
		// 95
		{41: 183, 48: 197, 86},
		{49: 198},
		{199},
		{1: 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 92, 45: 92, 92},
		{86, 27: 86, 86, 86, 86, 41: 183, 48: 201},
		// 100
		{71, 27: 205, 206, 207, 208, 57: 204, 68: 203, 202},
		{211},
		{70, 26: 209},
		{69, 26: 69},
		{44, 26: 44},
		// 105
		{43, 26: 43},
		{42, 26: 42},
		{41, 26: 41},
		{27: 205, 206, 207, 208, 57: 210},
		{68, 26: 68},
		// 110
		{1: 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 45: 93, 93},
		{1: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 27: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 183, 48: 214, 62: 213},
		{228},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 215, 50: 216},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 183, 225, 48: 224},
		// 115
		{84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 217, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 84, 51: 218},
		{83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 27: 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 83, 45: 83},
		{75, 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 221, 64: 220, 219},
		{76},
		{74, 26: 222},
		// 120
		{73, 26: 73},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 223},
		{72, 26: 72},
		{78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78, 78},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 226},
		// 125
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 183, 48: 227},
		{77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77, 77},
		{1: 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 45: 94, 94},
		{86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 27: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 183, 48: 232, 53: 231, 63: 230},
		{236},
		// 130
		{82, 26: 234},
		{81, 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 215, 50: 233},
		{80, 26: 80},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 215, 50: 235},
		{79, 26: 79},
		// 135
		{1: 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 45: 95, 95},
		{1: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 27: 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 183, 48: 239, 53: 238},
		{240, 26: 234},
		{1: 146, 148, 154, 142, 143, 144, 145, 164, 152, 156, 147, 149, 151, 158, 155, 153, 157, 159, 163, 161, 167, 166, 162, 165, 160, 27: 176, 177, 178, 179, 172, 175, 150, 141, 174, 168, 169, 171, 170, 173, 47: 215, 50: 233},
		{1: 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 45: 96, 96},
		// 140
		{86, 41: 183, 48: 242},
		{243},
		{1: 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 97, 45: 97, 97},
		{1: 119, 121, 126, 105, 116, 117, 118, 110, 124, 127, 120, 122, 123, 129, 135, 125, 128, 130, 134, 132, 113, 112, 133, 111, 131, 45: 115, 54: 106, 114, 109, 58: 245, 108, 107},
		{1: 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 98, 45: 98, 98},
	}
)

var yyhintDebug = 0

type yyhintLexer interface {
	Lex(lval *yyhintSymType) int
	Error(s string)
}

type yyhintLexerEx interface {
	yyhintLexer
	Reduced(rule, state int, lval *yyhintSymType) bool
}

func yyhintSymName(c int) (s string) {
	x, ok := yyhintXLAT[c]
	if ok {
		return yyhintSymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yyhintlex1(yylex yyhintLexer, lval *yyhintSymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyhintEofCode
	}
	if yyhintDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yyhintSymName(n), n, n, lval)
	}
	return n
}

func yyhintParse(yylex yyhintLexer, parser *hintParser) int {
	const yyError = 71

	yyEx, _ := yylex.(yyhintLexerEx)
	var yyn int
	var yylval yyhintSymType
	var yyVAL yyhintSymType
	yyS := make([]yyhintSymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyhintDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yyhintSymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yyhintlex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyhintXLAT[yychar]; !ok {
			yyxchar = len(yyhintSymNames) // > tab width
		}
	}
	if yyhintDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyhintParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyhintTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyhintDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyhintDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyhintDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yyhintSymName(yychar), yystate)
			}
			msg, ok := yyhintXErrors[yyhintXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyhintXErrors[yyhintXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyhintXErrors[yyhintXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyhintXErrors[yyhintXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyhintTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yyhintSymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyhintParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyhintTabOfs
					if yyn > 0 { // hit
						if yyhintDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyhintDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyhintDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyhintDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyhintSymName(yychar))
			}
			if yychar == yyhintEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyhintReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yyhintSymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyhintParseTab[yyS[yyp].yys][x]) + yyhintTabOfs
	/* reduction by production r */
	if yyhintDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yyhintSymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			parser.result = yyS[yypt-0].hints
		}
	case 2:
		{
			if yyS[yypt-0].hint != nil {
				yyVAL.hints = []*ast.TableOptimizerHint{yyS[yypt-0].hint}
			}
		}
	case 3:
		{
			if yyS[yypt-0].hint != nil {
				yyVAL.hints = append(yyS[yypt-2].hints, yyS[yypt-0].hint)
			} else {
				yyVAL.hints = yyS[yypt-2].hints
			}
		}
	case 4:
		{
			parser.warnUnsupportedHint(yyS[yypt-3].ident)
			yyVAL.hint = nil
		}
	case 5:
		{
			parser.warnUnsupportedHint(yyS[yypt-3].ident)
			yyVAL.hint = nil
		}
	case 6:
		{
			parser.warnUnsupportedHint(yyS[yypt-3].ident)
			yyVAL.hint = nil
		}
	case 7:
		{
			parser.warnUnsupportedHint(yyS[yypt-3].ident)
			yyVAL.hint = nil
		}
	case 8:
		{
			parser.warnUnsupportedHint(yyS[yypt-4].ident)
			yyVAL.hint = nil
		}
	case 9:
		{
			yyVAL.hint = &ast.TableOptimizerHint{
				HintName: model.NewCIStr(yyS[yypt-4].ident),
				QBName:   model.NewCIStr(yyS[yypt-2].ident),
				HintData: yyS[yypt-1].number,
			}
		}
	case 10:
		{
			parser.warnUnsupportedHint(yyS[yypt-5].ident)
			yyVAL.hint = nil
		}
	case 11:
		{
			parser.warnUnsupportedHint(yyS[yypt-3].ident)
			yyVAL.hint = nil
		}
	case 12:
		{
			yyVAL.hint = &ast.TableOptimizerHint{
				HintName: model.NewCIStr(yyS[yypt-3].ident),
				QBName:   model.NewCIStr(yyS[yypt-1].ident),
			}
		}
	case 13:
		{
			yyVAL.hint = &ast.TableOptimizerHint{
				HintName: model.NewCIStr(yyS[yypt-3].ident),
				QBName:   model.NewCIStr(yyS[yypt-1].ident),
			}
		}
	case 14:
		{
			yyVAL.hint = &ast.TableOptimizerHint{
				HintName: model.NewCIStr(yyS[yypt-3].ident),
				HintData: model.NewCIStr(yyS[yypt-1].ident),
			}
		}
	case 15:
		{
			yyVAL.ident = ""
		}
	case 20:
		{
			yyVAL.hint = &ast.TableOptimizerHint{
				QBName: model.NewCIStr(yyS[yypt-0].ident),
			}
		}
	case 21:
		{
			yyVAL.hint = &ast.TableOptimizerHint{
				Tables: []ast.HintTable{yyS[yypt-0].table},
				QBName: model.NewCIStr(yyS[yypt-1].ident),
			}
		}
	case 22:
		{
			h := yyS[yypt-2].hint
			h.Tables = append(h.Tables, yyS[yypt-0].table)
			yyVAL.hint = h
		}
	case 23:
		{
			yyVAL.table = ast.HintTable{
				TableName: model.NewCIStr(yyS[yypt-1].ident),
				QBName:    model.NewCIStr(yyS[yypt-0].ident),
			}
		}
	case 24:
		{
			yyVAL.table = ast.HintTable{
				DBName:    model.NewCIStr(yyS[yypt-3].ident),
				TableName: model.NewCIStr(yyS[yypt-1].ident),
				QBName:    model.NewCIStr(yyS[yypt-0].ident),
			}
		}
	case 25:
		{
			h := yyS[yypt-0].hint
			h.Tables = []ast.HintTable{yyS[yypt-2].table}
			h.QBName = model.NewCIStr(yyS[yypt-3].ident)
			yyVAL.hint = h
		}
	case 26:
		{
			yyVAL.hint = &ast.TableOptimizerHint{}
		}
	case 28:
		{
			yyVAL.hint = &ast.TableOptimizerHint{
				Indexes: []model.CIStr{model.NewCIStr(yyS[yypt-0].ident)},
			}
		}
	case 29:
		{
			h := yyS[yypt-2].hint
			h.Indexes = append(h.Indexes, model.NewCIStr(yyS[yypt-0].ident))
			yyVAL.hint = h
		}
	case 36:
		{
			yyVAL.ident = strconv.FormatUint(yyS[yypt-0].number, 10)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}