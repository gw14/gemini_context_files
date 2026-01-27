function Get-AppStatus{
	#A function that finds a specific process and returns basic info about it.
	param(
	[Parameter(Mandatory=$true)]
	[string]$ProcessName
	)
	
	get-process $ProcessName
}
