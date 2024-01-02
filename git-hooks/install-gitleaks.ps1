param(
    [string]$InstallDir,
    [string]$DownloadLink
)

# Create the installation directory if it doesn't exist
New-Item -ItemType Directory -Force -Path $InstallDir

# Download Gitleaks zip
Invoke-WebRequest -Uri $DownloadLink -OutFile "gitleaks.zip"

# Extract the contents
Add-Type -AssemblyName System.IO.Compression.FileSystem
[System.IO.Compression.ZipFile]::ExtractToDirectory('gitleaks.zip', $InstallDir)

# Find the extracted Gitleaks executable
$exePath = Get-ChildItem -Path $InstallDir -Recurse -Filter "gitleaks.exe" | Select-Object -ExpandProperty FullName

# Move the executable to the installation directory
Move-Item -Path $exePath -Destination "$InstallDir\gitleaks" -Force

# Clean up the zip file
Remove-Item -Path 'gitleaks.zip'

Write-Output "Gitleaks installed to: $InstallDir"