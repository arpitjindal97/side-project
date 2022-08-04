FROM temp:latest as base_img

# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:latest
WORKDIR /work
COPY --from=base_img /work/apiserver /work/apiserver
EXPOSE 8080
ENTRYPOINT /work/apiserver