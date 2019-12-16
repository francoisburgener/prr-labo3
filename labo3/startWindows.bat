@echo off
set /a i = %1-1
set /a cnt=0;
setlocal ENABLEDELAYEDEXPANSION
FOR %%i in (%*) DO (
	call set arg[!cnt!]=%%i 
	call set /a cnt+=1 
)
FOR /l %%G IN (0 1 %i%) DO (
	call set /a apt=%%G+1
	call start cmd /k go run main.go -proc %%G -N %1 -apt %%arg[!apt!]%%
)
endlocal