pipeline {
    agent any

    tools { go '1.24.1' }

    stages {
        // stage('Build') {
        //     steps {
        //         sh 'make --version'
        //         sh 'make build'
        //     }
        // }

        // stage('Test') {
        //     steps {
        //         sh 'make test'
        //     }
        // }

        // stage('Docker Build') {
        //     steps {
        //         sh 'docker build -t nghex/tripadvisor-scraper-api:latest .'
        //     }
        // }

        stage('Docker Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'nghex', passwordVariable: 'PASSWORD', usernameVariable: 'USERNAME')]) {
                    sh 'echo $USERNAME'
                    // sh "docker login -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
                    // sh 'docker push nghex/tripadvisor-scraper-api:latest'
                }
            }
        }
    }
}