function Get-KubeResource(){

[Cmdletbinding(DefaultParameterSetName = "Namespace")]
param(
[Parameter(Mandatory=$true, ParameterSetName = "Namespace")]
[string]$Namespace,

[Parameter(Mandatory=$true, ParameterSetName = "Labelselect")]
[string]$Labelselector
)

process{
	switch ($PSCmdlet.ParameterSetName) {
		'Namespace' {
		    kubectl get pod --namespace $Namespace	
		}
		'Labelselect' {
			kubectl get pod --all-namespaces -l $Labelselector
		}
	}
}


}
