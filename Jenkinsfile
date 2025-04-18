pipeline {
    agent any

    tools {
        go 'golang'
    }

    stages {
        stage('Build') {
            steps {
                sh '''
                    go version

                    echo "Building..."
                    go mod download
                    go build -o main cmd/api/main.go
                '''
            }
        }

        stage('Test') {
            steps {
                sh '''
                    echo "Testing..."
                    go test ./... -v
                '''
            }
        }
    }
}