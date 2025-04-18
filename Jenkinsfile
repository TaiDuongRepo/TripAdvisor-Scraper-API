pipeline {
    agent any

    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.23-alpine'
                    reuseNode true
                }
            }
            steps {
                sh '''
                    go version

                    make test
                '''
            }
        }
    }
}