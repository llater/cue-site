#!/usr/bin/env groovy

pipeline {
    agent {
        label 'master'
    }
    options {
        buildDiscarder(logRotator(daysToKeppStr: '60', numToKeepStr: '100'))
        timeout(time: 20)
    }
    stages {
   
    }
}
