// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "cgodbc.h"

#include <stdio.h>

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
