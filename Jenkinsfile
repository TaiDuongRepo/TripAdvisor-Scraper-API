pipeline {
    agent any

    tools { go '1.24.1' }

    stages {
        stage('Build') {
            steps {
                sh 'make --version'
                sh 'make build'
            }
        }

        stage('Test') {
            steps {
                sh 'make test'
            }
        }
    }
}