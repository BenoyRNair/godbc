// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "godbc"
import "fmt"

var env godbc.GS_HANDLE;

func init()
{
	env.GS_AllocHandle ( 1, godbc.NULL_HANDLE );
	env.GS_SetEnvAttr ( godbc.GS_ATTR_ODBC_VERSION, godbc.GS_OV_ODBC3, 0 );
}

func main()
{
	listDataSources();
	listDrivers();
}

func listDataSources()
{
	fmt.Printf ( "Listing Data Sources (odbcinst -q -s) >>\n" );

	var x int;
	var dsn, desc string;

	direction := godbc.GS_FETCH_FIRST;

	for
	{
		x, dsn, desc = env.GS_DataSources ( direction );

		if ( ! godbc.GS_Succeeded ( x ) )
		{
			break;
		}

		direction = godbc.GS_FETCH_NEXT;

		fmt.Printf ( "DSN: %s - Desc: %s\n", dsn, desc );

		if x == godbc.GS_SUCCESS_WITH_INFO
		{
			fmt.Printf ( "\tdata truncation\n" );
		}
	}
}

func listDrivers()
{
	fmt.Printf ( "Listing Drivers (odbcinst -q -d) >>\n" );

	var x int;
	var driver, attr string;

	direction := godbc.GS_FETCH_FIRST;

	for
	{
		x, driver, attr = env.GS_Drivers ( direction );

		if ( ! godbc.GS_Succeeded ( x ) )
		{
			break;
		}

		direction = godbc.GS_FETCH_NEXT;

		fmt.Printf ( "Driver: %s - Attr: %s\n", driver, attr );

		if x == godbc.GS_SUCCESS_WITH_INFO
		{
			fmt.Printf ( "\tdata truncation\n" );
		}
	}
}
