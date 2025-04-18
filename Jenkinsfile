pipeline {
    agent any

    tools { go '1.24.1' }

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