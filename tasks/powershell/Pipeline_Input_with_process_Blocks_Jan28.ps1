function Get-AppStatus{
        #A function that finds a specific process and returns basic info about it.
        param(
        [Parameter(Mandatory=$true,ValueFromPipeline)]
        [string]$ProcessName
        )
        process{
          get-process $ProcessName
        }
}


