@ECHO OFF
mklink /D "C:\GO\app" "E:\Claims1MED"
mklink /D "C:\GO\engine\src" "\\192.168.10.220\GO.Dev\src"
mklink /D "C:\GO\resource\images" "\\192.168.10.220\wwwroot\MediCSV5\images"
mklink /D "C:\GO\resource\scripts" "\\192.168.10.220\wwwroot\MediCSV5\scripts"
mklink /D "C:\GO\resource\app\db" "\\192.168.10.220\wwwroot\MediCSV5\app\db"
mklink /D "C:\GO\resource\app\css" "\\192.168.10.220\wwwroot\MediCSV5\app\css"
mklink /D "C:\GO\resource\app\scripts" "\\192.168.10.220\wwwroot\MediCSV5\app\scripts"
mklink /D "C:\GO\resource\engine\db" "\\192.168.10.220\wwwroot\MediCSV5\engine\db"
mklink /D "C:\GO\resource\engine\css" "\\192.168.10.220\wwwroot\MediCSV5\engine\css"
mklink /D "C:\GO\resource\engine\scripts" "\\192.168.10.220\wwwroot\MediCSV5\engine\scripts"

