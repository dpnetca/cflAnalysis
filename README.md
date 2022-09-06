# cflAnalysis

## Under Development
This is very much still a under development, plan to eventually move the pkg/cfl into a seperate stand alone package once complete


## Overview
Leveraging the CFL API to do some analysis http://api.cfl.ca/docs 


## Plans
* create CLI script for analysis options with parameterized inputs
* come up wth consistent strategy for where and API calls are done
  * option 1 - (current preferred approach) as is done in schedule.go create a struct
  * option 2 - gather the information outside the analysis functions and pass it in
  * option 3 - TBD
* update overUnder functions to return information instead of printing it
* complete the pkg/cfl api package and seperate into it's own stand alone package
* add some rudementary preditive analysis, on unplayed games to predict winner/score
