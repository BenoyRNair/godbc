// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * 22-Nov-09 Benoy R Nair	First draft
 * 23-Nov-09 Benoy R Nair	For SQLDriverConnect()
 * 23-Nov-09 Benoy R Nair	For SQLGetInfo()
 * 27-Nov-09 Benoy R Nair	For SQLTables(), SQLNumResultCols(), SQLGetData()
 */
#ifndef _CGODBC_H_
#define _CGODBC_H_

#include <sqltypes.h>
#include <sql.h>
#include <sqlext.h>

#include <stdio.h>
#include <stdlib.h>

int GO_SetEnvAttr ( SQLHANDLE environmentHandle
	, SQLINTEGER attribute
	, int value
	, SQLINTEGER stringLength );

int GO_DriverConnect ( SQLHANDLE connectionHandle
	, int windowHandle
	, char * inConnectionString
	, SQLSMALLINT stringLength1
	, SQLCHAR * outConnectionString
	, SQLSMALLINT bufferLength
	, SQLSMALLINT * stringLength2Ptr
	, SQLUSMALLINT driverCompletion );

int GO_GetInfo_String ( SQLHANDLE connectionHandle
	, SQLUSMALLINT infoType
	, SQLCHAR * infoValue
	, SQLSMALLINT bufferLength
	, SQLSMALLINT * stringLength );

int GO_GetInfo_Uint ( SQLHANDLE connectionHandle
	, SQLUSMALLINT infoType
	, SQLUSMALLINT * infoValue );

int GO_GetInfo_Int ( SQLHANDLE connectionHandle
	, SQLUSMALLINT infoType
	, SQLSMALLINT * infoValue );

int GO_Tables ( SQLHANDLE statementHandle
	, char * catalogName
	, SQLSMALLINT nameLength1
	, char * schemaName
	, SQLSMALLINT nameLength2
	, char * tableName
	, SQLSMALLINT nameLength3
	, char * tableType
	, SQLSMALLINT nameLength4 );

int GO_GetData_String ( SQLHANDLE statementHandle
	, SQLUSMALLINT columnNumber
	, SQLCHAR * targetValue
	, SQLINTEGER bufferLength
	, SQLINTEGER * indicator );
#endif
