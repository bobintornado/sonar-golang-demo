# sonar-golang-demo

Demo project contain bad golang codes for generating sonarqube reports

## Usage

Assuming GoMetaLinter is installed as `gometalinter.v1`, running following command to generate the checkstyle xml report

```
gometalinter.v1 --checkstyle > report.xml
```

Then use sonar-scanner to analyse the project
