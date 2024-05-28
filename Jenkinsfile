pipeline {
  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }
  triggers {
    cron('@weekly')
  }
  agent {
    docker {
      image 'golang:1.17'
      label "docker"
      args "-v /tmp:/.cache"
    }
  }
  stages {
    stage("Prepare dependencies") {
      steps {
        sh 'go install github.com/jstemmer/go-junit-report/v2@latest'
        sh 'go mod download'
      }
    }
    stage("Test") {
      steps {
        sh 'go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml'
      }
      post {
        failure {
            mail to: 'sysadmin@brightbox.co.uk',
             subject: "Gobrightbox Tests Failed: ${currentBuild.fullDisplayName}",
             body: "${env.BUILD_URL}"
        }
        always {
          junit "report.xml"
        }
      }
    }
  }
}
