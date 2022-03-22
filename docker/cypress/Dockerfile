FROM cypress/included:9.1.1

WORKDIR /root

RUN npm install "cypress-failed-log@2.9.2"

ENV CYPRESS_VIDEO=false
ENV CYPRESS_baseUrl=http://gitline:3001

COPY cypress.json .
COPY cypress cypress