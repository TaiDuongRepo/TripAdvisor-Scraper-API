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

                    echo "Testing..."
                    go test ./... -v
                '''
            }
        }
    }
}