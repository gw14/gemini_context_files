function Get-AppStatus{
        #A function that finds a specific process and returns basic info about it.
        param(
        [Parameter(Mandatory=$true,ValueFromPipeline)]
        [string]$ProcessName
        )
        begin{
	$ProcessList=@()
	}
	process{
          foreach ($process in get-process $ProcessName){
	  	$ProcessList +=[pscustomobject]@{"Name"=$process.ProcessName;"ID"=$process.Id;"Cpu"=$process.cpu}
	  }
        }
	end{
		return $ProcessList
	}
}


