pipeline {
    agent any
    environment {
        GIT_REPO_NAME = 'Demo_app'
        GITHUB_BRANCH = 'main'
        GIT_USER_NAME = 'rajesh7979'
        PATH = "$PATH:/usr/share/maven"
    }
    stages {
        stage ('Checkout') {
            steps {
                git branch: 'main', credentialsId: 'github', url: 'https://github.com/rajesh7979/Demo_app.git'
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
                    sh -s 'docker build -t fis-demo/app-image.${BUILD_ID}  .'
                    sh -s 'docker images'
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
