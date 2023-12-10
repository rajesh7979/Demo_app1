pipeline {
    agent any
    tools {
      maven 'maven3'
      jdk 'JDK8'
      jfrog 'cli'         
    }
    environment {
        DOCKER_IMAGE_NAME = "slk.jfrog.io/fis-demo-dockerhub/app-image.${BUILD_ID}.${env.BUILD_NUMBER}"
        ARTIFACTORY_ACCESS_TOKEN = credentials('jf_access_token')
        WEBHOOK_URL = credentials("webhook_url")
        BUILD_NAME = "${JOB_NAME}"
        BUILD_NO = "${env.BUILD_NUMBER}"
       
    }
    options {
        office365ConnectorWebhooks([
            [name: "Office 365", Url: credentials("webhook_url"), notifyBackToNormal: true, notifyFailure: true, notifyRepeatedFailure: true, notifySuccess: true, notifyAborted: true]
        ])
      }
    stages {
        stage ('Checkout') {
            steps {
                git branch: 'main', credentialsId: 'github', url: 'https://github.com/rajesh7979/Demo_app1.git'
            }
        }
        stage ('Build') {
            steps {        
                sh 'mvn clean install'
                sh 'cp ./webapp/target/*.war ./'
                sh 'pwd'
                sh 'docker images'
            }
        }
        stage('Push artifacts into artifactory') {
             steps {
                   sh 'curl -fL https://getcli.jfrog.io | sh'
                   sh './jfrog rt u --url https://slk.jfrog.io/artifactory --access-token ${ARTIFACTORY_ACCESS_TOKEN} ./*.war  fis-demo-release/'
                   jf 'rt build-publish'
                 //  sh  './jfrog rt bp  --url https://slk.jfrog.io/artifactory --access-token ${ARTIFACTORY_ACCESS_TOKEN} ${JOB_NAME} ${BUILD_NUMBER}'
             }
        }
        /*stage('Upload'){
            steps{
                rtUpload (
                 serverId:'jfrog-artifactory',
                  spec: '''{
                   "files": [
                      {
                      "pattern": "webapp/target/*.war",
                      "target": "java-application-generic-local"
                      }
                            ]
                           }''',
                  )
            }
        }*/
        /*stage ('Publish build file info') {
            steps {
                rtPublishBuildInfo (
                    serverId: "jfrog-artifactory"
                )
            } 
        } */   
        stage ('SonarQube analysis') {
            steps {
                withSonarQubeEnv('sonar') {
                sh 'mvn sonar:sonar'
                }
            }
        }   
        stage('Build Docker image') {
			steps {
				script {
					docker.build("$DOCKER_IMAGE_NAME", './')
				}
			}
		}
		stage('Trivy Scan') {
            steps {
                // Install trivy
                sh 'sudo apt-get -y install wget apt-transport-https gnupg lsb-release'
                sh 'wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -'
                sh  'echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | sudo tee -a /etc/apt/sources.list.d/trivy.list'
                sh  'sudo apt-get -y update'
                sh  'sudo apt-get -y install trivy'
                // Scan all vuln levels

                sh 'curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3'
                sh 'chmod 700 get_helm.sh'
                sh './get_helm.sh'
                sh 'helm repo add aquasecurity https://aquasecurity.github.io/helm-charts/'
                sh 'helm repo update'
                sh 'helm search repo trivy'
          //      sh 'helm install my-trivy aquasecurity/trivy'
               sh 'trivy -h'
               sh 'mkdir -p reports'
               sh 'sudo apt-get -y install tree'
               sh 'trivy image $DOCKER_IMAGE_NAME  --output report.html || true'
               sh 'tree'
               sh 'ls -lrth'
              // publishHTML([allowMissing: false, alwaysLinkToLastBuild: false, keepAll: false, reportDir: 'reports', reportFiles: 'index.html', reportName: 'Trivy Scan', reportTitles: 'Trivy Scan', useWrapperFileDirectly: true])
              

            }
		}    
        stage('Scan and push image') {
			steps {
				dir('${WORKSPACE}') {
					// Scan Docker image for vulnerabilities
				//	jf 'docker scan $DOCKER_IMAGE_NAME'

					// Push image to Artifactory
					jf 'docker push $DOCKER_IMAGE_NAME'
				}
			}
		}

		stage('Publish build info') {
			steps {
				jf 'rt build-publish'
			}
		}
        /*stage ('Build docker image') {
            steps {
                script {
                    sh 'sudo docker build -t fis-demo/app-image.${BUILD_ID}  .'
                    sh 'sudo docker images'
                }
            }
        }*/
        stage('Notification') {
            steps {
                office365ConnectorSend webhookUrl: '$WEBHOOK_URL',
                message: 'build is success',
                status: 'Success'            
            }
        }
    }        
       
    //Job Status check
    post {
        success {
            publishHTML([allowMissing: false, alwaysLinkToLastBuild: false, keepAll: false, reportDir: '.', reportFiles: 'report.html', reportName: 'Trivy Scan', reportTitles: 'Trivy Scan', useWrapperFileDirectly: true])
            echo 'Build successfully!'
         //   office365ConnectorSend(
          //  status: "Build Success",
        //    webhookUrl: "$WEBHOOK_URL",
         //   color: '00ff00',
          //  message: "Success"
        //    )      
        }

        failure {
           // office365ConnectorSend(
         //    status: "Build Failed",
          //  webhookUrl: "$WEBHOOK_URL",
        //    color: 'ff4000',
          //  message: "The build has failed, please check build logs"
        // )
            echo 'Build Failed'
        }
    }
}
