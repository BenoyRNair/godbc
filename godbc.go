// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * 22-Nov-09 Benoy R Nair	First draft
 * 23-Nov-09 Benoy R Nair	For SQLDriverConnect()
 * 23-Nov-09 Benoy R Nair	For SQLGetDiagRec(), SQLGetInfo()
 */
package godbc 

/*
#include <stdio.h>
#include <stdlib.h>

#include <sqltypes.h>
#include <sqlext.h>
#include <sql.h>

#include "cgodbc.h"
*/
import "C"
import "bytes"
import "unsafe"

type GS_HANDLE struct
{
	GsHandle C.SQLHANDLE;
}

const
(
	GS_ATTR_ODBC_VERSION = 200;
	GS_OV_ODBC3 = 3;
	GS_FETCH_FIRST = 2;
	GS_FETCH_NEXT = 1;
	GS_SUCCESS_WITH_INFO = 1;

// Handle type identifiers
	GS_HANDLE_ENV = 1;
	GS_HANDLE_DBC = 2;
	GS_HANDLE_STMT = 3;
	GS_HANDLE_DESC = 4;

// Options for GS_DriverConnect
	GS_DRIVER_NOPROMPT = 0;
	GS_DRIVER_COMPLETE = 1;
	GS_DRIVER_PROMPT = 2;
	GS_DRIVER_COMPLETE_REQUIRED = 3;

// Options for GS_GetInfo
	GS_MAX_CONCURRENT_ACTIVITIES = 1;
	GS_DBMS_NAME = 17;
	GS_DBMS_VER = 18;
	GS_GETDATA_EXTENSIONS = 81;

// GS_GETDATA_EXTENSIONS bitmasks
	GS_GD_ANY_COLUMN = uint32 ( 1 );
	GS_GD_ANY_ORDER = uint32 ( 2 );

	BUFFER_SIZE = 256;
)

var
(
	NULL_HANDLE GS_HANDLE;
)

func ( inputHandle * GS_HANDLE ) GS_AllocHandle ( handleType int
	, outputHandle * GS_HANDLE )
	int
{
	return ( int ( C.SQLAllocHandle ( C.SQLSMALLINT ( handleType )
		, unsafe.Pointer ( inputHandle.GsHandle )
		, &outputHandle.GsHandle ) ) );
}

func ( environmentHandle * GS_HANDLE ) GS_SetEnvAttr ( attribute int
	, value int
	, stringLength int )
	int
{
	return ( int ( C.GO_SetEnvAttr ( unsafe.Pointer ( environmentHandle.GsHandle )
		, C.SQLINTEGER ( attribute )
		, C.int ( value )
		, C.SQLINTEGER ( stringLength ) ) ) );
}

func ( environmentHandle * GS_HANDLE ) GS_DataSources ( direction int )
	( int, string, string )
{
	var dsn, desc * C.SQLCHAR;

	dsn = ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );
	desc = ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );

	var intDsn, intDesc C.SQLSMALLINT;

	returnInt := int ( C.SQLDataSources ( unsafe.Pointer ( environmentHandle.GsHandle )
		, C.SQLUSMALLINT ( direction )
		, dsn
		, BUFFER_SIZE
		, &intDsn
		, desc
		, BUFFER_SIZE
		, &intDesc ) );

	retServer := toStringByLength ( dsn, int ( intDsn ) );
	retDesc := toStringByLength ( desc, int ( intDesc ) );

	C.free ( unsafe.Pointer ( dsn ) );
	C.free ( unsafe.Pointer ( desc ) );

	return returnInt, retServer, retDesc
}

func ( environmentHandle * GS_HANDLE ) GS_Drivers ( direction int )
	( int, string, string )
{
	var driver, attr * C.SQLCHAR;

	driver = ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );
	attr = ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );

	var intDriver, intAttr C.SQLSMALLINT;

	returnInt := int ( C.SQLDrivers ( unsafe.Pointer ( environmentHandle.GsHandle )
		, C.SQLUSMALLINT ( direction )
		, driver
		, BUFFER_SIZE
		, &intDriver
		, attr
		, BUFFER_SIZE
		, &intAttr ) );

	retDriver := toStringByLength ( driver, int ( intDriver ) );
	retAttr := toStringByLength ( attr, int ( intAttr ) );

	C.free ( unsafe.Pointer ( driver ) );
	C.free ( unsafe.Pointer ( attr ) );

	return returnInt, retDriver, retAttr;
}

func ( connectionHandle * GS_HANDLE ) GS_DriverConnect ( windowHandle int
	, inConnection string
	, driverCompletion int )
	( int, string )
{
	var stringLength2 C.SQLSMALLINT;
	outConnectionString := ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );

	returnInt := int ( C.GO_DriverConnect ( unsafe.Pointer ( connectionHandle.GsHandle )
		, C.int ( windowHandle )
		, C.CString ( inConnection )
		, C.SQLSMALLINT ( len ( inConnection ) )
		, outConnectionString
		, BUFFER_SIZE
		, &stringLength2
		, C.SQLUSMALLINT ( driverCompletion ) ) );

	retOutConnectionString := toStringByLength ( outConnectionString, int ( stringLength2 ) );

	C.free ( unsafe.Pointer ( outConnectionString ) );

	return returnInt, retOutConnectionString;
}

func ( handle * GS_HANDLE ) GS_GetDiagRec ( handleType int
	, recNumber int )
	( int, string, int, string )
{
	sqlState := ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );
	var nativeError C.SQLINTEGER;
	messageText := ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );
	var textLength C.SQLSMALLINT;

	returnInt := int ( C.SQLGetDiagRec ( C.SQLSMALLINT ( handleType )
		, unsafe.Pointer ( handle.GsHandle )
		, C.SQLSMALLINT ( recNumber )
		, sqlState
		, &nativeError
		, messageText
		, BUFFER_SIZE
		, &textLength ) );

	C.free ( unsafe.Pointer ( sqlState ) );
	C.free ( unsafe.Pointer ( messageText ) );

	sqlStateString := toStringTillNull ( sqlState );
	messageTextString := toStringByLength ( messageText, int ( textLength ) );

	return returnInt, sqlStateString, int ( nativeError ), messageTextString;
}

func ( connectionHandle * GS_HANDLE ) GS_GetInfo_String ( infoType int )
	( int, string )
{
	stringValue := ( * C.SQLCHAR ) ( C.calloc ( BUFFER_SIZE, 1 ) );
	var length C.SQLSMALLINT;

	returnInt := int ( C.GO_GetInfo_String ( unsafe.Pointer ( connectionHandle.GsHandle )
		, C.SQLUSMALLINT ( infoType )
		, stringValue
		, BUFFER_SIZE
		, &length ) );

	returnString := toStringByLength ( stringValue, int ( length ) );

	C.free ( unsafe.Pointer ( stringValue ) );

	return returnInt, returnString;
}

func ( connectionHandle * GS_HANDLE ) GS_GetInfo_Uint ( infoType int )
	( int, uint32 )
{
	var intValue C.SQLUSMALLINT;

	returnInt := int ( C.GO_GetInfo_Uint ( unsafe.Pointer ( connectionHandle.GsHandle )
		, C.SQLUSMALLINT ( infoType )
		, &intValue ) );

	return returnInt, uint32 ( intValue );
}

func ( connectionHandle * GS_HANDLE ) GS_GetInfo_Int ( infoType int )
	( int, int )
{
	var intValue C.SQLSMALLINT;

	returnInt := int ( C.GO_GetInfo_Int ( unsafe.Pointer ( connectionHandle.GsHandle )
		, C.SQLUSMALLINT ( infoType )
		, &intValue ) );

	return returnInt, int ( intValue );
}

func toStringByLength ( buf * C.SQLCHAR, length int )
	string
{
	strbuf := make ( []byte, length );

	for j := int (0); j < length; j++
	{
		strbuf [j] = * ( * byte ) ( unsafe.Pointer ( uintptr ( unsafe.Pointer ( buf ) ) + uintptr ( j ) ) );
	}

	return bytes.NewBuffer ( strbuf ).String();
}

func toStringTillNull ( buf * C.SQLCHAR )
	string
{
	var length int;
	for length = 0 ; * ( * byte ) ( unsafe.Pointer ( uintptr ( unsafe.Pointer ( buf ) ) + uintptr ( length ) ) ) != 0; length++
	{
	}

	strbuf := make ( []byte, length);

	for j:= int (0); j < length; j++
	{
		strbuf [j] = * ( * byte ) ( unsafe.Pointer ( uintptr ( unsafe.Pointer ( buf ) ) + uintptr ( j ) ) );
	}

	return bytes.NewBuffer ( strbuf ).String();
}

func GS_Succeeded ( rc int )
	bool
{
	return ( ( ( rc ) & ( ^1 ) ) == 0 )
}
