pipeline {
    agent any

    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.24.1-alpine'
                    reuseNode true
                }
            }
            steps {
                sh '''
                    go version

                    echo "Building..."
                    go build -o main cmd/api/main.go
                '''
            }
        }

        stage('Test') {
            agent {
                docker {
                    image 'golang:1.24.1-alpine'
                    reuseNode true
                }
            }
            steps {
                sh '''
                    echo "Testing..."
                    go test ./... -v
                '''
            }
        }
    }
}