// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef _CGODBC_H_
#define _CGODBC_H_

#include <sqltypes.h>
#include <sql.h>
#include <sqlext.h>

int GO_SetEnvAttr ( SQLHANDLE environmentHandle
	, SQLINTEGER attribute
	, int value
	, SQLINTEGER stringLength );
#endif
