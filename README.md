Sample gRPC code being created from following udemy course
Will be used to add gRPC to bakery project later

## SSL in Windows

You need to run the powershell script ssl.ps1 in folder ssl

This requires openssl.exe application.  This exe can be downloaded from https://www.openssl.org/ but it is easier to use the one in your git folder. On this machine it is located in `C:\Program Files\Git\usr\bin` . Ensure that this is in your Environment variables path.
ssl.ps1 is not digitally signed. To force it to run enter `Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass` in your powershell application 
After running the files generated by openssl will be added to the ssl folder 