# PlayingWithGoAndVue
Playing with GoLang and Vue JS


This is a demo application developed using GoLang for backend. 
This instructions assume that you are using Ubuntu Linux. 
However, to run on any other platform, please make sure you download all the dependencies. Please see the **DownloadDependecy.sh** for the full list of required package. 
The SQLITE3 driver is not compatible with windows. 
The solutions is to, use MinGW gcc (windows XP or later).

1. Install GoLang 1.9+.
2. Download the project.
3. Go to the folder and open terminal/command prompt.
4. To download dependecies, run "./DownloadDependencies.sh"
4. To run the application, "./run.sh".

## Tips: 
- If error occurs, first check port number in main.go file. The port number is in the main function.

```go
func main() {
   ...
   ng.Run(":8080")
   ...
}
```
- Still don't work, create issue.
