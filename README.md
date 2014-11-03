wat - Wolfram Alpha command line Tool
===

A simple command-line tool to query Wolfram Alpha's API

Installation
===
A simple go get will suffice: ```go get github.com/seen/wat```

Configuration
===
Wolfram Alpha requires an APP ID for querying it's API. You can
get one here: https://developer.wolframalpha.com/

`wat` loads the APP ID from the environment variable `WOLFRAM_API_ID`.
Set it with your APP ID from Wolfram Alpha and you're good to go!

Usage Examples
===

Epoch timestamps
```
~ » wat 1415034693291/1000 unix                                                                                        
Input interpretation
	1415034693291/1000 (Unix time)

Corresponding Gregorian time and date
	5:11:33 pm UTC  |  Monday, November 3, 2014

Time difference from now (5:56:32 pm UTC)
	44 minutes 60 seconds ago
	2700 seconds ago

Time offset from PST
	+8 hours

Clocks
	Coordinated Universal Time
	
	5:11:33 pm  UTC
	Monday, November 3               Portland, Oregon
	
	9:11:33 am  PST
	Monday, November 3
```

Byte conversions
```
~ » wat 12346578 bytes                                                                                                 
Input interpretation
	12346578 bytes

Unit conversions
	12.35 MB  (megabytes)
	0.01235 GB  (gigabytes)
	98.77 Mb  (megabits)
	0.09877 Gb  (gigabits)
	9.877×10^7 bits

Comparisons
	 ~~ 4 × data storage capacity of a 200-character per inch 2400-foot IBM 7330 7-track magnetic tape (~~ 2×10^7 b )
	 ~~ 8.6 × 3.5" floppy disk capacity (high density) ( 1440 kB )
	 ~~ 15 × 3.5" floppy disk capacity (Mac double-sided formatted) ( 800 KiB )

Interpretation
	information

Basic unit dimensions
	[information]

Corresponding quantities
	Time to transfer at T1 speed:
		  | 64 seconds
		  | 1.1 minutes
		  | (assuming T1 speed  =  1544 kb/s)
	Time to transfer at cable modem speed:
		  | (16 to 99) seconds
		  | (0.3 to 1.6) minutes
		  | (assuming cable modem speed ~~ 1 to 6 Mb/s)
	Time to transfer at ADSL (US DSL) speed:
		  | (4.1 to 66) seconds
		  | (0.069 to 1.1) minutes
		  | (assuming downstream ADSL speed ~~ 1.5 to 24 Mb/s)
	Time to transfer at 10BASE-T ethernet speed:
		  | 9.9 seconds
		  | (assuming 10BASE-T ethernet speed  =  10 Mb/s)
	Time to transfer at 100BASE-TX ethernet speed:
		  | 0.99 seconds
		  | (assuming 100BASE-TX ethernet speed  =  100 Mb/s)
```
