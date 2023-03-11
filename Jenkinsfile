def mainDir="."
def dockerImgName="msago"
def dockerRegistory="yuhojin"
def ecrLoginHelper="docker-credential-ecr-login"
def region="ap-northeast-1"
def ecrUrl="297064282309.dkr.ecr.ap-northeast-2.amazonaws.com"
def repository="go_msa"
def deployHost=""

pipeline {
    agent any

    stages {

        stage("Pull codes from Github"){
            steps {
                checkout scm
            }
        }

        stage("Test Cli") {
            steps {
               sh '''
                    export GOPATH=$WORKSPACE
                    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
                    go version
                '''
            }
        }

        stage("Build Docker Image"){
            steps {
                sh "/usr/local/bin/docker build -t ${dockerRegistory}/${dockerImgName}:$currentBuild.number ."
            }
        }

        stage("Push To AWS ECR") {
            steps {
                script {
                    def ecrCredentials = ecrLogin()
                    withCredentials([[$class: 'AmazonWebServicesCredentialsBinding', credentialsId: ecrCredentials]]) {
                        sh "docker login -u AWS -p ${ecrCredentials.secretKey} ${ecrCredentials.accessKey}"
                    }
                }
            }
        }
    }
}