// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * 22-Nov-09 Benoy R Nair	First draft
 * 23-Nov-09 Benoy R Nair	For SQLDriverConnect()
 * 23-Nov-09 Benoy R Nair	For SQLGetInfo()
 */
#include "cgodbc.h"

int GO_SetEnvAttr ( SQLHANDLE environmentHandle
	, SQLINTEGER attribute
	, int value
	, SQLINTEGER stringLength )
{
	return ( SQLSetEnvAttr ( ( SQLHENV ) environmentHandle
		, attribute
		, ( SQLPOINTER ) value
		, stringLength ) );
}

int GO_DriverConnect ( SQLHANDLE connectionHandle
	, int windowHandle
	, char * inConnectionString
	, SQLSMALLINT stringLength1
	, SQLCHAR * outConnectionString
	, SQLSMALLINT bufferLength
	, SQLSMALLINT * stringLength2Ptr
	, SQLUSMALLINT driverCompletion )
{
	return ( SQLDriverConnect ( ( SQLHDBC ) connectionHandle
		, ( SQLHWND ) windowHandle
		, ( SQLCHAR * ) inConnectionString
		, stringLength1
		, outConnectionString
		, bufferLength
		, stringLength2Ptr
		, driverCompletion ) );
}

int GO_GetInfo_String ( SQLHANDLE connectionHandle
	, SQLUSMALLINT infoType
	, SQLCHAR * infoValue
	, SQLSMALLINT bufferLength
	, SQLSMALLINT * stringLength )
{
	return ( SQLGetInfo ( ( SQLHDBC ) connectionHandle
		, infoType
		, ( SQLPOINTER ) infoValue
		, bufferLength
		, stringLength ) );
}

int GO_GetInfo_Uint ( SQLHANDLE connectionHandle
	, SQLUSMALLINT infoType
	, SQLUSMALLINT * infoValue )
{
	return ( SQLGetInfo ( ( SQLHDBC ) connectionHandle
		, infoType
		, ( SQLPOINTER ) infoValue
		, 0
		, 0 ) );
}

int GO_GetInfo_Int ( SQLHANDLE connectionHandle
	, SQLUSMALLINT infoType
	, SQLSMALLINT * infoValue )
{
	return ( SQLGetInfo ( ( SQLHDBC ) connectionHandle
		, infoType
		, ( SQLPOINTER ) infoValue
		, 0
		, 0 ) );
}
