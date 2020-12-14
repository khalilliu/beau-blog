FROM nginx:alpine
LABEL MAINTAINER="lkhaliliu@gmail.com"

COPY .docker-compose/nginx/conf.d/admin.conf /etc/nginx/conf.d/admin.conf
COPY --from=0 /bb_admin/dist /usr/share/nginx/html/admin
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/admin.conf
RUN ls -al /usr/share/nginx/html
CMD ls -al /usr/share/nginx/html