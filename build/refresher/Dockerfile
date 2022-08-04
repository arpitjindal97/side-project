# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM temp:latest as base_img

FROM alpine:latest
WORKDIR /work
COPY --from=base_img /work/refresher /work/refresher
EXPOSE 8080
ENTRYPOINT /work/refresher