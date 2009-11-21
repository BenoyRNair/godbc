// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

	BUFFER_SIZE = 256;
)

var
(
	NULL_HANDLE GS_HANDLE;
)

func ( outputHandle * GS_HANDLE ) GS_AllocHandle ( handleType int
	, inputHandle GS_HANDLE )
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
