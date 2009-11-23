// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * 22-Nov-09 Benoy R Nair	First draft
 * 23-Nov-09 Benoy R Nair	For SQLDriverConnect()
 * 23-Nov-09 Benoy R Nair	For SQLGetDiagRec(), SQLGetInfo()
 */
package main

import "godbc"
import "fmt"
import "os"

var env, dbc, stmt godbc.GS_HANDLE;

func init()
{
	godbc.NULL_HANDLE.GS_AllocHandle ( godbc.GS_HANDLE_ENV, &env );
	env.GS_SetEnvAttr ( godbc.GS_ATTR_ODBC_VERSION, godbc.GS_OV_ODBC3, 0 );
	env.GS_AllocHandle ( godbc.GS_HANDLE_DBC, &dbc );
	dbc.GS_AllocHandle ( godbc.GS_HANDLE_STMT, &stmt );
}

func main()
{
	listDataSources();
	listDrivers();
	connect();
	testGetInfo();
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

		fmt.Printf ( "DSN: %s Desc: %s\n", dsn, desc );

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

		fmt.Printf ( "Driver: %s Attr: %s\n", driver, attr );

		if x == godbc.GS_SUCCESS_WITH_INFO
		{
			fmt.Printf ( "\tdata truncation\n" );
		}
	}
}

func connect()
{
	x, str := dbc.GS_DriverConnect ( 0, "DSN=dsn1mysql;UID=root;PWD=password", godbc.GS_DRIVER_COMPLETE );

	if ( godbc.GS_Succeeded ( x ) )
	{
		fmt.Printf ( "Connected. Details: %s\n", str );
	}
	else
	{
		fmt.Printf ( "Unable to connect. Check the login credentials provided in the code.\n" );
		extractError ( "connect()", &dbc, godbc.GS_HANDLE_DBC );
		os.Exit (1);
	}
}

func testGetInfo()
{
	_, dbmsName := dbc.GS_GetInfo_String ( godbc.GS_DBMS_NAME );
	_, dbmsVer := dbc.GS_GetInfo_String ( godbc.GS_DBMS_VER );
	_, dataSupport := dbc.GS_GetInfo_Uint ( godbc.GS_GETDATA_EXTENSIONS );
	_, maxConcurrentActivities := dbc.GS_GetInfo_Int ( godbc.GS_MAX_CONCURRENT_ACTIVITIES );

	fmt.Printf ( "DBMS Name: %s\n", dbmsName );
	fmt.Printf ( "DBMS Version: %s\n", dbmsVer );

	if ( dataSupport & godbc.GS_GD_ANY_ORDER ) != uint32 ( 0 )
	{
		fmt.Printf ( "Columns can be retrieved in any order\n" );
	}
	else
	{
		fmt.Printf ( "Columns must be retrieved in order\n" );
	}

	if ( dataSupport & godbc.GS_GD_ANY_COLUMN ) != uint32 ( 0 )
	{
		fmt.Printf ( "Can retrieve columns before last bound one\n" );
	}
	else
	{
		fmt.Printf ( "Columns must be bound after last bound one\n" );
	}

	if maxConcurrentActivities == 0
	{
		fmt.Printf ( "Max Concurrent Acitivities: No limit or undefined\n" );
	}
	else
	{
		fmt.Printf ( "Max Concurrent Activities: %d\n", maxConcurrentActivities );
	}
}

func extractError ( function string
	, handle * godbc.GS_HANDLE
	, handleType int )
{
	var x, nativeError int;
	var sqlState, messageText string;
	recNumber := 1;

	fmt.Printf ( "The driver reported the following diagnostics whilst running %s\n\n", function );

	for
	{
		x, sqlState, nativeError, messageText = handle.GS_GetDiagRec ( handleType, recNumber );

		recNumber++;

		if ( godbc.GS_Succeeded ( x ) )
		{
			fmt.Printf ( "%s:%d:%d:%s\n", sqlState, recNumber, nativeError, messageText );
		}
		else
		{
			fmt.Printf ( "unable to extract error\n" );
			break;
		}
	}
}
