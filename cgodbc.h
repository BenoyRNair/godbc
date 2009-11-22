// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * 22-Nov-09 Benoy R Nair	First draft
 * 23-Nov-09 Benoy R Nair	For SQLDriverConnect()
 */
#ifndef _CGODBC_H_
#define _CGODBC_H_

#include <sqltypes.h>
#include <sql.h>
#include <sqlext.h>

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
#endif
