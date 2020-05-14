FROM hayd/alpine-deno:1.0.0

WORKDIR /opt/app

ENV PORT=80

# COPY deps.ts .
# RUN deno cache deps.ts

COPY . .
RUN deno cache main.ts

CMD ["run", "--allow-env", "--allow-net", "main.ts"]