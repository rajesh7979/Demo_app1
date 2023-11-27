pipeline {
    agent any
    environment {
        GIT_REPO_NAME = 'Demo_app'
        GITHUB_BRANCH = 'main'
        GIT_USER_NAME = 'rajesh7979'
    }
    stages {
        stage ('Checkout') {
            steps {
                git credentialsId: 'github', url: 'https://github.com/rajesh7979/Demo_app.git'
            }
        }
        stage ('Build') {
            steps {
                sh 'mvn clean install'
            }
        }
        stage ('Build docker image') {
            steps {
                script {
                    sh 'docker build -t nagarjunacse02/test-image .'
                    sh 'docker images'
                }
            }
        }
        
    }
    //Job Status check
    post {
        success {
            echo 'Build successfully!'
        }
 
        failure {
            echo 'Build Failed'
        }
    }
}
