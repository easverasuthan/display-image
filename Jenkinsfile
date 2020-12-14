pipeline {

    agent {
        node {
            label 'slave_node1'
        }
    }
    stages {
        
        stage('Code Checkout') {
            steps {
             checkout scm
            }
        }
         stage ('Deploying our Api'){
            steps{
             checkout scm
             sh 'docker build -t display-image .'
             sh 'docker tag display-image 603389930669.dkr.ecr.us-east-1.amazonaws.com/display-image:latest'
             sh '$(aws ecr get-login --no-include-email --region us-east-1 > /dev/null)'
             sh 'docker push 603389930669.dkr.ecr.us-east-1.amazonaws.com/display-image:latest'
             sh 'aws ecs update-service --cluster pratilipi --service display-image --force-new-deployment --region us-east-1'
            }

         }
    }   
}
