
FROM akrck02/weblogic:10.3.6

# Args
ARG eHDIV_ENABLE=true
ARG eHDIV_MAIL=
ARG eHDIV_CONSOLE_URL=http://www.otc.ejiedes.eus/hdiv-console-services

ARG eHOSTNAME=wls1036
ARG eADMIN_PORT_SSL=7002
ARG eDEBUG_PORT=8453
ARG eADMIN_PORT=7001
ARG eMEMORY=512m

# Environment variables
ENV HDIV_ENABLE=${eHDIV_ENABLE}
ENV CONFIG_JVM_ARGS=-Dweblogic.security.SSL.ignoreHostnameVerification=true
ENV HOSTNAME=${eHOSTNAME}
ENV TERM=xterm
ENV HDIV_MAIL=${eHDIV_MAIL}
ENV USER_MEM_ARGS="-Djava.security.egd=file:/dev/./urandom -Xms${eMEMORY} -Xmx${eMEMORY} -XX:PermSize=${eMEMORY} -XX:MaxPermSize=${eMEMORY}"
ENV DOMAIN_NAME=base_domain
ENV PATH=/usr/java/jdk-7/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/srv/wls1036/wlserver/common/bin:/srv/wls1036/user_projects/domains/base_domain:/srv/wls1036/user_projects/domains/base_domain/bin:/srv/wls1036/scripts
ENV ADMIN_HOST=localhost
ENV ADMIN_PORT_SSL=${eADMIN_PORT_SSL}
ENV PWD=/srv/wls1036/user_projects/domains/base_domain
ENV DEBUG_PORT=${eDEBUG_PORT}
ENV ADMIN_PORT=${eADMIN_PORT}
ENV MW_HOME=/srv/wls1036
ENV JAVA_HOME=/usr/java/jdk-7
ENV debugFlag=true
ENV LANG=en_US.UTF-8
ENV JAVA_VERSION=1.7.0_80
ENV TZ=Europe/Madrid
ENV SHLVL=1
ENV HOME=/srv/wls1036
ENV PRE_DOMAIN_HOME=/srv/wls1036/user_projects
ENV HDIV_CONSOLE_URL=${HDIV_CONSOLE_URL}




# Add files required to build this image
COPY ./resources/data/* /tmp/domain-resources/

# Configure user and workdir
USER wls1036
WORKDIR /srv/wls1036/user_projects/domains/base_domain
RUN /srv/wls1036/scripts/create-wls-resources.sh /tmp/domain-resources

# Start WebLogic (container startup)
CMD ["/srv/wls1036/scripts/startWebLogic.sh"]