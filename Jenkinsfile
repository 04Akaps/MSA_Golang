def mainDir="."
def dockerImgName="MSA_Go"
def dockerRegistory="hojin"
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
                // script {
                //     docker.build("${dockerRegistory}/${dockerImgName}:$currentBuild.number")
                // }
                // sh '/usr/local/bin/docker build -t hojin/MSA_Go:$currentBuild.number .' 
                 sh '/usr/local/bin/docker --version' 
            }
        }

        // stage("Push To Image to ECR") {
        //     steps {
        //         sh 'docker push ${ecrUrl}/${dockerImgName}:${currentBuild.number}'
        //     }
        // }

    }
}