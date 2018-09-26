# PlayingWithGoAndVue
Playing with GoLang and Vue JS

This is a demo application developed using GoLang for backend. 
This instructions assume that you are using Ubuntu Linux. 
This application runs on google app engine standard. 

However, to run on any other platform, please make sure you download all the dependencies. Please see the **DownloadDependecy.sh** for the full list of required package. 

1. Install GoLang 1.9+.
2. Install Google App Engine SDKs for golang. See instructions here https://cloud.google.com/appengine/docs/standard/go/download
3. When you run 
   ```bash
     gcloud components list
   ```
   Output should be

#### Components    

|     Status    |                         Name                         |            ID            |    Size   |
|---------------|------------------------------------------------------|--------------------------|-----------|
| Not Installed | Cloud Bigtable Command Line Tool                     | cbt                      |   6.4 MiB |
| Not Installed | Cloud Bigtable Emulator                              | bigtable                 |   4.3 MiB |
| Not Installed | Cloud Datalab Command Line Tool                      | datalab                  |   < 1 MiB |
| Not Installed | Cloud Datastore Emulator (Legacy)                    | gcd-emulator             |  38.1 MiB |
| Not Installed | Cloud Pub/Sub Emulator                               | pubsub-emulator          |  33.4 MiB |
| Not Installed | Cloud SQL Proxy                                      | cloud_sql_proxy          |   3.8 MiB |
| Not Installed | Emulator Reverse Proxy                               | emulator-reverse-proxy   |  14.5 MiB |
| Not Installed | Google Cloud Build Local Builder                     | cloud-build-local        |   4.5 MiB |
| Not Installed | Google Container Local Builder                       | container-builder-local  |   4.5 MiB |
| Not Installed | Google Container Registry's Docker credential helper | docker-credential-gcr    |   1.8 MiB |
| Not Installed | gcloud app Java Extensions                           | app-engine-java          | 107.5 MiB |
| Not Installed | gcloud app PHP Extensions                            | app-engine-php           |           |
| Not Installed | gcloud app Python Extensions (Extra Libraries)       | app-engine-python-extras |  28.5 MiB |
| Installed     | App Engine Go Extensions                             | app-engine-go            | 153.3 MiB |
| Installed     | BigQuery Command Line Tool                           | bq                       |   < 1 MiB |
| Installed     | Cloud Datastore Emulator                             | cloud-datastore-emulator |  17.7 MiB |
| Installed     | Cloud SDK Core Libraries                             | core                     |   8.5 MiB |
| Installed     | Cloud Storage Command Line Tool                      | gsutil                   |   3.5 MiB |
| Installed     | gcloud Alpha Commands                                | alpha                    |   < 1 MiB |
| Installed     | gcloud Beta Commands                                 | beta                     |   < 1 MiB |
| Installed     | gcloud app Python Extensions                         | app-engine-python        |   6.2 MiB |
| Installed     | kubectl                                              | kubectl                  |   < 1 MiB |


4. Download the project to ~go/src/ or $GOPATH/src/ folder
5. Go to the folder and open terminal/command prompt.
6. To download dependecies, run "./DownloadDependencies.sh"
7. To run the application go to the **main** folder and run "dev_appserver.py app.yaml" on terminal.


## Tips: 
- App engine runs on port 8080 in development environment.
- Still don't work, create issue.
