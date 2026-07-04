@echo off
setlocal enabledelayedexpansion

cd /d "%~dp0"
if %errorlevel% neq 0 exit /b %errorlevel%

set "PYTHON={{ .PYTHON }}"

if exist "__pip_install__" (
  ".\python\!PYTHON!" -m ensurepip
  if %errorlevel% neq 0 exit /b %errorlevel%
  
  for %%f in ("__pip_install__"\*.whl) do (
    ".\python\!PYTHON!" -m pip install --no-cache-dir --no-index --find-links="__pip_install__" "%%f"
    if !errorlevel! neq 0 exit /b !errorlevel!
  )
  
  rd /s /q "__pip_install__"
  if %errorlevel% neq 0 exit /b %errorlevel%
)

".\python\!PYTHON!" {{ .RUN }} %*
if %errorlevel% neq 0 exit /b %errorlevel%

endlocal
