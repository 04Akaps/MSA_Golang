def mainDir="."
def ecrLoginHelper="docker-credential-ecr-login"
def region="ap-northeast-1"
def ecrUrl="297064282309.dkr.ecr.ap-southeast-2.amazonaws.com/msago"
def repository="MSA_GO"
def deployHost=""


pipeline {
    agent any

    stages {
        stage("Pull codes from Github")
            steps {
                checkout scm
            }
        stage("Test Cli") {
            steps {
                echo ls -al
            }
        }
    }
}