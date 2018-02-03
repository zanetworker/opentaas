FROM jenkins/jenkins:lts-alpine

# Create Jenkins Folders
USER root
RUN mkdir /var/log/jenkins && chown -R jenkins:jenkins /var/log/jenkins
RUN mkdir /var/cache/jenkins && chown -R jenkins:jenkins /var/cache/jenkins
RUN mkdir /var/jenkins_jobs && chown -R jenkins:jenkins /var/jenkins_jobs

USER jenkins
ENV JENKINS_OPTS="--handlerCountMax=300 --logfile=/var/log/jenkins/jenkins.log --webroot=/var/cache/jenkins/war"
ENV JAVA_OPTS="-Djenkins.install.runSetupWizard=false"

# Copy init scripts to run upon Jenkins start
USER root
COPY out/security.groovy /usr/share/jenkins/ref/init.groovy.d/
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt

# Fix file permissions for /usr/share/jenkins/ref/
USER root
RUN chown -R jenkins:jenkins /usr/share/jenkins/ref/
USER jenkins

EXPOSE 5000

# Copy custom start script to run parametrization during container start
#COPY start-jenkins.sh /usr/local/bin/start-jenkins.sh
#ENTRYPOINT ["/usr/local/bin/start-jenkins.sh"]
#ENTRYPOINT ["/bin/sh"]