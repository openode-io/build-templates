FROM hayd/deno:alpine

WORKDIR /opt/app

ENV PORT=80

# COPY deps.ts .
# RUN deno cache deps.ts

COPY . .
RUN deno cache main.ts

CMD ["run", "--allow-env", "--allow-net", "main.ts"]
