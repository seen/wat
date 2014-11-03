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
~ » wat 1415034693291 milliseconds                                                                                                                                                                                                                                           sean@phalanx
Input interpretation
	1415034693291 ms  (milliseconds)

Unit conversions
	1.415×10^9 seconds
	2.358×10^7 minutes
	393065 hours
	16378 days
	44.84 average Gregorian years

Comparisons as period
	 ~~ 0.27 × orbital period of Neptune (~~ 165 Julian years )
	 ~~ 0.53 × orbital period of Uranus (~~ 84 Julian years )
	 ~~ 1.5 × orbital period of Saturn (~~ 29 Julian years )

Comparisons as half‐life
	 ~~ 8.5 × half-life of cobalt-60 ( 1.6635×10^8 s )
	 ~~ 17 × half-life of sodium-22 ( 8.2108×10^7 s )

Offsets from current time
	1.415×10^12 ms  (milliseconds) from now | 2:25:26 am PDT  |  Saturday, September 6, 2059
		1.415×10^12 ms  (milliseconds) before now | 4:02:19 pm PST  |  Wednesday, December 31, 1969

Interpretations
	time
	period
```

Byte conversions
```
~ » wat 12346578 bytes                                                                                                                                                                                                                                                       sean@phalanx
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
