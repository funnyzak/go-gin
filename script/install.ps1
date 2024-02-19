$taskName = "go-gin"
$programPath = Join-Path $PSScriptRoot "go-gin.exe"
$workingDir = $PSScriptRoot

# check if Administrator
$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")
if (-not $isAdmin) {
    Write-Host "Please run this script as Administrator."
    exit
}

if ($args[0] -eq "enable") {
    # first check if the task exists
    $task = Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue
    if ($task -ne $null) {
        Write-Host "Task $taskName already exists."
        $status = Get-ScheduledTask -TaskName $taskName | Select-Object State
        if ($status.State -eq "Running") {
            Write-Host "Task $taskName is already running."
        } else {
            Start-ScheduledTask -TaskName $taskName
            Write-Host "Task $taskName started."
        }
    } else {
        Write-Host "Creating task $taskName..."
        $action = New-ScheduledTaskAction -Execute $programPath -WorkingDirectory $workingDir
        $trigger = New-ScheduledTaskTrigger -AtStartup
        $settings = New-ScheduledTaskSettingsSet -AllowStartIfOnBatteries -DontStopIfGoingOnBatteries -StartWhenAvailable -DontStopOnIdleEnd
        $principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -LogonType ServiceAccount
        Register-ScheduledTask -TaskName $taskName -Action $action -Trigger $trigger -Settings $settings -Principal $principal
        Start-ScheduledTask -TaskName $taskName
        Write-Host "Task $taskName created."
    }
} elseif ($args[0] -eq "disable") {
    # first check if the task exists
    $task = Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue
    if ($task -eq $null) {
        Write-Host "Task $taskName does not exist."
    } else {
        Write-Host "Deleting task $taskName..."
        Stop-ScheduledTask -TaskName $taskName
        Unregister-ScheduledTask -TaskName $taskName -Confirm:$false
        Write-Host "Task $taskName deleted."
    }
} elseif ($args[0] -eq "start") {
    # first check if the task exists
    $task = Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue
    if ($task -eq $null) {
        Write-Host "Task $taskName does not exist. Please enable it first."
    } else {
        Write-Host "Starting task $taskName..."
        Start-ScheduledTask -TaskName $taskName
        Write-Host "Task $taskName started."
    }
} elseif ($args[0] -eq "stop") {
    # first check if the task exists
    $task = Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue
    if ($task -eq $null) {
        Write-Host "Task $taskName does not exist. Please enable it first."
    } else {
        Write-Host "Stopping task $taskName..."
        Stop-ScheduledTask -TaskName $taskName
        Write-Host "Task $taskName stopped."
    }
} elseif ($args[0] -eq "restart") {
    # first check if the task exists
    $task = Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue
    if ($task -eq $null) {
        Write-Host "Task $taskName does not exist. Please enable it first."
    } else {
        Write-Host "Restarting task $taskName..."
        Restart-ScheduledTask -TaskName $taskName
        Write-Host "Task $taskName restarted."
    }
} elseif ($args[0] -eq "status") {
    # first check if the task exists
    $task = Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue
    if ($task -eq $null) {
        Write-Host "Task $taskName does not exist. Please enable it first."
    } else {
        $status = Get-ScheduledTask -TaskName $taskName | Select-Object State
        if ($status.State -eq "Running") {
            Write-Host "Task $taskName is running."
        } else {
            Write-Host "Task $taskName is not running."
        }
    }
} else {
    Write-Host "Usage: go-gin.ps1 [install|uninstall|enable|disable|start|stop|restart|status]"
    Write-Host "  enable   - Enable the service."
    Write-Host "  disable  - Disable the service."
    Write-Host "  start    - Start the service."
    Write-Host "  stop     - Stop the service."
    Write-Host "  restart  - Restart the service."
    Write-Host "  status   - Show the status of the service."
    exit
}
